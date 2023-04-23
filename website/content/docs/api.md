---
layout: default
title: Api
nav_order: 2
has_children: true
---

# Grpc Api

Grpc is main interface of API Brew. By default, 9009 port is used to communicate to API Brew grpc service.

Grpc services:

* [**Authentication**](/docs/proto#authentication-service) - Authentication
* [**Resource**](/docs/proto#resource-service) - Resource operations
* [**Record**](/docs/proto#record-service) - Record operations (CRUD)
* [**Namespace**](/docs/proto#namespace-service) - Namespace operations (CRUD)
* [**Datasource**](/docs/proto#datasource-service) - Data Source operations (CRUD)
* [**User**](/docs/proto#user-service) - User operations (CRUD)
* [**Extension**](/docs/proto#extension-service) - Extension operations (CRUD)

# Rest Api

You can see Rest api docs on swagger, by default swagger is available on http://localhost:9009/docs/index.html

Rest services:
* [**Authentication**](/docs/openapi#authentication) - Authentication
* [**Resource**](/docs/openapi#resourceget) - Resource operations
* [**Record**](/docs/openapi#recordget) - Record operations (CRUD)
* [**Namespace**](/docs/openapi#namespaceget) - Namespace operations (CRUD)
* [**Datasource**](/docs/openapi#datasourceget) - Data Source operations (CRUD)
* [**User**](/docs/openapi#userget) - User operations (CRUD)
* [**Extension**](/docs/openapi#extensionget) - Extension operations (CRUD)
