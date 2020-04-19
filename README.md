# Olist Code Challenge

## 1 - Introduction

The purpose of this file is to present information about the work developed to solve the code challenge prepared by the company **Olist** that can be founded in the following link: 

*Website*: https://github.com/olist/work-at-olist/blob/master/README.md

In order to summarise, the project comprehends the implementation of a back-end application for the management of books along with their authors in a library. It is composed by a **REST** API developed using **Go** programming language, in addition to a **Postgres** database.

Throughout this documentation, a few aspects will be highlighted, such as, the configuration of environment variables of the **Postgres** database and the procedures adopted to run the project with **Docker** containers.

Finally, the last section illustrates the details of the deployment of the solution on **Heroku** hosting service.

(P.S. This is the job opportunity that I applied for: Desenvolvedor GO - Presencial (SP) - https://olist.gupy.io/jobs/98273)

## 2 - API Documentation

The documentation of the API implemented in **Go** programming language was developed following the **OpenAPI 3.0** specification. Inside the directory **api-docs** there is a script named **openapi-json-to-html.py**. When running it using the **openapi.json** file, it is generated a HTML page named **index.html** within the current directory that illustrates details about the API *endpoints*.