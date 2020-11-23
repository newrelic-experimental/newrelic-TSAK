[![New Relic Experimental header](https://github.com/newrelic/opensource-website/raw/master/src/images/categories/Experimental.png)](https://opensource.newrelic.com/oss-category/#new-relic-experimental)


![TSAK architecture](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/documentation/images/architecture.png)


# TSAK

TSAK - as translated as a "Telemetry Swiss Army Knife", the instrument designed to aide the development of following category of tools:
* Telemetry stream converters and translators.
* On-host Telemetry acquisition and processing.
* Generic programmable data converision tool.
* Rule-Based Expert system shell, geared towards telemetry data processing.
* To be a part of telemetry processing pipeline
* To generate a simulated telemetry data for Monitoring and Observability systems development, testing, training and teaching and POC demonstrations.

The development of TSAK is ongoing with scheduled monthly releases.

Current in-development version is [0.2-pre2](https://github.com/newrelic-experimental/newrelic-TSAK/tree/0.2-pre2)

## Installation

As of current, there is no process of generating and provision of the binary packages. After TSAK binary is generated, please place it to */usr/local/bin* or other location of your desire. There is no pre-defined places of where TSAK is looking for it's scripts, so pick the place most appropriate to them, with permissions that allow user under which you will execute TSAK to read them. If you are planning to use "net/icmp" module, you may consider to elevate privileges, as use of the ICMP requires that.

To build the tool from the sources, please refer "Building" section of this document.


## Getting Started

Consider running TSAK like this:
```bash
tsak -help
```
As it is always a good start. You shall familiarize yourself with the existing command line keys. Here is a "Cheat sheet" for some command line keys that you may find most helpful:

### Logs and debugging

This group of command line parameters will get you a loggin output from the TSAK application itself and from the TSAK scripts.

Key | Description
----|------------
-debug | Turn on all logging output up to tracing
-stdout | Send logging output to stdout
-log | Path to a log file
-production | Changes TSAK settings to be tuned to a production mode.
-logage | TSAK will rotate log file when reached maximum age
-logsize | TSAK will rotate the log when reached maximum size
-tracern | If given, will send a traces to New Relic as log entries.
-error | Clip log output to errors
-warning | Clip log output to warnings
-info | Clip log output to info


### Establishing the conversion pipeline

This group of command line parameters will help you to define a conversion pipeline.

Key | Description
----|------------
-in | Path to the script which will be running at "Protocol side"
-proc | Path to the script which will be running as "Processing stage"
-out | Path to the script which will be running as a "Feeder"

### NewRelic related command line parameters

This group of command line parameters will help you to set-up communication with New Relic SaaS service

Key | Description
----|------------
-nrapi | New Relic API key for insert operations
-nrapiq | New Relic API key for query operations
-account | You New Relic account number
-evtapi | URL for Event API endpoint
-logapi | URL for Log API endpoint
-queryapi | URL for Query API endpoint
-metricapi | URL for Metric API endpoint

### Generic application command line parameters

This group of command line parameters, controlling generic behavior of the TSAK application

Key | Description
----|------------
-production | Changes TSAK settings to be tuned to a production mode.
-name | Will specify a name for the TSAK application
-id | Unique identifier for the application
-conf | Path to the TSAK-script file, which will be pre-executed to all TSAK virual machines
-run | Path to the "exclusive" TSAK script. If you are running TSAK in "exclusive mode", all other scripts in -in/-proc/-out will be ignored.
-clips | Path to the CLIPS main script. If provided, CLIPS thread is executed in "exclusive mode", and terminates TSAK after script is finished.

And this example will demonstrate you how you can run a script in exclusive mode:
```bash
tsak -stdout -production -debug  -name "helloworld" -conf ./config.example/tsak.conf -run ./examples/run/helloworld.script
```
This script will load configuration file *./config.example/tsak.conf* to all TSAK VM's, then executed TSAK script *./examples/run/helloworld.script*

```golang
fmt = import("fmt")

fmt.Println("Hello world")
fmt.Println("The answer is ", ANSWER)
fmt.Println("Not an answer is ", NO_ANSWER)
```
and the script output as expected is
```
Hello world
The answer is  42
Not an answer is  41
```

## Usage

![TRAPD architecture](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/documentation/images/snmptrapd.png)

Here is more complicated example of the TSAK use. This is a functioning SNMP trap server. It does have all three components, such as:

1. [Protocol component](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/trapd/in.script) - SNMP server which listens the UDP port, accepts and parses the packet.
2. [Processing component](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/trapd/proc.script) - Resolving OID to Symbol using MIB's
3. [Feeder component](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/trapd/out.script) - sending events to the New Relic.

Run the TSAK-powered snmptrapd in foreground:
```bash
tsak -stdout -production -nrapi (New Relic Insert API key) -account (New Relic account number) -name "trapd" -in ./examples/trapd/in.script -out ./examples/trapd/out.script -proc ./examples/trapd/proc.script
```
When application is started, you can send a simulated SNMP traps using [that](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/examples/trapd/sendatrap.sh) script

## Building

TSAK will bring all golang modules that it requires automatically during the build process. All, except one, you have to assist building *"github.com/Keysight/clipsgo"*. After you check-out most recent *clipsgo* module source code from https://github.com/Keysight/clipsgo. Check out this module to your GOPATH then cd to the root directory of the module and issue command
```bash
make clips
```
This will download and build CLIPS library. Copy libclips.a and libclips.so from *clips_source* to the directory listed in your LD_LIBRARY_PATH, for example */usr/local/lib*. Optionally, you may consider to copy a binary file *clips* from the same directory *clips_source* to any directory in your PATH, for example */usr/local/bin*. This CLIPS shell is a great learning tool to get you familiarize with CLIPS. In Linux OS, you may need to issue *ldconfig* after you copy the file. Then build your TSAK code as usual. Dependency on external *libclips* shared library is only external dependency as of now.

The build process of the TSAK itelf are very straightforward. Just check out the source code, change directory to the root of the TSAK source code and run *make*

```bash
➜  newrelic-TSAK git:(main) ✗ make
=== newrelic-TSAK === [ deps             ]: Installing package dependencies required by the project...
=== newrelic-TSAK === [ tools-compile    ]: building tools:
=== newrelic-TSAK === [ tools            ]: Installing tools required by the project...
go: finding module for package github.com/goreleaser/goreleaser
go: found github.com/goreleaser/goreleaser in github.com/goreleaser/goreleaser v0.147.2
=== newrelic-TSAK === [ compile          ]: building commands:
=== newrelic-TSAK === [ compile          ]:     ./bin/darwin/tsak
```

## Getting started with CLIPS Rule-Based expert system

TSAK provide you with a programmatic access to an Expert System shell from your scripts. You can convert telemetry to facts, load facts to an Expert System, run the rules and exporting facts to the pipeline converted to a JSON.

If you never worked with CLIPS, start with study of the standard CLIPS documentations:

1. [User Guide](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/documentation/clips/clips_ug.pdf) - the very first steps in understanding on how to creat and operate of an Expert Sytem shell. What are the facts, rules, how to create and run your processing.
2. [Basic Programming Guide](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/documentation/clips/clips_bpg.pdf) - the next step after you done with [User Guide](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/documentation/clips/clips_ug.pdf). [Basic Programming Guide](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/documentation/clips/clips_bpg.pdf) is an excellent source of the in-depth information on CLIPS shell programming.

## Continue to Learn

First, there are [exhaustive list](https://github.com/newrelic-experimental/newrelic-TSAK/tree/main/examples) of TSAK-script samples, bundled with TSAK source code. [Here is a description](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/documentation/example.md) of what some of those scripts are demonstrating.

Next, consider to go through the ["Telemetry Swiss Army Knife in 10 minutes"](https://github.com/newrelic-experimental/newrelic-TSAK/blob/main/documentation/presentations/Telemetry%20Swiss%20Army%20Knife%20in%2010%20minutes/Telemetry%20Swiss%20Army%20Knife%20in%2010%20minutes.pdf) presentation deck.

Link to a proper Documentation will be published when we get the documenting of the TSAK going.

## Support


New Relic has open-sourced this project. This project is provided AS-IS WITHOUT WARRANTY OR DEDICATED SUPPORT. Issues and contributions should be reported to the project here on GitHub.

We encourage you to bring your experiences and questions to the [Explorers Hub](https://discuss.newrelic.com) where our community members collaborate on solutions and new ideas.

You can always use an "Issues" tab of the GitHub project to open a support ticket or report an issue.

https://github.com/newrelic-experimental/newrelic-TSAK/issues

## Contributing

We encourage your contributions to improve TSAK! Keep in mind when you submit your pull request, you'll need to sign the CLA via the click-through using CLA-Assistant. You only have to sign the CLA one time per project. If you have any questions, or to execute our corporate CLA, required if your contribution is on behalf of a company, please drop us an email at opensource@newrelic.com.

**A note about vulnerabilities**

As noted in our [security policy](../../security/policy), New Relic is committed to the privacy and security of our customers and their data. We believe that providing coordinated disclosure by security researchers and engaging with the security community are important means to achieve our security goals.

If you believe you have found a security vulnerability in this project or any of New Relic's products or websites, we welcome and greatly appreciate you reporting it to New Relic through [HackerOne](https://hackerone.com/newrelic).

## License

TSAK is licensed under the [Apache 2.0](http://apache.org/licenses/LICENSE-2.0.txt) License.

TSAK also uses source code from third-party libraries. You can find full details on which libraries are used and the terms under which they are licensed in the third-party notices document.
