# Get Started with Flux
Flux is InfluxData's new data language designed for querying, analyzing, and acting on data stored in InfluxDB.
Its takes the power of InfluxQL and the TICKscript and combines them into a single, unified data scripting language.

## Using Flux with the Sandbox
Starting with the 1.7 releases of Chrongoraf and InfluxDB, Flux is now completely integrated within 
the [InfluxData Sandbox](https://github.com/influxdata/sandbox).

There will be ongoing updates and refinements made to the language and the implementation.  So, you can always grab the
latest by starting the Sandbox with the `-nightly` flag to pull the nightly builds of InfluxDB and
Chronograf.

```bash
./sandbox up -nightly
```

### Flux via CLI
Integrated within the InfluxCLI is a Flux Read-Eval-Print-Loop (REPL) Command Line Interface.  

To access the Flux REPL:
```
// Enter the docker container containing the InfluxCLI from the command prompt
$ ./sandbox enter influxdb  

// Once inside the container, start the InfluxCLI tool using the -type flux option
# influx -type flux

//The following should appear
Connected to http://localhost:8086 version 1.7.0
InfluxDB shell version: 1.7.0
Enter a Flux query
>
```

__Remember to use CTL+D to exit the Flux REPL.__

### Get started with the Flux Editor via Chronograf
The Flux Editor makes working with Flux a visual process. It consists of 3 panes:

1. **[The Schema Explorer](#schema-explorer)** Allows you to explore the actual structure of your data as you're building Flux scripts.
2. **[The Script Editor](#script-editor)** Where the actual Flux code is written and displayed.
3. **[The Function Explorer](#function-explorer)** An online quick reference for many Flux functions.

![Flux Editor](../images/flux-editor.png)

Each pane can be minimized, maximized, or closed depending on how you want to work.

### Schema Explorer
The "Explore" pane of the Flux Editor allows you to visual explore the structure of your data.
This is incredibly helpful as you're building out Flux queries.

![Schema Explorer](../images/flux-editor-explore.png)

### Script Editor
Flux queries are written in the "Script" pane of the Flux Editor.
You can also use various 'helper tools' around the Script Editor to rapidly build out queries.
As queries are updated in the Flex Builder, the are updated in the script editor.

![Script Editor](../images/flux-editor-script.png)

### Function Explorer
The "Build" pane is a visual representation of your Flux script that used to both visualize and build your script.
As queries are updated in the Flex Builder, the are updated in the script editor.
At any point in your Flux query, you can use the `yield()` function to visualize the current state of your query.

![Flux Builder](../images/flux-function-explorer.png)

## Learn the basics of the Flux language
Flux draws inspiration from programming languages such as Lisp, Elixir, Elm,
Javascript and others, but is specifically designed for analyzing and acting on data.
For an introduction into the Flux syntax, view the
[Flux Getting Started](https://docs.influxdata.com/flux/latest/introduction/getting-started/)
section of the documentation.

You can also [explore a walkthrough of the language including a handful of simple expressions and what they mean.](https://github.com/influxdata/platform/blob/nc-training/TRAIN.md#learning-flux)

## Additional Information
[Flux Documentation](https://docs.influxdata.com/flux/latest/introduction/getting-started/)  
[Flux Specification](https://github.com/influxdata/platform/blob/master/query/docs/SPEC.md)
[Flux Introduction Slides](https://speakerdeck.com/pauldix/flux-number-fluxlang-a-new-time-series-data-scripting-language) 
