@echo off

if "%1"=="" (
@echo.
@echo sandbox commands:
@echo   up           - spin up the sandbox environment
@echo   down         - tear down the sandbox environment
@echo   restart      - restart the sandbox
@echo   influxdb     - attach to the influx cli
@echo.
@echo   enter influxdb^|kapacitor^|chronograf^|telegraf - enter the specified container
@echo   logs  influxdb^|kapacitor^|chronograf^|telegraf - stream logs for the specified container
@echo.
@echo   delete-data  - delete all data created by the TICK Stack 
@echo   docker-clean - stop and remove all running docker containers and images
@echo   rebuild-docs - rebuild the documentation image
@echo.
)

if "%1"=="up" (
    echo Spinning up Docker Images...
    echo If this is your first time starting sandbox this might take a minute...
    docker-compose -p sandbox up -d
    echo Opening tabs in browser...
    sleep 3
    start http://localhost:8888
    start http://localhost:3000
    )

if "%1"=="down" (
    echo Stopping and removing running containers...
    docker-compose -p sandbox down
    )

if "%1"=="restart" (
    echo Stopping all processes...
    docker-compose -p sandbox down > NUL 2>&1
    echo Starting all processes...
    docker-compose -p sandbox up -d > NUL 2>&1
    echo Services available"
    )

if "%1"=="influxdb" (
    echo Following the logs from the influxdb container...
    docker exec -it sandbox_influxdb_1 /usr/bin/influx
    )

if "%1"=="enter" (
if "%2"=="" (
    echo sandbox enter ^(influxdb^|chronograf^|kapacitor^|telegraf^)
    )
if "%2"=="influxdb" (
    echo Entering /bin/bash session in the influxdb container...
    docker exec -it sandbox_influxdb_1 /bin/bash
    )
if "%2"=="chronograf" (
    echo Entering /bin/sh session in the chronograf container...
    docker exec -it sandbox_chronograf_1 /bin/sh
    )
if "%2"=="kapacitor" (
    echo Entering /bin/bash session in the kapacitor container...
    docker exec -it sandbox_kapacitor_1 /bin/bash
    )
if "%2"=="telegraf" (
    echo Entering /bin/bash session in the telegraf container...
    docker exec -it sandbox_telegraf_1 /bin/bash
    )
)

if "%1"=="logs" (
if "%2"=="" (
    echo sandbox logs ^(influxdb^|chronograf^|kapacitor^|telegraf^)
    )
if "%2"=="influxdb" (
    echo Following the logs from the influxdb container...
    docker logs -f sandbox_influxdb_1
    )
if "%2"=="chronograf" (
    echo Following the logs from the chronograf container...
    docker logs -f sandbox_chronograf_1
    )
if "%2"=="kapacitor" (
    echo Following the logs from the kapacitor container...
    docker logs -f sandbox_kapacitor_1
    )
if "%2"=="telegraf" (
    echo Following the logs from the telegraf container...
    docker logs -f sandbox_telegraf_1
    )
)

if "%1"=="delete-data" (
    echo deleting all influxdb, kapacitor and chronograf data...
    del /F /S /Q kapacitor\data influxdb\data chronograf\data
    )

if "%1"=="docker-clean" (
    echo Stopping all running containers...
    FOR /F %%C in ('docker ps -a -q') DO docker stop %%C > NUL 2>&1
    echo Removing all containers...
    FOR /F %%C in ('docker ps -a -q') DO docker rm %%C > NUL 2>&1
    echo Removing TICK images...
    docker rmi sandbox_documentation influxdb:1.2.2 telegraf:1.2.1 kapacitor:1.2.0 quay.io/influxdb/chronograf:1.2.0-beta7 > NUL 2>&1
    )

if "%1"=="rebuild-docs" (
    echo Rebuilding documentation container...
    docker build -t sandbox_documentation documentation/ > NUL 2>&1
    echo Restarting...
    docker-compose -p sandbox down > NUL 2>&1
    docker-compose -p sandbox up -d > NUL 2>&1
    )

