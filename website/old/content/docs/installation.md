Installation
========

## Easy install 
```bash
curl -L https://raw.githubusercontent.com/tislib/apibrew/master/deploy/easy-install/run.sh | bash
```

It will install both API Brew and dhctl for you

## Docker [standalone]

Docker standalone mode is most basic mode to run API Brew

### Prerequisites

* Docker

### Run

*Run docker container*

```bash
docker run -d -p 9009:9009 -v ${PWD}/data:/var/lib/postgresql/data tislib/apibrew:full-latest
```

## Docker [normal]

For running application on docker normal mode you need to have postgresql database running on your host

### Prerequisites
* Docker
* Postgresql

### Configure
see [config.json](../files/config.json)

You need to update systemDataSource and initDataSources in config.json

After that you need to mount config.json to /app/config.json

You also need to create two database on your postgresql server. see [init.sql](../files/init.sql)

### Run

*Run docker container*

```bash
docker run -d -p 9009:9009 -v ${PWD}/config.json:/app/config.json tislib/apibrew:latest
```

## Docker compose
[see](https://github.com/tislib/apibrew/tree/master/deploy/docker-compose)

## Kubernetes
[see](https://github.com/tislib/apibrew/tree/master/deploy/kubernetes)

## Client 
Let's install our client **dhctl**

You can download client binary from release page https://github.com/tislib/apibrew/releases/latest (download dhctl-OS-ARCH)

You can also use go install if you have go runtime on your local
```
go install github.com/tislib/apibrew/cmd/dhctl@latest
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



