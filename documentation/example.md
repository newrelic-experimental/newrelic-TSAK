![TSAK architecture](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/documentation/images/architecture.png)

# Here, you will find a description of the some of sample TSAK scripts

Nothing demonstrates system capabilities better, than a collection of the sample scripts and solutions. We are striving to bring as much examples, which demonstrates TSAK system, as possible. Most of those scripts are designed to be running in "exclusive run" mode, when they are passed as a value to a *-run* command line parameter.

## Basic scripts

This scripts is pretty much showing you the "Hello world" type of applications.

Script | Description
----|------------
[helloworld.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/helloworld.script) | A true HelloWorld script, also demonstrating that the values imported through configuration file, available from inside the TSAK script.
[log1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/log1.script) | Demonstration of how you can send a messages to the a logging.
[loop.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/loop.script) | All your -in/-proc/-out script logic, must be created inside the loop that checking if exit from the TSAK application is requested
[cron1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/cron1.script) | Yes, there are 3 basic threads inside TSAk application, but nothing says that you can not run your own periodic background scripts. There is internal "cron" in TSAK tool and you can create TSAK script functions and run them periodically.

## How to create a -in/-proc/-out pipeline

The core of the TSAK application is a in-line processing of the data, which flowing from "IN" to "PROC" and then to "OUT". Here is the demo of a simple pipeline.

Script | Description
----|------------
[loop_in1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/in/loop_in1.script) | Generate and send to a PROC thread a test JSON data with counter every second
[loop_proc1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/proc/loop_proc1.script) | Processor, receiving data from the channel, prints and send to "OUT"
[loop_out1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/out/loop_out1.script) | Receives the data from Processor and prints them.


## New Relic related scripts

TSAK tool is tightly integrated with New Relic SaaS services. You can transparently send Logs, Events, Metrics to the New Relic platform from your TSAk scripts.

Script | Description
----|------------
[out.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/trapd/out.script) | How to send a pre-formed events to a New Relic with one line of the code
[query1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/query1.script) | Send an NRQL query and parse result

## Network and SNMP scripts

Those scripts demonstrates on how you can interact with network services

Script | Description
----|------------
[trapd.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/trapd.script) | Listening for SNMP trap packets and prinis VarBind from received packets
[mibparse2.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/mibparse2.script) | Loads all MIB files and resove OID to Symbol and Symbol to OID.
[ping1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/ping1.script) | Sends ICMP packet to destination host and handles the ICMP Echo, returning RTT time. So, basically your common *ping* available from the TSAK script. Note, you have to be _root_, to run this script.


## Data management and AI scripts

Working with and generating the data is one of the main "reasons for existence" for the TSAK tool

Script | Description
----|------------
[djson.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/djson.script) | Demonstrates on how you can create, and parse JSON data
[input1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/input1.script) | Show you, how you can read data from STDIN and generate JSON with that data
[stat1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/stat1.script) | How to do a statistical computations.
[aimsimple1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/aimsimple1.script) | Create and train a Neural Net, then find the patterns in your telemetry samples.



## CLIPS examples

How you can interact with Rule-Based expert system from your TSAK scripts.

Script | Description
----|------------
[clips5.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/clips5.script) | Demonstrates you how you can push the facts and commands through different pipelines from the script to CLIPS
[clips6.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/clips6.script) | How to submin the facts, load rules and run the rules processing.
[testmain1.clips](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/clips/testmain1.clips) | That is what is your "exclusive main" CLIPS script may look like.
