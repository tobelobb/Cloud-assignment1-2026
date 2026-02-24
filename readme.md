Assignment 1
countryinfo API

combines restcountries api with currency api to make add some features and give selected info

Endpoints are:
/countryinfo/v1/status/
gives back status codes if the api is up and running
/countryinfo/v1/info/{country code}
returns information about a given country based on the country code
/countryinfo/v1/exchange/{country code}
gives exchangerates to neigboring countries


main recieves initial and sends the request for a handler for each endpoint

handlers contains handlers for each idividual endpoint + root

models contains structs used for data handling with json

USE of AI
AI was consultet on how suggestions on how to structure the project
rewrite poor code or non functional code
formatting of json responses
