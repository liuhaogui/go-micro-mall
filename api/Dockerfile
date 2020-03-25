FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY micro-api /micro-api
ENV CONSUL_ADDRESS=127.0.0.1:8500 \
    MICRO_REGISTRY=consul \
    MICRO_REGISTRY_ADDRESS=127.0.0.1:8500 \
    MICRO_API_HANDLER=http
ENTRYPOINT /micro-api
LABEL Name=micro Version=0.0.1
EXPOSE 9091 9091
