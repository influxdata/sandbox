#!/bin/sh
# wait_for_influxdb.sh

set -e

until wget http://influxdb:8086/ping; do
  echo "InfluxDB is unavailable - sleeping"
  sleep 5
done

echo "InfluxDB is up - Starting Flux"
exec ./ifqld
