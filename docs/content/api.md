# Table of contents

* [**Grpc API**](#grpc-api)
    * [**Authentication**](#authentication)
    * [**Resource**](#resource)
    * [**Record**](#record)
    * [**Namespace**](#namespace)
    * [**Datasource**](#data-source)
    * [**User**](#user)
    * [**Extension**](#extension)
    * [**ExternalCall**](#external-call)
* [**Rest API**](#rest-api)
    * [**Swagger**](#swagger)
* [**SDK and code generation**](#sdk)
    * [**Golang**](#golang)
    * [**Python**](#python)
    * [**Nodejs**](#nodejs)
* [**CLI**](#cli)

# Grpc Api

Grpc is main interface of Data Handler. By defaul 9009 port is used to communicate to Data handler grpc service.

Grpc services:

* [**Authentication**](#authentication) - Authentication
* [**Resource**](#resource) - Resource operations
* [**Record**](#record) - Record operations (CRUD)
* [**Namespace**](#namespace) - Namespace operations (CRUD)
* [**Datasource**](#data-source) - Data Source operations (CRUD)
* [**User**](#user) - User operations (CRUD)
* [**Extension**](#extension) - Extension operations (CRUD)

## Resource

Proto file: [resource.proto](https://github.com/tislib/data-handler/blob/master/proto/stub/resource.proto)

Resource service is for managing resources inside Data handler.
For detailed information about resource
see [link](general.md#resource)

Methods:
* [**create**]  

* [**ExternalCall**](#external-call) - External call service is for extension purposes and it is not available by
  Datahandler itself, instead, External call service is required to be implemented on extension service, So data handler
  is communicating with extension service through this service