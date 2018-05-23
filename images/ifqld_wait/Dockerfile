FROM quay.io/influxdb/ifqld:v0.0.6
ADD ./wait_for_influxdb.sh ./

ENTRYPOINT ["./wait_for_influxdb.sh"]
