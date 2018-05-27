ARG IFQL_TAG
FROM quay.io/influxdb/ifqld:$IFQL_TAG
ADD ./wait_for_influxdb.sh ./

ENTRYPOINT ["./wait_for_influxdb.sh"]
