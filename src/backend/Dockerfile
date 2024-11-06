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

EXPOSE 8080
ENTRYPOINT ["./gcloud-serverless-gym"]