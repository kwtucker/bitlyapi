### REQUIREMENTS

This coding exercise is an opportunity for you to show us how you break down product requirements into actual code, as well as to demonstrate quality coding style, experience and creativity in problem solving. It will also introduce you to the Bitly API and expose you to some of the functionality the back-end team develops on a daily basis.

#### Bitly Access Token
* In order to make use of our API, you will need a Bitly Access Token.
  * Sign up for a Bitly account if you do not already have one.
  * Visit [this page](https://bitly.is/accesstoken) to get your Access Token.
  * Due to security concerns, do not include your access token as part of the submission. We will immediately reject any submission containing an access token!!!

#### Language and Framework
* Choose languages and frameworks you are relatively comfortable with.
* Bonus points if you can complete the problems in Go.

#### Project Setup
* Send us a runnable server with some descriptions of endpoint functionality and example requests (`curl` or similar is fine).
* Please include a short writeup in your README describing the major design decisions behind your solution.
* Make sure to list the dependencies of your project as well as how to install them (we may not be experts in your chosen language).

#### The Problem

We would like your server to expose an endpoint to provide the average number of clicks, per country, within the last 30 days, for the Bitlinks in a user's default group.

The following [Bitly API](https://dev.bitly.com/v4_documentation.html) endpoints will be helpful in your solution:
* [https://api-ssl.bitly.com/v4/user](https://dev.bitly.com/v4_documentation.html#operation/getUser) - provides user information including the users default group
* [https://api-ssl.bitly.com/v4/groups/{group_guid}/bitlinks](https://dev.bitly.com/v4_documentation.html#operation/getBitlinksByGroup) - provides paged information about the Bitlinks for a provided group
* [https://api-ssl.bitly.com/v4/bitlinks/{bitlink}/countries](https://dev.bitly.com/v4_documentation.html#operation/getMetricsForBitlinkByCountries) - user clicks, broken down by country, for a provided Bitlink

Construct an API service that exposes this data as JSON over HTTP. You can assume the API endpoint you write will take an access token as input.

Note that this differs slightly from the results returned by our endpoint at [https://api-ssl.bitly.com/v4/groups/{group_guid}/countries](https://dev.bitly.com/v4/#operation/getGroupMetricsByCountries). The format of your response data could, however, mirror the format of data returned by that endpoint.

##### Notes
* Bitly Glossary:
  * Bitlink - the short link created by the Bitly platform
  * decode/user click - a metric collected each time a Bitlink is accessed and performs the redirect
  * group - a way of organizing and controlling access to data between multiple users in an organization
  * Long URL - the destination redirect of a Bitlink
* The problem description is intentionally vague; this is a test of your ability to explore the dataset and expose a logical, reasonably performant API to the data.