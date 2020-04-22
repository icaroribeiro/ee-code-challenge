# Estratégia Educacional Code Challenge

## 1 - Introduction

The purpose of this file is to present information about the work developed to solve the code challenge prepared by the company **Estratégia Educacional** that can be founded in the following link: 

*Website*: https://github.com/estrategiahq/desafio-software-engineer

In order to summarise, the project comprehends the implementation of a back-end and a front-end applications for the management of tags related to users Github starred repositories. It is composed by a **REST** API developed using **Go** programming language and web pages developed using ReactJS programming languages with Redux, in addition to a **Postgres** database.

Throughout this documentation, a few aspects will be highlighted, such as, the configuration of environment variables of the **Postgres** database and the procedures adopted to run the project with **Docker** containers.

Finally, the last section named **Project Dynamics** illustrates a brief report of how the solution works in practice when using the back-end and front-end applications together.

## 2 - API Documentation

The documentation of the API implemented in **Go** programming language was developed following the **OpenAPI 3.0** specification. Inside the directory **api-docs** there is a script named **openapi-json-to-html.py**. When running it using the **openapi.json** file, it is generated a HTML page named **index.html** within the current directory that illustrates details about the API *endpoints*.


## 3 - Project Organization

The developed solution is organized according to the structure of directories and files summarized below:

### 3.1 - Back-end

### 3.2 - Front-end

### 3.3 - Postgres

The **postgresdb/scripts/1-create_tables.sql** file contains instructions for creating the **repositories** and **user_repositories** tables, as detailed below:

#### 3.2.1 - Tables

**Authors**

In the **authors** table each record contains the data of an author.

This way, the **id** field refers to the unique identifier of the author and the **name** field refers to his/her full name.

| Fields        | Data type     | Extra                |
|:--------------|:--------------|:---------------------|
| id            | INTEGER       | NOT NULL PRIMARY KEY |
| name          | VARCHAR (255) | NOT NULL             |

**Books**

In the **books** table each record contains the data of a book.

This way, the **id** field refers to the unique identifier of the book and the **name**, **edition**, **publication_year** and **authors** fields refer to its full name, edition, year of publication and a list of ids of all its authors, respectively.

| Fields           | Data type     | Extra                |
|:-----------------|:--------------|:---------------------|
| id               | INTEGER       | NOT NULL PRIMARY KEY |
| name             | VARCHAR (255) | NOT NULL             |
| edition          | INTEGER       | NOT NULL             |
| publication_year | INTEGER       | NOT NULL             |
| authors          | INTEGER ARRAY | NOT NULL             |

The list of all ids of its authors is configured as follows:

```
*application/json*

"authors": [
        <The id of an author>
    ]
```









