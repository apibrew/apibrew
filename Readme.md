Data Handler [![Build Status](https://app.travis-ci.com/tislib/data-handler.svg?branch=master)](https://app.travis-ci.com/tislib/data-handler)
======================

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
[init.json](data/init.example.json)

```
 docker run -v <path-to-init.json>:/data/init.json  tislib/data-handler:latest
```

## Installation (client)
### Installing via go install

```
go install github.com/tislib/data-handler/cmd/dhctl@latest
```

You need to configure client pointing to right server:
[config](data/dhctl.example.config) and move this file to ~/.dhctl/config

Sample commands
```
dhctl get resources
dhctl get user -n system
```
For detailed documentation: [link](docs/dhctl.md)