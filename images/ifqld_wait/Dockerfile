FROM "${IFQL_TAG}"
ADD ./wait_for_influxdb.sh ./

ENTRYPOINT ["./wait_for_influxdb.sh"]
