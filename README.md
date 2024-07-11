### Section 3: Testing REST API Integrations
https://github.com/federicoleon/golang-restclient

https://api.mercadolibre.com/countries/AR


#### RESTCLIENT
**restclient** is a Go library that simplifies making HTTP requests, especially for RESTful APIs. It abstracts away much of the boilerplate code required to make HTTP calls, handle responses, and manage errors, making it easier to work with external APIs.

In this lesson, we're using the **github.com/federicoleon/golang-restclient** package. This package is commonly used for making REST API calls and provides functionalities for **MOCKING** HTTP requests and responses, which is particularly useful for unit testing.

**Notice that:**
If the mockup server is started using rest.StartMockupServer() but no mock responses are added using rest.AddMockups(), the request will not be sent to the real server. Instead, the mockup server will handle the request and return a default response indicating that no mock was found for the request.

When the mockup server is running, it intercepts all HTTP requests made using the rest client. If a request is made and there is no corresponding mockup configured using rest.AddMockups(), the mockup server will not forward the request to the actual API endpoint. Instead, it typically returns a response indicating that the mockup was not found.


**Tools**
https://codebeautify.org/jsonminifier (reformat json into a string for testing http response)


### Section 4: Testing the whole app
#### Gin-gonic
https://github.com/gin-gonic/gin
