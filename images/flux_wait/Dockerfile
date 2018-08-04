ARG FLUX_TAG
FROM quay.io/influxdb/flux:$FLUX_TAG
ADD ./wait_for_influxdb.sh ./

ENTRYPOINT ["./wait_for_influxdb.sh"]
