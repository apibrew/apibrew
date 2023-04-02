Data Handler [![Build Status](https://app.travis-ci.com/tislib/data-handler.svg?branch=master)](https://app.travis-ci.com/tislib/data-handler)
======================

# Overview

**Data Handler** is an application to manage your data in a CRUD fashion

There are two main things in data-handler:

* Resource - This is your schema entry
* Record - This is your data

Data Handler by default uses postgresql as a database for your data but it also supports various databases (mysql, mongo, etc.)

You can define your schema, and it will prepare you CRUD APIs (Rest, Grpc) and Swagger docs

**Data Handler can either use your existing database schema or create new schema for you**

Quick Example:
![](http://static.tisserv.net/dh_overview.gif)

country.hcl
```
schema {
   resource "country" {
    property "name" {
      type   = "string"
      length = 124
    }
    property "description" {
      type   = "string"
      length = 124
    }
   }
}
```
```
dhctl apply -f country.hcl
```

Swagger: http://localhost:9009/docs/index.html
```
# Create Country
curl -X POST --location "http://localhost:9009/country" \
    -H "Authorization: <token>" \
    -d "{
   \"name\": \"Country1\",
   \"description\": \"Sample Country 1\"
}"

# List Countries
curl "http://localhost:9009/country" -H "Authorization: <token>"
```


# Quick Start

Let's run application on standalone mode:

```
docker run -d -p 9009:9009 tislib/data-handler:full-latest
```

Let's install our client
```
go install github.com/tislib/data-handler/cmd/dhctl@latest
```

Now let's configure our client to point to server
```
mkdir -p ~/.dhctl
nano  ~/.dhctl/config
```
Paste config to there
```
type: server
servers:
  - name: local
    host: 127.0.0.1:9009
    authentication:
      username: admin
      password: admin
defaultServer: local
```


[![Build Status](docs/static/Overview.png)](https://app.travis-ci.com/tislib/data-handler)

**Website: [https://datahandler.talehibrahimli.com/](https://datahandler.talehibrahimli.com/)**

Data Handler is a **Low Code software** that allows to create various Grpc and Rest APIs from various database platforms

You can use **Data Handler** for following purposes:

1. You have a database and you want to build CRUD like APIs on it without coding or by minimal coding
2. You want to manage data source alongside with creating APIs. For example
    1. You want to create new table/collection in your datasource and create API without much coding
3. You want to manage multiple data source and do operations on top of them with auto created APIs

# Installation

## Installation (server)

### Running with docker

You need to copy and adjust init config for data-handler
[init.json](../old/examples/data/init.example.json)

```
 docker run -v <path-to-init.json>:/app/config.json  tislib/data-handler:latest
```

## Installation (client)

### Installing via go install

```
go install github.com/tislib/data-handler/cmd/dhctl@latest
```

You need to configure client pointing to right server:
[config](../old/examples/data/dhctl.example.config) and move this file to ~/.dhctl/config

Sample commands

```
dhctl get resources
dhctl get user -n system
```

For detailed documentation: [link](docs/content/old/dhctl.md)