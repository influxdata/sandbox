# Understanding the Sandbox

The InfluxData Sandbox runs a complete InfluxData setup and provides a convenient way to learn to use all of the products in concert through the Chronograf management user interface. It also contains a learning portal with links to relevant pages. The Sandbox is built with [Docker](https://www.docker.com/).

### Configuration changes

To change the configuration for any of the products, just change the config file in the respective directory and restart the Sandbox:

```bash
$ vi ./influxdb/config/influxdb.conf
# Make some changes...
$ ./sandbox restart
# terminal output...
```

### Data

After initial startup you will see the data directories for `influxdb`, `chronograf` and `kapacitor` created. This is where the data for `sandbox` is persisted.

```bash
.
├── chronograf
│   └── data
├── influxdb
│   ├── config
│   └── data
├── kapacitor
│   ├── config
│   └── data
└── telegraf
    └── telegraf.conf
```

### InfluxDB

In the `sandbox` InfluxDB is collecting data that is created by `telegraf`, forwarding it to `kapacitor` for alerting, and serving dashboard queries from `chronograf`. The API is available at `http://localhost:8086`

If you have an existing installation of InfluxDB on your computer, you can use the `influx` cli tool to run commands against the `sandbox` instances without additional configuration. If you do not have the tools then use the Sandbox to attach to the CLI:

```bash
$ ./sandbox influxdb
Entering the influx cli...
Connected to http://localhost:8086 version 1.2.2
InfluxDB shell version: 1.2.2
> SHOW DATABASES
name: databases
name
----
telegraf
_internal
```

### Kapacitor

In the `sandbox` Kapacitor catches and processes data coming in from InfluxDB and can have tasks created on it by Chronograf. The API is available at `http://localhost:9092`

While a downloaded copy of the `kapacitor` cli will work to run commands against the Sandbox instance, you will need to use the cli in the container if you do not have it locally:

```bash
$ ./sandbox enter kapacitor
Entering /bin/bash session in the kapacitor container...
root@d5d99840dbc7:/# kapacitor stats ingress
Database   Retention Policy Measurement             Points Received
_internal  monitor          cq                                  146
_internal  monitor          database                            292
_internal  monitor          httpd                               146
_internal  monitor          queryExecutor                       146
....
```

### Chronograf

In the `sandbox` Chronograf acts as the control hub for the different products. It can query data from and perform management function for InfluxDB and create tasks on Kapacitor. The Chronograf web interface is available at `http://localhost:8888`.

### Telegraf

In the `sandbox` Telegraf is collecting data. It has the following plugins configured by default:

* [`outputs.influxdb`](https://github.com/influxdata/telegraf/tree/master/plugins/outputs/influxdb)
* [`inputs.docker`](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/docker)
* [`inputs.cpu`](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/system)
* [`inputs.system`](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/system)
* [`inputs.influxdb`](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/influxdb)
