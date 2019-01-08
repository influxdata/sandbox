FROM alpine:3.6

EXPOSE 3010:3000

RUN mkdir -p /documentation

COPY builds/documentation /documentation/
COPY static/ /documentation/static

CMD ["/documentation/documentation", "-filePath", "/documentation/"]
