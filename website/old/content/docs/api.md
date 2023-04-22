---
layout: default
title: Api
nav_order: 2
has_children: true
---

# Grpc Api

Grpc is main interface of API Brew. By default, 9009 port is used to communicate to API Brew grpc service.

Grpc services:

* [**Authentication**](proto.md#authentication-service) - Authentication
* [**Resource**](proto.md#resource-service) - Resource operations
* [**Record**](proto.md#record-service) - Record operations (CRUD)
* [**Namespace**](proto.md#namespace-service) - Namespace operations (CRUD)
* [**Datasource**](proto.md#datasource-service) - Data Source operations (CRUD)
* [**User**](proto.md#user-service) - User operations (CRUD)
* [**Extension**](proto.md#extension-service) - Extension operations (CRUD)

# Rest Api

You can see Rest api docs on swagger, by default swagger is available on http://localhost:9009/docs/index.html

Rest services:
* [**Authentication**](openapi.md#authentication) - Authentication
* [**Resource**](openapi.md#resourceget) - Resource operations
* [**Record**](openapi.md#recordget) - Record operations (CRUD)
* [**Namespace**](openapi.md#namespaceget) - Namespace operations (CRUD)
* [**Datasource**](openapi.md#datasourceget) - Data Source operations (CRUD)
* [**User**](openapi.md#userget) - User operations (CRUD)
* [**Extension**](openapi.md#extensionget) - Extension operations (CRUD)
