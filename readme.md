Assignment 1
countryinfo API

combines restcountries api with currency api to make add some features and give selected info

Endpoints are:
/countryinfo/v1/status/<br>
gives back status codes if the api is up and running<br>
/countryinfo/v1/info/{country code}<br>
returns information about a given country based on the country code<br>
/countryinfo/v1/exchange/{country code}<br>
gives exchangerates to neigboring countries<br>
<br>
<br>
main recieves initial and sends the request for a handler for each endpoint<br>
<br>
handlers contains handlers for each idividual endpoint + root<br>
<br>
models contains structs used for data handling with json<br>
<br>
USE of AI<br>
AI was consultet on how suggestions on how to structure the project<br>
rewrite poor code or non functional code<br>
formatting of json responses<br>
write the segments marked with a made with AI comment above<br>
