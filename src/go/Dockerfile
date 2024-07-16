# Dockerfile.production

FROM golang:1.22-alpine as builder

ENV APP_HOME /go/src/gcloud-serverless-gym

WORKDIR "$APP_HOME"

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o gcloud-serverless-gym

FROM alpine

ENV APP_HOME /go/src/gcloud-serverless-gym

COPY --from=builder "$APP_HOME"/gcloud-serverless-gym /

COPY --from=gcr.io/datadoghq/serverless-init:1-alpine /datadog-init /app/datadog-init
ENTRYPOINT ["/app/datadog-init"]
ENV DD_SERVICE=gcloud-go
ENV DD_ENV=gcloud-go
ENV DD_VERSION=1
ENV DD_TRACE_SAMPLE_RATE 1

EXPOSE 8080
CMD ["./gcloud-serverless-gym"]