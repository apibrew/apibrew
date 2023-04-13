Data Handler [![Build Status](https://app.travis-ci.com/tislib/data-handler.svg?branch=master)](https://app.travis-ci.com/tislib/data-handler)
======================


See [docs](docs/index.md) for detailed information

# Overview

Data Handler is a **Low Code software** that allows to create various Grpc and Rest APIs from various database platforms

You can use **Data Handler** for following purposes:

1. You have a database and you want to build CRUD like APIs on it without coding or by minimal coding
2. You want to manage data source alongside with creating APIs. For example
    1. You want to create new table/collection in your datasource and create API without much coding
3. You want to manage multiple data source and do operations on top of them with auto created APIs

**Data Handler** is an application to manage your data in a CRUD fashion

There are two main things in data-handler:

* Resource - This is your schema entry
* Record - This is your data

Data Handler by default uses postgresql as a database for your data but it also supports various databases (mysql,
mongo, etc.)

You can define your schema, and it will prepare you CRUD APIs (Rest, Grpc) and Swagger docs

**Data Handler can either use your existing database schema or create new schema for you**

Quick Example:
![](http://static.tisserv.net/dh_overview.gif)

country.yml

```
type: resource
name: country
properties:
  - name: name
    type: STRING
    length: 255
    required: true
    unique: true
  - name: description
    type: STRING
    length: 255
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

# [Tutorials](docs/content/tutorials/index.md)

you can find various tutorials [here](docs/content/tutorials/index.md)

# [Docs](docs/index.md)

- [Tutorials](docs/content/tutorials/index.md) - Tutorials
- [Installation](docs/content/installation.md) - Installation
- [General](docs/content/general.md) - General information about Data Handler
- [API](docs/content/api.md) - GRPC and Rest API documentation
- [CLI](docs/content/dhctl/dhctl.md) - CLI interface(dhctl) documentation
- [Proto](docs/content/proto.md) - Proto documentation
- [SDK](docs/content/sdk.md) - SDK documentation
