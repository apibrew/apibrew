---
title: About API Brew
linkTitle: About
menu: { main: { weight: 10 } }
---

{{% blocks/cover height="auto" %}}

[//]: # (<img style="width:150px;display: block; margin:0 auto; margin-top:-100px" src="/logo/logo.svg"/>)

<video autoplay="autoplay" muted playsinline id="myVideo" controls loop style="width:1000px; ">
  <source src="http://static.tisserv.net/apibrew_demo_recording_1_compress.mov" type="video/mp4">
</video>


Apibrew is an open-source tool that allows you to quickly and easily create CRUD APIs for your data. With Apibrew, you
define your schema in a declarative way, and the tool generates the corresponding APIs for you, minimizing the amount of
coding you have to do.

Apibrew is designed to be low-code, which means that you can create APIs without writing a lot of code. It is also
database agnostic and supports various database platforms, including PostgreSQL, MongoDB, MySQL, Redis, and more.
Apibrew supports multiple databases, so you can define multiple databases and perform operations on them.

{{% /blocks/cover %}}

{{% blocks/section color="white" %}}

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
* ***CLI support*** - API Brew has a cli tool to manage your resources, dataSources, etc. It is called `dhctl`
* ***Docker*** - API Brew is dockerized. You can run it on docker
* ***Docker Compose*** - API Brew is docker-compose ready. You can run it on docker-compose, see [docker-compose](deploy/docker-compose)
* ***Kubernetes*** - API Brew is kubernetes ready. You can run it on kubernetes, see [kubernetes](deploy/kubernetes)

## Use Cases

* Creating backend for your mobile application or website
* Creating backend for your existing database
* Managing your data in a CRUD fashion
* Creating Standardized, well documented APIs for your data


{{% /blocks/section %}}
