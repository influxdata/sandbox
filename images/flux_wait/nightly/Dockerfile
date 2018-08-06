ARG FLUX_TAG
FROM quay.io/influxdb/flux:$FLUX_TAG
ENTRYPOINT ["./wait_for_influxdb.sh"]
#CMD ["fluxd","--storage-hosts","influxdb:8082"]
