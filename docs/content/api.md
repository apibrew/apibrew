
# Grpc Api

Grpc is main interface of Data Handler. By default, 9009 port is used to communicate to Data handler grpc service.

Grpc services:

* [**Authentication**](proto#authentication-service) - Authentication
* [**Resource**](proto#resource-service) - Resource operations
* [**Record**](proto#record-service) - Record operations (CRUD)
* [**Namespace**](proto#namespace-service) - Namespace operations (CRUD)
* [**Datasource**](proto#datasource-service) - Data Source operations (CRUD)
* [**User**](proto#user-service) - User operations (CRUD)
* [**Extension**](proto#extension-service) - Extension operations (CRUD)

# Rest Api

You can see Rest api docs on swagger, by default swagger is available on http://localhost:9009/docs/index.html

Rest services:
* [**Authentication**](openapi.md#authentication) - Authentication
* [**Resource**](openapi#resourceget) - Resource operations
* [**Record**](openapi#recordget) - Record operations (CRUD)
* [**Namespace**](openapi#namespaceget) - Namespace operations (CRUD)
* [**Datasource**](openapi#datasourceget) - Data Source operations (CRUD)
* [**User**](openapi#userget) - User operations (CRUD)
* [**Extension**](openapi#extensionget) - Extension operations (CRUD)
