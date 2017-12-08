# TICK Sandbox

This repo is a quick way to get the entire TICK Stack spun up and working together. It uses [docker](https://www.docker.com/) to spin up the full TICK stack in a connected fashion. This is heavily tested on Mac and should mostly work on linux and Windows.

To get started you need a running docker installation. If you don't have one, you can download Docker for [Mac](https://www.docker.com/docker-mac) or [Windows](https://www.docker.com/docker-windows), or follow the installation instructions for Docker CE for your [Linux distribution](https://docs.docker.com/engine/installation/#server).

### Running 

To run the `sandbox`, simply use the convenient cli:

```bash
$ ./sandbox
sandbox commands:
  up           -> spin up the sandbox environment
  down         -> tear down the sandbox environment
  restart      -> restart the sandbox
  influxdb     -> attach to the influx cli
  
  enter (influxdb||kapacitor||chronograf||telegraf) -> enter the specified container
  logs  (influxdb||kapacitor||chronograf||telegraf) -> stream logs for the specified container
  
  delete-data  -> delete all data created by the TICK Stack
  docker-clean -> stop and remove all running docker containers
  rebuild-docs -> rebuild the documentation container to see updates
```

To get started just run `./sandbox up`. You browser will open two tabs:

- `localhost:8888` - Chronograf's address. You will use this as a management UI for the full stack
- `localhost:3000` - Documentation server. This contains a simple markdown server for tutorials and documentation.

> NOTE: Make sure to stop any existing installations of `influxdb`, `kapacitor` or `chronograf`. If you have them running the sandbox will run into port conflicts and fail to properly start. In this case stop the existing processes and run `./sandbox restart`.

To configure the connection to InfluxDB from Chronograf just fill in `http://influxdb:8086` as the URL:

![Configure Influx](./documentation/static/images/configure-influxdb.png)

Once you have configured the InfluxDB URL you should see your dashboard:

![Dashboard](./documentation/static/images/dashboard.png)

Then click on the gear icon and select `Kapacitor`:

![Kapacitor Configuration](./documentation/static/images/kapacitor-config.png)

Finally, enter `http://kapacitor:9092` as the URL and click `Connect Kapacitor`:

![Kapacitor Configuration](./documentation/static/images/configure-kapacitor.png)

Then you are ready to get started with the TICK Stack!

Visit `http:localhost:8888` and click the host to see your dashboard, then check out the tutorials at `http://localhost:3000/tutorials`.

>Note: see [influx-stress](https://github.com/influxdata/influx-stress) to create data for your sandbox. 
>

![Dashboard](./documentation/static/images/sandbox_dashboard.jpg)
