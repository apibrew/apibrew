<p align="center">
<img src="https://apibrew.io/ApiBrew%20Logo.svg" width="500px">
</p>

# Api Brew    -    https://apibrew.io
[![build](https://github.com/apibrew/apibrew/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/apibrew/apibrew/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/apibrew/apibrew)](https://goreportcard.com/report/github.com/apibrew/apibrew)
[![Go Reference](https://pkg.go.dev/badge/github.com/apibrew/apibrew.svg)](https://pkg.go.dev/github.com/apibrew/apibrew)
[![Docker Pulls](https://img.shields.io/docker/pulls/tislib/apibrew)](https://hub.docker.com/r/tislib/apibrew)
[![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/tislib/apibrew)](https://hub.docker.com/r/tislib/apibrew)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/apibrew/apibrew)
![GitHub](https://img.shields.io/github/license/apibrew/apibrew)

Docs: https://apibrew.io/docs/getting-started


## Introduction

### Run your instance

```shell
docker run tislib/apibrew:full-latest -v data:/var/lib/postgresql/data -p 9009:9009
```

**Everything is a *Resource***. And **Everything has a *CRUD API***

With API Brew, you can create CRUD APIs for your data in a few minutes

### Create first resource

country.yml
```yaml
type: resource
name: Book
properties:
  title:
    type: STRING
    unique: true
    required: true
  description:
    type: STRING
```

```bash
apbr apply -f country.yml
```
So you are ready, you have fully established Rest API for book resource

<img src="https://apibrew.io/files/book-swagger.png" width="300"/>

You can build entire application with resources and references between them (like relations in relational databases)

### Change its behaviour with power of nano code

Everything can be written by resources, not?
**Let's extend our Book resource with help of nano code**

BookLogic.js
```javascript
const book = resource('Book')

book.beforeCreate((book) => {
  if (!book.description) {
    book.description = 'No description'
  }
});
```
```bash
apbr deploy -f BookLogic.js --override
```

See the docs for nano: https://apibrew.io/docs/nano

So we have extended our book resource with help of nano code

So, with **API Brew**, you can create your application with resources and you can customize behavior of your resources with extensions

## About
API Brew is a **Low Code software** that allows to create various Grpc and Rest APIs from various database platforms

## Features

* ***Declarative*** - *API Brew* is declarative. You can define your schema in a declarative way, it will create your APIs
* ***Low Code*** - With API Brew, you can create APIs for your data without coding. But you can also extend your APIs with
  extensions, so you can customize behavior of your Resources/Apis
* ***Rest API*** - As you create resources, Rest Apis for them is made automatically
* ***Grpc*** - As you create resources, Grpc Apis for them is made automatically
* ***Database agnostic*** - API Brew is using Postgresql database by default, but it also supports various databases. Including Mongo, Mysql, Redis, etc.
* ***CRUD*** - Crud is on the heart of API Brew.
* ***Swagger*** - Swagger docs are generated automatically
* ***Authentication*** - API Brew supports various authentication methods. Including JWT authentication etc.
* ***Authorization*** - API Brew supports authorization. You can define permissions for your resources
* ***Multi Database*** - You can define multiple databases and do operations on top of them
* ***Scalable*** - API Brew is scalable. You can run it on multiple instances, and it will work as expected, as API Brew does not have any data internally, you can scale it.
* ***Extensible*** - API Brew is extensible. You can extend your resources with extensions. You can also extend your APIs with extensions
* ***CLI support*** - API Brew has a cli tool to manage your resources, dataSources, etc. It is called `apbr`
* ***Docker*** - API Brew is dockerized. You can run it on docker
* ***Docker Compose*** - API Brew is docker-compose ready. You can run it on docker-compose
* ***Kubernetes*** - API Brew is kubernetes ready. You can run it on kubernetes

## Use Cases

* Creating backend for your mobile application or website
* Creating backend for your existing database
* Managing your data in a CRUD fashion
* Creating Standardized, well documented APIs for your data
