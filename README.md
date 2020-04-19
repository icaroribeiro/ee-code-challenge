# Estratégia Educacional Code Challenge

## 1 - Introduction

The purpose of this file is to present information about the work developed to solve the code challenge prepared by the company **Estratégia Educacional** that can be founded in the following link: 

*Website*: https://github.com/estrategiahq/desafio-software-engineer

In order to summarise, the project comprehends the implementation of a back-end and a front-end applications for the management of tags related to users' starred repositories from Github. It is composed by a **REST** API developed using **Go** programming language and web pages developed using ReactJS programming languages with Redux, in addition to a **Postgres** database.

Throughout this documentation, a few aspects will be highlighted, such as, the configuration of environment variables of the **Postgres** database and the procedures adopted to run the project with **Docker** containers.

Finally, the last section named **Project Dynamics** illustrates a brief report of how the solution works in practice when using the back-end and front-end applications together.

## 2 - API Documentation

The documentation of the API implemented in **Go** programming language was developed following the **OpenAPI 3.0** specification. Inside the directory **api-docs** there is a script named **openapi-json-to-html.py**. When running it using the **openapi.json** file, it is generated a HTML page named **index.html** within the current directory that illustrates details about the API *endpoints*.