FROM golang:1.12-stretch as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags "--s -extldflags '-static' -X main.Version=git:$CI_BUILD_REF" -o "api-hub" ./...

FROM alpine
WORKDIR /app
COPY --from=builder /app/api-hub .
COPY data/ data/
COPY ui/ ui/
RUN chmod +x api-hub
EXPOSE 10000
CMD [ "./api-hub" ]
