# Using the Telegraf socket_listener Plugin

The Telegraf [socket_listener](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/socket_listener) plugin allows you to send arbitrary measurements to Telegraf using UDP or TCP. It is often the fastest way to get your metrics into InfluxDB, as it will parse various [inputs](https://docs.influxdata.com/telegraf/v1.5/concepts/data_formats_input/) and take care of things like automatic retries for you.

In order to use the Telegraf socket_listener Plugin with the InfluxData Sandbox, we will need to expose a port for the Telegraf container to receive data. We can do that by adding a `ports` section under `telegraf` in `./docker-compose.yml`.

It should look like this:

```
telegraf:
    # Full tag list: https://hub.docker.com/r/library/telegraf/tags/
    image: telegraf:1.5.3
    environment:
        HOSTNAME: "telegraf-getting-started"
    # Telegraf requires network access to InfluxDB
    links:
        - influxdb
    volumes:
        # Mount for telegraf configuration
        - ./telegraf/:/etc/telegraf/
        # Mount for Docker API access
        - /var/run/docker.sock:/var/run/docker.sock
    ports:
        # The socket_listener plugin listens or port 8094
        - "8094:8094/udp"
    depends_on:
        - influxdb
```

We'll also need to configure the Telegraf plugin to receive UDP packets. Open `./telegraf/telegraf.conf` in your editor and add the following lines at the end of the file:

```
[[inputs.socket_listener]]
  service_address = "udp://:8094"
  data_format = "influx"
```

Once you've made those changes you can apply them by restarting the Sandbox:

```
$ ./sandbox restart
Stopping all processes...
Starting all processes...
Services available!
$
```

Now let's test that Telegraf can receive data via UDP. We'll send a metric using the InfluxDB [Line Protocol](https://docs.influxdata.com/influxdb/v1.5/write_protocols/line_protocol_tutorial/), which Telegraf understands. To send the data, we'll use two common command line utilities, `echo` and `nc`, or netcat.

Both of these utilites are included on macOS and most Linux distributions; if you're using Windows, you can [install Ubuntu from the Windows Store](https://www.microsoft.com/store/productId/9NBLGGH4MSV6) and use these tools from within the Windows Subsystem for Linux (WSL).

On macOS, enter the following commands:

```
$ echo "my_measurement,my_tag_key=my_tag_value value=1" | nc -u -4 -w 1 localhost 8094
$ echo "my_measurement,my_tag_key=my_tag_value value=2" | nc -u -4 -w 1 localhost 8094
$ echo "my_measurement,my_tag_key=my_tag_value value=3" | nc -u -4 -w 1 localhost 8094
$ echo "my_measurement,my_tag_key=my_tag_value value=4" | nc -u -4 -w 1 localhost 8094
$ echo "my_measurement,my_tag_key=my_tag_value value=5" | nc -u -4 -w 1 localhost 8094
```

The `echo` command prints the text within quotes, which in this case is data encoded using the InfluxDB line protocol. Next, we use  the `|` character to "pipe" the text into the next command, `nc`. We provide several arguments to `nc`: `-u`, `-4`, and `-w 1`. The first argument tells `nc` to send data using UDP; the second argument says that we should use IPv4, and the third argument tells `nc` to wait one second before terminating the connection.

On Linux or WSL you can substitute `-q` for `-w 1`, which tells `nc` to quit once the data has been sent.

Telegraf has the ability to aggregate and process data it receives before sending it on using an output plugin. In the Sandbox, Telegraf is configured with a five second "flush" interval, which means that Telegraf will collect measurements for five seconds, aggregate them, and send the resulting value to InfluxDB. Since we have not specified the type of aggregation we'd like, Telegraf will default to computing the mean of the measurements received during the interval.

Now that we've sent some measurements to Telegraf, let's verify that they made their way into the InfluxDB database. Open Chronograf at [http://localhost:8888](http://localhost:8888) and navigate to the data explorer.

Select the `telegraf.autogen` database, then the `my_measurement` field. You should see a graph with some values!

Because of the aggregation, the values on the graph likely will not match the values you entered at the command line unless you waited more than five seconds between each command. Play around with sending different values in different intervals to Telegraf to better understand how the aggregates are computed!
