FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY config-srv /config-srv
ENTRYPOINT /config-srv
LABEL Name=micro Version=0.0.1
EXPOSE 9600 9600
