[![New Relic Experimental header](https://github.com/newrelic/opensource-website/raw/master/src/images/categories/Experimental.png)](https://opensource.newrelic.com/oss-category/#new-relic-experimental)

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

>[Simple steps to start working with the software similar to a "Hello World"]

## Usage

>[**Optional** - Include more thorough instructions on how to use the software. This section might not be needed if the Getting Started section is enough. Remove this section if it's not needed.]

## Building

>[**Optional** - Include this section if users will need to follow specific instructions to build the software from source. Be sure to include any third party build dependencies that need to be installed separately. Remove this section if it's not needed.]

## Testing

>[**Optional** - Include instructions on how to run tests if we include tests with the codebase. Remove this section if it's not needed.]

## Support

New Relic has open-sourced this project. This project is provided AS-IS WITHOUT WARRANTY OR DEDICATED SUPPORT. Issues and contributions should be reported to the project here on GitHub.

>[Choose 1 of the 2 options below for Support details, and remove the other one.]

>[Option 1 - no specific thread in Community]
>We encourage you to bring your experiences and questions to the [Explorers Hub](https://discuss.newrelic.com) where our community members collaborate on solutions and new ideas.

>[Option 2 - thread in Community]
>New Relic hosts and moderates an online forum where customers can interact with New Relic employees as well as other customers to get help and share best practices. Like all official New Relic open source projects, there's a related Community topic in the New Relic Explorers Hub.
>You can find this project's topic/threads here: [URL for Community thread]

## Contributing

We encourage your contributions to improve Salesforce Commerce Cloud for New Relic Browser! Keep in mind when you submit your pull request, you'll need to sign the CLA via the click-through using CLA-Assistant. You only have to sign the CLA one time per project. If you have any questions, or to execute our corporate CLA, required if your contribution is on behalf of a company, please drop us an email at opensource@newrelic.com.

**A note about vulnerabilities**

As noted in our [security policy](../../security/policy), New Relic is committed to the privacy and security of our customers and their data. We believe that providing coordinated disclosure by security researchers and engaging with the security community are important means to achieve our security goals.

If you believe you have found a security vulnerability in this project or any of New Relic's products or websites, we welcome and greatly appreciate you reporting it to New Relic through [HackerOne](https://hackerone.com/newrelic).

## License

[Project Name] is licensed under the [Apache 2.0](http://apache.org/licenses/LICENSE-2.0.txt) License.

>[If applicable: [Project Name] also uses source code from third-party libraries. You can find full details on which libraries are used and the terms under which they are licensed in the third-party notices document.]
