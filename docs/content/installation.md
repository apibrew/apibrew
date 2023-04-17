Installation
========

## Docker [standalone]

Docker standalone mode is most basic mode to run data handler

### Prerequisites

* Docker

### Run

*Run docker container*

```bash
docker run -d -p 9009:9009 -v ${PWD}/data:/var/lib/postgresql/data tislib/data-handler:full-latest
```

## Docker [normal]

For running application on docker normal mode you need to have postgresql database running on your host

### Prerequisites
* Docker
* Postgresql

### Configure
see [config.json](files/config.json)

You need to update systemDataSource and initDataSources in config.json

After that you need to mount config.json to /app/config.json

You also need to create two database on your postgresql server. see [init.sql](files/init.sql)

### Run

*Run docker container*

```bash
docker run -d -p 9009:9009 -v ${PWD}/config.json:/app/config.json tislib/data-handler:latest
```

## Docker compose ([see](https://github.com/tislib/data-handler/tree/master/deploy/docker-compose))

## Kubernetes ([see](https://github.com/tislib/data-handler/tree/master/deploy/kubernetes))


