FROM alpine:3.4

EXPOSE 3000:3000

RUN mkdir -p /documentation

COPY builds/documentation /documentation/
COPY static/ /documentation/static

CMD ["/documentation/documentation", "-filePath", "/documentation/"]