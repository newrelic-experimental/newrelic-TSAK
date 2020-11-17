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

## New Relic related scripts

## Network and SNMP scripts

Script | Description
----|------------
[trapd.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/trapd.script) | Listening for SNMP trap packets and prinis VarBind from received packets
[mobparse2.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/mibparse2.script) | Loads all MIB files and resove OID to Symbol and Symbol to OID.


## Data management and AI scripts

Script | Description
----|------------
[djson.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/djson.script) | Demonstrates on how you can create, and parse JSON data
[input1.script](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/run/input1.script) | Show you, how you can read data from STDIN and generate JSON with that data


## CLIPS examples
