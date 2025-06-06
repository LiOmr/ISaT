package main

import (
	"context"
	"database/sql"
	"fmt"
	delivery "github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/message/delivery/grpc"
	generatedMessage "github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/message/delivery/grpc/gen"
	MessageRepo "github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/message/repo/message"
	ReportRepo "github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/message/repo/report"
	MessageUsecase "github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/message/usecase/message"
	ReportUsecase "github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/message/usecase/report"
	grpcmetrics "github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/metrics"
	"github.com/go-park-mail-ru/2024_2_SaraFun/internal/pkg/middleware/grpcMetricsMiddleware"
	"github.com/go-park-mail-ru/2024_2_SaraFun/internal/utils/config"
	"github.com/go-park-mail-ru/2024_2_SaraFun/internal/utils/connectDB"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{"stdout", "/tmp/sparkit_logs"},
		ErrorOutputPaths: []string{"stderr", "/tmp/sparkit_err_logs"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
			TimeKey:    "ts",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	//defer logger.Sync()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("failed to sync logger", zap.Error(err))
		}
	}()
	envCfg, err := config.NewConfig(logger)
	if err != nil {
		log.Fatal(err)
	}
	connStr, err := connectDB.GetConnectURL(envCfg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("env", zap.Any("envCfg", envCfg))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(16)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(0)
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	metrics, err := grpcmetrics.NewGrpcMetrics("message")
	if err != nil {
		log.Fatalf("Error initializing grpc metrics: %v", err)
	}
	metricsMiddleware := grpcMetricsMiddleware.NewMiddleware(metrics, logger)
	messageRepo := MessageRepo.New(db, logger)
	reportRepo := ReportRepo.New(db, logger)
	messageUsecase := MessageUsecase.New(messageRepo, logger)
	reportUsecase := ReportUsecase.New(reportRepo, logger)
	messageDelivery := delivery.NewGRPCHandler(reportUsecase, messageUsecase, logger)
	gRPCServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}), grpc.ChainUnaryInterceptor(metricsMiddleware.ServerMetricsInterceptor))
	generatedMessage.RegisterMessageServer(gRPCServer, messageDelivery)

	router := mux.NewRouter()
	router.Handle("/api/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:    ":8034",
		Handler: router,
	}

	go func() {
		fmt.Println("Starting the server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	go func() {
		listener, err := net.Listen("tcp", ":8084")
		if err != nil {
			log.Printf("net listen error: %s", err.Error())
		}
		fmt.Println("grpc server running")
		if err := gRPCServer.Serve(listener); err != nil {
			log.Fatalf("bad serve")
		}
		fmt.Println("gRPC server stopped")
	}()
	fmt.Println("wait signal")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	gRPCServer.GracefulStop()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}
}
