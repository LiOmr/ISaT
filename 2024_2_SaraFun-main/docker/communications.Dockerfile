FROM alpine:latest

ENV EXECUTABLE=communications

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

COPY --from=sparkit-builder:latest --chmod=755 /application/${EXECUTABLE} /app

CMD /app/${EXECUTABLE}