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

## Installation

As of current, there is no process of generating and provision of the binary packages. After TSAK binary is generated, please place it to */usr/local/bin* or other location of your desire. There is no pre-defined places of where TSAK is looking for it's scripts, so pick the place most appropriate to them, with permissions that allow user under which you will execute TSAK to read them. If you are planning to use "net/icmp" module, you may consider to elevate privileges, as use of the ICMP requires that.

To build the tool from the sources, please refer "Building" section of this document.


## Getting Started

```bash
tsak -help
```
Is always a good start. Please familiarize yourself with the existing command line keys. Here is a "Cheat sheet" for some command line keys that you may find most helpful:

### Logs and debugging

Key | Description
----|------------
-debug | Turn on all logging output up to tracing
-stdout | Send logging output to stdout
-log | Path to a log file
-production | Changes TSAK settings to be tuned to a production mode. For example, switch log format from text to json

## Usage

>[**Optional** - Include more thorough instructions on how to use the software. This section might not be needed if the Getting Started section is enough. Remove this section if it's not needed.]

## Building

TSAK will bring all golang modules that it requires automatically during the build process. All, except one, you have to assist building *"github.com/Keysight/clipsgo"*. After you check-out most recent *clipsgo* module source code from https://github.com/Keysight/clipsgo. Check out this module to your GOPATH then cd to the root directory of the module and issue command
```bash
make clips
```
This will download and build CLIPS library. Copy libclips.a and libclips.so from *clips_source* to the directory listed in your LD_LIBRARY_PATH, for example */usr/local/lib*. Optionally, you may consider to copy a binary *clips* from the same directory *clips_source* to any directory in your PATH, for example */usr/local/bin*. This CLIPS shell is a great learning tool to get you familiarize with CLIPS. In Linux OS, you may need to issue *ldconfig* after you copy the file. Then build your TSAK code as usual. Dependency on external *libclips* shared library is only external dependency as of now.


## Support

New Relic has open-sourced this project. This project is provided AS-IS WITHOUT WARRANTY OR DEDICATED SUPPORT. Issues and contributions should be reported to the project here on GitHub.

We encourage you to bring your experiences and questions to the [Explorers Hub](https://discuss.newrelic.com) where our community members collaborate on solutions and new ideas.

You can always use an "Issues" tab of the GitHub project to open a support ticket or report an issue.

https://github.com/newrelic-experimental/newrelic-TSAK/issues

## Contributing

We encourage your contributions to improve Salesforce Commerce Cloud for New Relic Browser! Keep in mind when you submit your pull request, you'll need to sign the CLA via the click-through using CLA-Assistant. You only have to sign the CLA one time per project. If you have any questions, or to execute our corporate CLA, required if your contribution is on behalf of a company, please drop us an email at opensource@newrelic.com.

**A note about vulnerabilities**

As noted in our [security policy](../../security/policy), New Relic is committed to the privacy and security of our customers and their data. We believe that providing coordinated disclosure by security researchers and engaging with the security community are important means to achieve our security goals.

If you believe you have found a security vulnerability in this project or any of New Relic's products or websites, we welcome and greatly appreciate you reporting it to New Relic through [HackerOne](https://hackerone.com/newrelic).

## License

TSAK is licensed under the [Apache 2.0](http://apache.org/licenses/LICENSE-2.0.txt) License.

TSAK also uses source code from third-party libraries. You can find full details on which libraries are used and the terms under which they are licensed in the third-party notices document.
