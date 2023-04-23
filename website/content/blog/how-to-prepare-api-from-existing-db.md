---
title: How to prepare API from existing DB in 5 mins?
date: 2023-04-23
---

## Introduction

With help of API Brew, you can build API from your existing DB. In this article, we will show you how to do it.

## Prerequisites

* API Brew server is running and accessible
* dhctl is installed
* You have existing DB

For installing API Brew server, please refer to [Installation](/docs/installation#easy-install)

## Connect API Brew to your exising DB

We need to prepare new Data Source for our DB. For this, we will use dhctl.

```yaml
type: datasource
name: mydb
backend: POSTGRESQL
postgresqlParams:
  host: localhost
  port: 5432
  dbName: dh_db_2
  username: root
  password: root
```

Now let's test connection to our DB

```bash
dhctl data-source status --name=mydb
```

You will see following output

```textmate
INFO[0000] DataSource name: mydb                        
INFO[0000] ConnectionAlreadyInitiated: false            
INFO[0000] TestConnection: true                         
```

## Prepare API from DB

Let's assume that, we have author table in DB, and we want to expose it as API.
For testing purposes you can execute this query on DB to have author table

```SQL
create table public.author
(
    id          bigserial primary key,
    name        character varying(255) not null,
    description character varying(255)
);
```

Now, let's see what tables we have in our DB, We will get this information from API Brew client

```bash
dhctl data-source list-entities --name=mydb
```

You will see following output

```textmate
CATALOG ENTITY  EDITABLE 
public  author  editable                                   
```

```bash
dhctl data-source prepare --name=mydb > schema.yml
```

This command will generate schema.yml file, which contains all information about resources for our DB.

The schema.yml file will look like this
```yaml
annotations:
  AutoCreated: "true"
  DisableVersion: "true"
name: author
namespace: default
properties:
  - annotations:
      Identity: "true"
      SourceDef: int8
    mapping: id
    name: id
    primary: true
    required: true
    type: INT64
  - annotations:
      SourceDef: varchar(255)
    length: 255
    mapping: name
    name: name
    required: true
    type: STRING
  - annotations:
      SourceDef: varchar(255)
    length: 255
    mapping: description
    name: description
    type: STRING
sourceConfig:
  catalog: public
  dataSource: mydb
  entity: author
type: resource
```

Now, let's apply our resources

```bash
dhctl apply -f schema.yml -m=false
```

## Test API

So, Bingo! We have our API ready. You can check it by visiting Swagger UI

You can use following url to see swagger UI http://localhost:9009/docs/index.html

You can also call curl directly to see the result
```bash
## Create author
curl -X 'POST' \
  'http://localhost:9009/author' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer <token>' \
  -H 'Content-Type: application/json' \
  -d '{
  "description": "123",
  "name": "abc"
}'

## List authors
curl -X 'GET' \
  'http://localhost:9009/author' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer <token>'
```

## Conclusion
With help of API Brew, you can build API from your existing DB. In this article, we have shown you how to do it.
