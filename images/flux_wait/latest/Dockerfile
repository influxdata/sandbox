ARG FLUX_TAG
FROM quay.io/influxdb/fluxd:$FLUX_TAG
ADD ./latest/wait_for_influxdb.sh ./

ENTRYPOINT ["./wait_for_influxdb.sh"]
