![banner]
# What api-hub does?
`api-hub` is very simple repository for storing, serving and invoking REST API.

The project is designed to solve the following problems:
* There is a small-to-mid-size project designed using Micro-Services that communicates via REST and documented using Swagger / OpenAPI.
* As developer I want to have the documentation of all services in a single place, where I can find them easily.
* A DevOps must be able to configure the pipelines easily to upload the API specification on build.
* DevOps must be able to deploy the application easily in the cloud or any other machine, that supports Docker.
* As developer, I would like to be able to invoke some API manually and Swagger-UI and/or Redoc is enough.

# What api-hub does not do?

* Security is not implemented at all. All get/post/delete methods are completely unprotected. So don't install in in a public place.
* Validation/Linting of the published API. Lint rules may vary, so it's up to developers to validate their specification before publishing it. Here are some tools that might do the trick:
 * https://github.com/Redocly/openapi-cli
 * https://github.com/wework/speccy
 * https://github.com/stoplightio/spectral
 * https://github.com/whitlockjc/oval
 * https://github.com/paypal/openapilint
 * https://hub.docker.com/r/shipchain/openapi-spellcheck
* Conversion - if you want to support multiple API formats, look at https://github.com/lucybot/api-spec-converter
* Write the API for you! :)
 * If you prefer to write the specification first - you can start with https://github.com/Redocly/openapi-template and the use https://openapi-generator.tech/ to generate you server stubs/clients.
 * Or you can generate swagger/openapi directly from your code - check what tools are supported for your programming language.


# Alternatives

Speaking of visualisation and documenting the REST API, the following alternatives has also been considered, but not implemented:
* https://github.com/LucyBot-Inc/documentation-starter
* https://github.com/slatedocs/slate

This project looks a lot like APIs.guru. It is a great tool, if you want to publish your API publicly.
* https://apis.guru/
* https://github.com/APIs-guru/openapi-directory


# What more can be done?

Automatic code samples generation for Redoc. There are tools, that can generate some code snippets, but apparently, they are not build into Redoc. Using those tools the published specification may be updated with some code examples:
* https://github.com/ErikWittern/openapi-snippet
* https://github.com/Redocly/openapi-sampler

Metrics is something else that is worth adding. It might be useful to know which API are used/browsed/downloaded most.

CORS Proxy like `cors-anywhere`.
