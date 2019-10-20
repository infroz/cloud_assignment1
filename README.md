# Assignment 1 for Cloud Technologies at NTNU Gjøvik
## Task Description
In this assignment, you are going to develop a REST web application in Golang that provides the client to retrieve information about animal species. For this purpose, you will interrogate an existing web service and return the result in a given output format.

The REST web service you will be using is the [Global Biodiversity Information Facility](https://www.gbif.org/what-is-gbif), GBIF. It is based on open standards and open access, and it allows the client to query various aspects of Earth's biodiversity.

[The documentation of GBIF is available online](https://www.gbif.org/developer/summary). To use the GBIF web service through GET requests you do **not** need an API token! Nevertheless, be **mindful of rate limits** - we will talk about mitigation strategies in class.

In addition to the information provided by the GBIF database, you will need to use the [restcountries.eu](https://restcountries.eu) API for obtaining information about countries and their capitals, country codes, currency, etc.

The final web service should be deployed on Heroku. The initial development should occur on your local machine. For the submission, you will both need to provide a URL to the deployed Heroku service as well as your repository.
## Heroku
This application has been deployed on Heroku 
Link to Heroku [https://lit-garden-19973.herokuapp.com/](https://lit-garden-19973.herokuapp.com/ "https://lit-garden-19973.herokuapp.com/")

## How to use the API

host:port/conservation/v1/
												
Method GET
 - country/{Alpha2Code for a country}
	 - Returns data on a given country in JSON format
 - diag/
	 - Returns status code on API's used by this Rest application. Also gives uptime in seconds JSON format
 - species/
	 - Returns data on a given specieskey JSON formt

