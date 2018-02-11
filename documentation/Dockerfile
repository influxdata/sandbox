FROM alpine:3.4

EXPOSE 3010:3010

RUN mkdir -p /documentation

COPY builds/documentation /documentation/
COPY static/ /documentation/static

CMD ["/documentation/documentation", "-filePath", "/documentation/"]
