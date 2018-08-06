ARG FLUX_TAG
FROM quay.io/influxdb/ifqld:$FLUX_TAG
ADD ./latest/wait_for_influxdb.sh ./

ENTRYPOINT ["./wait_for_influxdb.sh"]
