FROM quay.io/influxdb/ifqld:latest
ADD ./wait_for_influxdb.sh ./

ENTRYPOINT ["./wait_for_influxdb.sh"]
