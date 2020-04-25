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

The following directories contain the **REST** API implementation using **Go** programming language.

**back-end/graphql**: it contains the schema for querying starred repositories related to Github users implemented based on the Github v4 API with **GraphQL**.

**back-end/handlers**: it contains the handling of API requests, as well as the elaboration of API responses.

**back-end/handlers_test**: it contains the tests for handling requests to the API *endpoint* using the **Go** language test package.

**back-end/middlewares**: it contains intermediate validations of parameters transmitted through API requests.

**back-end/models**: it contains the definition of the data entities used by both the API and the database.

**back-end/postgresdb**: it contains the implementation directed to the database configuration along with **CRUD** operations (*create*, *read*, *update* and *delete*).

**back-end/postgresdb_test**: it contains the tests of the implementation of **CRUD** operations using the **Go** language test package.

**back-end/router**: it contains a router that exposes the routes associated with the API *endpoints*.

**back-end/router/routes**: it contains the routes associated with the API *endpoints*.

**back-end/server**: it contain an abstraction of the server that allows to "attach" some resources in order to make them available during the API requests. Here, it's used to store the Github personal access token as well as other structure that holds attributes to manage the data.

**back-end/utils**: it contains supporting functions, such as, to elaborate the API responses in JSON-like format and to generate random data used during the tests.

**back-end/.env**: it contains the environment variables for the configuration of the **development** environment.

**back-end/.test.env**: it contains the environment variables for the configuration of the **test** environment.

The **back-end/.env** file contains the environment variables referring to the personal access token required to authenticate to Github and the connection to **Postgres** database, as well as the exposure of the access address for HTTP communication, as indicated below:

```
GITHUB_PERSONAL_ACCESS_TOKEN=<Token>
```

```
DB_USERNAME=user
DB_PASSWORD=password
DB_HOST=db
DB_PORT=5432
DB_NAME=db
```

```
HTTP_SERVER_HOST=0.0.0.0
HTTP_SERVER_PORT=8080
```

In order to not compromise the integrity of the database used by the project in terms of data generated from the execution of the test cases, two Postgres databases will be used.

In this sense, to facilitate future explanations regarding the details of the databases, consider that the database used for the storage of data in a "normal" actions is the **development** database and the one used for the storage of data resulting from the test cases is the **test** database.

These databases are named **db** and **test-db** by the environment variables **DB_NAME** of the **back-end/.env** file and **TEST_DB_NAME** of the **back-end/.test.env** file, respectively.

(P.S. It is necessary to pay special attention to the database environment variables defined in these two previous files in case they are changed.)

### 3.2 - Front-end

The following directories contain the **graphical interfaces** implementation using **React** programming language and **Redux**.

**front-end/src/actions**: it contains the **React** constants and action creators.

**front-end/src/components**: it contains components and their shared styles that are used to determine the layout of the application.

**front-end/src/containers**: it contains a "submission" container to connect the navigation logic to the application state using **Redux**. This container is also used to manage the repositories that will be presented in the repository table.

**front-end/src/pages**: it contains the application pages.

**front-end/src/reducers**: it contains the mechanisms to manage the update of the application state based on actions triggered using **Redux**, enabling the asynchronous dispatching of actions.

**front-end/src/routes**: it contains the routes associated with the API *endpoints* of the **back-end** application.

**front-end/src/service**: it contains functions to manage external communication with the API developed in the **back-end** application using the HTTP client based on *promise* named **axios**.

**front-end/src/store**: it contains the implementation of the creation of the **store** of the **Redux**.

**front-end/.env**: it contains the environment variables for the configuration of the **development** environment.

The **front-end/.env** file contains the environment variables referring to the *host* that composes each *URL* to access the API *endpoints* and the maximum number of suggested tags for each repository while editing its related tags.

```
REACT_APP_API_HOST=<Host>
```

In continuity, the maximum number of suggested tags for each repository is configured by the **REACT_APP_MAX_TAGS_NBR** environment variable:

```
REACT_APP_MAX_TAGS_NBR=<Maximum number of suggested tags>
```

### 3.3 - Postgres

The **postgresdb/scripts/1-create_tables.sql** file contains instructions for creating the **repositories** and **user_repositories** tables, as detailed below:

#### 3.2.1 - Tables

**Repositories**

In the **repositories** table each record contains the data of a repository.

As follows, the **id** field refers to the unique identifier of the repository and the **name**, **description**, **url** and **language** fields refers to its name, description, url and language, respectively.

| Fields      | Data type     | Extra                |
|:------------|:--------------|:---------------------|
| id          | VARCHAR (255) | NOT NULL PRIMARY KEY |
| name        | VARCHAR (255) | NOT NULL             |
| description | VARCHAR (255) | NOT NULL             |
| url         | VARCHAR (255) | NOT NULL             |
| language    | VARCHAR (255) | NOT NULL             |

**User Repositories**

In the **user_repositories** table each record contains the data of a user repository.

As follows, the **user_id** and **repository_id** fields refers to the unique identifier of the user and the repository, respectively, and the tags field refers to a list of names associated to the related repository.

| Fields        | Data type           | Extra                |
|:--------------|:--------------------|:---------------------|
| user_id       | VARCHAR (255)       | NOT NULL PRIMARY KEY |
| repository_id | VARCHAR (255)       | NOT NULL PRIMARY KEY |
| tags          | VARCHAR (255) ARRAY |                      |

The **user_id** is the username of the user in Github.

#### 3.2.2 - Configurations of Docker database containers

To execute the solution through **Docker** containers, it is necessary to relate the environment variables of the **postgresdb/.env** and **postgresdb/.test.env** files with the corresponding environment variables directed to the development and test databases defined in the **back-end** application settings.

To do this, the environment variables of the **postgresdb/.env** and **postgresdb/.test.env** files must be associated with the environment variables of the **back-end/.env** and **back-end/.test.env** files, respectively.

Additionally, it is necessary to indicate that the environment variable **DB_HOST** of the **back-end/.env** and **back-end/.test.env** files must be related to the database **services** defined in the **docker-compose.yml** file.

The **docker-compose.yml** file contains the database services:

```
services:
  ...

  db:
    container_name: db
    build:
      context: ./postgresdb
      dockerfile: Dockerfile
    env_file:
      - ./postgresdb/.env
    ...

  test-db:
    container_name: test-db
    build:
      context: ./postgresdb
      dockerfile: Dockerfile
    env_file:
      - ./postgresdb/.test.env
    ...
```

**Development**

File **postgresdb/.env**:

```
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=db
```

File **back-end/.env**:

```
DB_USERNAME=user
DB_PASSWORD=password
DB_HOST=db
DB_PORT=5432
DB_NAME=db
```

**Test**

File **postgresdb/.test.env**:

```
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=test-db
```

File **back-end/.test.env**:

```
DB_USERNAME=user
DB_PASSWORD=password
DB_HOST=test-db
DB_PORT=5432
DB_NAME=test-db
```

**Important note**

After the project has been successfully executed, it is possible to check the data of the development and test databases resulting from the operations carried out at a command prompt with access to instructions directed to Docker:

```
$ docker exec -it <The id of the container of the corresponding database> /bin/bash
```

To do this, we need to inform the username and password that were previously defined by the database environment variables prior to accessing data.

```
$ psql -U <Username> <Database name>
```

In the case of the envinronment variables are kept as they were delivered, if the **id** of the container corresponds to the service named **db**, the data are obtained from the development database:

```
$ psql -U user db
```

On the other hand, if the **id** of the container corresponds to the service named **test-db**, the data are obtained from the test database:

```
$ psql -U user test-db
```

## 4 - How to execute the project?

**In the case of the environment variables of all the .env and .test.env files from all directories are kept as they were delivered I strongly believe that it will not be necessary any change before executing the project.**.

**Prior** to run the project, it is first necessary to configure two settings:

1. The personal access token required to authenticate to Github in order to get the Github starred repositories of a specific user by its username.

```
  ...
  back-end:
    container_name: back-end
    ...
    environment:
      - GITHUB_PERSONAL_ACCESS_TOKEN=<Github personal access token>
    ...
```

2. The IP address (*host*) configured by Docker so that the **front-end** application can communicate with the **back-end** application using HTTP.

The *host* corresponds to the value informed when executing a command at a command prompt with access to instructions directed to Docker:

```
$ docker-machine ip
```

After that, navigate to the project's root directory where the **docker-compose.yml** file is, and assign the *host* to the **REACT_APP_API_HOST** variable in the **front-end** service:

```
  ...
  front-end:
    container_name: front-end
    ...
    environment:
      - REACT_APP_API_HOST=<IP address configured by Docker>
    ...
```

Still at a command prompt with access to instructions directed to Docker where the docker-compose.yml file is, run the command:

```
$ docker-compose up -d
```

(P.S. Because of the dependencies related to the back-end services, it may take some time to them attach to other services properly. Then, to confirm if everything is up and running ok execute the command *docker container ls -a*)

If there are no errors, the API *endpoints* will be accessed using the address composed by the *host* and the HTTP server port **8080**. For example:

```
http://{host}:8080
```

In continuity, suppose the *host* is: 192.168.99.100. As a result, the API requests can be performed through a front-end client or test tool like Postman using the address as:

```
http://192.168.99.100:8080
```

In addition, it is also worth emphasizing that the entire configuration related to **Docker** was evaluated in this documentation based on the **DockerToolbox** tool for Windows.

## 5 - How to use the API *endpoints*?

The API request are performed through the HTTP server port **8080** and the API responses can be viewed by means of a **front-end** client or test tool, for example **Postman**.

**Important note**

The **userId** informed in the request URL of all API requests should be replaced by the **username** of a Github user whereas the **repositoryId** is just the **id** of the related Github repository.

In what follows, there is a guide that includes API requests for creating, obtaining, updating and deleting data from the database using the Github user named **icaroribeiro**.

(P.S. Before checking the following examples, consider that no data is recorded prior to this explanation.)

### Status

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8080/status
```

Response:

```
Code: 200 OK - In the case of the service has started up correctly and is ready to accept requests.
```

### Management of Repositories

#### Listing of User Starred Repositories from Github

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8080/users/icaroribeiro/githubStarredRepositories
```

Response:

```
Code: 200 OK - In the case of the user starred repositories are successfully obtained from Github.
```

```
*application/json**

Body: [
    {
        "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
        "name": "go",
        "description": "The Go programming language",
        "url": "https://github.com/golang/go",
        "language": "Go"
    }
]
```

#### Creation of a Repository

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8080/users/icaroribeiro/repository
```

```
Body: {
    "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
    "name": "go",
    "description": "The Go programming language",
    "url": "https://github.com/golang/go",
    "language": "Go"
}
```

Response:

```
Code: 201 Created - In the case of the repository is successfully created.
```

```
*application/json*

Body: {
    "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
    "name": "go",
    "description": "The Go programming language",
    "url": "https://github.com/golang/go",
    "language": "Go"
}
```

**Important note**

In brief, in addition to create the association between the user and its starred repository in the **user_repositories** table, it is also verified if the related repository was previously created in the **repositories** table. If not, then the repository is created in the table. On the other hand, it is checked if any data of the repository changed. If so, then the repository is updated in the table.

#### Edition of Tags by updating a Repository by its id

Request:

```
Method: HTTP PUT
```

```
URL: http://{host}:8080/users/icaroribeiro/repositories/MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==
```

```
Body: {
  	"tags": [
        "go", 
        "golang"
    ]
}
```

Response:

```
Code: 200 OK - In the case of the repository is successfully updated.
```

```
*application/json*

Body: {
    "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
    "name": "go",
    "description": "The Go programming language",
    "url": "https://github.com/golang/go",
    "language": "Go",
    "tags": [
        "go",
        "golang"
    ]
}
```

#### Deletion of a Repository by its id

Request:

```
Method: HTTP DELETE
```

```
URL: http://{host}:8080/users/icaroribeiro/repositories/MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==
```

Response:

```
Code: 200 OK - In the case of the repository is successfully deleted.
```

```
*application/json*

Body: {
    "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
    "name": "go",
    "description": "The Go programming language",
    "url": "https://github.com/golang/go",
    "language": "Go",
    "tags": [
        "go",
        "golang"
    ]
}
```

**Important note**

In brief, in addition to delete the association between the user and its starred repository from the **user_repositories** table, it is also verified if the related repository is associated with any other user in the table. If not, the repository is deleted from the **repositories** table. 

## 6 - Tests

In order to test the solution two **test sets** were developed.

The tests will be executed on the running **back-end** containers. 

To do this, at a command prompt with access to instructions directed to Docker, launch a bash terminal within the related **back-end** container:

```
$ docker exec -it <The id of the container of the corresponding back-end application> /bin/bash
```

### 6.1 Database

The tests that were developed are related to **CRUD** operations (*create*, *read*, *update* and *delete*) in the test database.

To execute them on the running **back-end** container, navigate to the **app/postgresdb_test** directory.

So, if you prefer to evaluate all tests at once, run the command:

```
$ go test -v
```

However, it is also possible to run each test separately using the commands:

**Tests of the CRUD operations directed to User Repositories**

```
$ go test -v -run=TestCreateUserRepository
```

```
$ go test -v -run=TestGetUserRepository
```

```
$ go test -v -run=TestGetAllUserRepositoriesByUserId
```

```
$ go test -v -run=TestGetAllUserRepositoriesByRepositoryId
```

```
$ go test -v -run=TestUpdateUserRepository
```

```
$ go test -v -run=TestDeleteUserRepository
```

**Tests of the CRUD operations directed to Repositories**

```
$ go test -v -run=TestCreateRepository
```

```
$ go test -v -run=TestGetRepository
```

```
$ go test -v -run=TestUpdateRepository
```

```
$ go test -v -run=TestDeleteRepository
```

### 6.2 Handlers

These tests are related to the API requests.

To execute them on the running **back-end** container, navigate to the **app/handlers_test** directory.

So, if you prefer to evaluate all tests at once, run the command:

```
$ go test -v
```

Nevertheless, it is also possible to run each test separately using the commands:

**Tests of the API requests directed to User Repositories**

```
$ go test -v -run=TestCreateUserRepository
```

```
$ go test -v -run=TestGetAllUserRepositories
```

```
$ go test -v -run=TestUpdateUserRepository
```

```
$ go test -v -run=TestDeleteUserRepository
```

**Tests of the API requests directed to Repositories**

```
$ go test -v -run=TestGetRepository
```

```
$ go test -v -run=TestUpdateRepository
```

## 7 - Project Dynamics

In what follows, there is a brief account of how the solution works in practice meeting the requirements specified in the comments of the code challenge.

### 7.1 - Home Page

The project's home page is presented by the **front-end** application when accessing the address in the following format where the *host* corresponds to the IP address configured by Docker:

``
http://{host}:3000
``

As previously explained, the *host* corresponds to the value informed when executing a command at a command prompt with access to instructions directed to Docker: 

```
$ docker-machine ip
```

To access the page with the table of the starred repositories, first, it is necessary to enter a **username** of a valid Github user in the text box shown and then click on the **get repositories** button.

By clicking on the **get repositories** button, a **set of evaluations** is developed involving repositories (as well as their related tags, if any) that can result in the creation, editing or even deletion of records from the database performed by API requests directed to the **back-end** application.

### 7.1.1 - Evaluation of the repositories associated with a Github user

Initially, **two** lists of repositories are formulated and, later, their repositories are compared with each other considering **only** the value of the unique identifier of each repository, that is, the **id** field.

A list of repositories is composed **only** by the current starred repositories of the Github user and they are obtained through the API request:

Request:

```
Method: HTTP GET
```

Response:

```
Code: 200 OK - In the case of the Github starred repositories are successfully obtained.
```

```
*application/json*

Body: [
    {
        "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
        "name": "go",
        "description": "The Go programming language",
        "url": "https://github.com/golang/go",
        "language": "Go"
    },
    {
        "id": "MDEwOlJlcG9zaXRvcnkxNTA0NTc1MQ==",
        "name": "compose",
        "description": "Define and run multi-container applications with Docker",
        "url": "https://github.com/docker/compose",
        "language": "Python"
    }
]
```

The other list of repositories is composed **only** by the repositories previously created in the **user_repositories** table that are associated with the same user and they are obtained through the API request:

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8080/users/icaroribeiro/repositories
```

Response:

```
Code: 200 OK - In the case of the user repositories are successfully obtained.
```

```
*application/json*

Body: [
    {
        "id": "MDEwOlJlcG9zaXRvcnkxNTA0NTc1MQ==",
        "name": "compose",
        "description": "Define and run multi-container applications with Docker",
        "url": "https://github.com/docker/compose",
        "language": "Python"
    },
    {
        "id": "MDEwOlJlcG9zaXRvcnkxMjM0NzE0",
        "name": "elixir",
        "description": "Elixir is a dynamic, functional language designed for building scalable and maintainable applications",
        "url": "https://github.com/elixir-lang/elixir",
        "language": "Elixir"
    }
]
```

In other words, the list of repositories above consists of database records that may have been inserted at any time when accessing the project using the same **username** as the Github user.

(P.S. As already indicated, after obtaining these **two** lists of repositories, their repositories will be compared with each other considering **only** their **id** fields.)

#### Creation of Repositories

Firstly, in the **front-end** application, it is analyzed **only** the current starred repositories of the Github user that are not registered in the **user_repositories** table. In short, it is "filtered" the repositories that are on the first list, but are not on the second list.

Each of these repositories will be associated with the related user in the **user_repositories** table through the API request:

Request:

```
Method: HTTP POST
```

```
URL: http://{host}:8080/users/icaroribeiro/repository
```

```
*application/json*

Body: {
    "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
    "name": "go",
    "description": "The Go programming language",
    "url": "https://github.com/golang/go",
    "language": "Go"
}
```

Response:

```
Code: 201 OK - In the case of the user repository is successfully created.
```

```
*application/json*

Body: {
    "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
    "name": "go",
    "description": "The Go programming language",
    "url": "https://github.com/golang/go",
    "language": "Go"
}
```

**Additional step**

Still on this scenario, the creation of a record in the **user_repositories** table can imply in one of the two additional steps to be executed:

1. If the repository that is associated with the Github user in the **user_repositories** table has not yet been created in the **repository** table, that is, the repository was never related to any user already recorded, then it will be created in the **repository** table. 

2. Otherwise, if the repository is already associated with any user already recorded, it means that the repository was previously created in the **repository** table. In this case, the data of the repository is updated **only** in the **repository** table. It is considered appropriate since its data may have changed over time, for example its **description**.

(P.S. Based on the explanation above, a repository is **only** created in the **repository** table when there is an association between it and some Github user in the **user_repositories** table through the **repository_id** field.)

#### Deletion of Repositories

Subsequently, in the **front-end** application, it is analyzed **only** the repositories associated with the user in the **user_repositories** table but are no longer starred by him/her in Github. In short, it is "filtered" the repositories that are on the second list, but are not on the first list.

Each of records related to these repositories (along with its tags, if any) will be deleted from the **user_repositories** table through the API request:

Request:

```
Method: HTTP DELETE
```

```
URL: http://{host}:8080/users/icaroribeiro/repositories/MDEwOlJlcG9zaXRvcnkxMjM0NzE0
```

Response:

```
Code: 200 OK - In the case of the repository is successfully deleted.
```

```
*application/json*

Body: {
    "id": "MDEwOlJlcG9zaXRvcnkxMjM0NzE0",
    "name": "elixir",
    "description": "Elixir is a dynamic, functional language designed for building scalable and maintainable applications",
    "url": "https://github.com/elixir-lang/elixir",
    "language": "Elixir"
}
```

**Additional step**

Still on this scenario, the deletion of a record in the **user_repositories** table can imply in one of the two additional steps to be executed:

1. If the repository that was associated with the Github user has not been related to any user already recorded in the **user_repositories** table, then it will be deleted in the **repository** table. 

2. Otherwise, if the repository is associated with any user already recorded, it means that the data of the repository is necessary and it can't be deleted from the **repository** table. In this case, nothing happens.

#### Update of Repositories

Finally, in the **front-end** application, it is analyzed **only** the repositories associated with the user in the **user_repositories** table that **also** are starred in Github. In short, the repositories that are in both lists.

These repositories will be updated in the **repository** table through the API request:

Request:

```
Method: HTTP PUT
```

```
URL: http://{host}:8080/repositories/MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==
```

```
Body: {
    "name": "golang",
    "description": "The Go programming language from Google",
    "url": "https://github.com/golang/go",
    "language": "Golang"
}
```

As already commented, it is considered suitable since its data may have changed over time. Above it is just a sample of the data of repository related to **name**, **description** and so forth that could have been modified.

### 7.2 Page with the Repository table

After accessing the project's home page, inserting a **username** of a valid Github user in the text box and clicking the **get repositories** button, the page with the starred repositories associated with the Github user is demonstrated with the components below:

In the upper left corner there is a text box with the message *search by tag* that is used to filter repositories by the name of one of their related tags and the search of a tag works with incomplete words.

Thus, if a repository has a tag named **go** and another repository has a tag called **golang**, when performing a search for repositories using the word **go**, or even just the letter **g**, both repositories will be maintained in the table as a result of the search process.

In the upper right corner there is a *link* called **Home** used to return to the project's home page.

Below these components there is a table of repositories filled out with the following data of the starred repositories: name, description and language.

(P.S. If the **username** entered on the project's home page does not belong to a valid Github user, or even **username** belongs to a valid Github user, but he/she does not have any associated starred repositories, the table of repositories is illustrated without any record.)

In continuity, if the table is filled out with repositories, the tags related to a given repository registered at a previous moment (if any) will be shown below the **Tags** column, in the line corresponding to the related repository.

To the right of the table there is a last column without title and each of its lines has a *link* called **edit** used to edit tags related to the corresponding repositories according to the selected line.

When clicking the **edit** component, a window for editing tags is shown. It contains a text box for editing tags that is **automatically** filled out with words according to two situations:

1. If the related record from the **user_repositories** table already has one or more related words from the **tags** column, these words are illustrated inside the text box.

2. On the other hand, and **only** in this circumstance, if the related record from the **user_repositories** table does not have any related tag, that is, the **tags** column is empty (or *null*) the text box can be filled out with one or more suggested words to identify the repository.

The proposed mechanism for tag suggestion is to get tags (if any) previously assigned to the same repository by other users in the **user_repositories** table through the API request:

Request:

```
Method: HTTP GET
```

```
URL: http://{host}:8080/repositories/MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==
```

Response:

```
Code: 200 OK - In the case of the repository is successfully obtained.
```

```
*application/json*

Body: {
    "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
    "name": "go",
    "description": "The Go programming language",
    "url": "https://github.com/golang/go",
    "language": "Go",
    "tags": [
        "google"
    ]
}
```

(P.S. The API request above is to get all data of the repository. Therefore, its **tags** field refers to all words assigned to the repository by all its associated users from the **user_repositories** table. In the case above, only the **google** tag would be considered as tags suggestion and displayed in the text box.)

As previously indicated, the maximum number of suggested tags for each repository is determined by the **REACT_APP_MAX_TAGS_NBR** environment variable from in the **front-end/.env** file. For example:

```
REACT_APP_MAX_TAGS_NBR=5
```

In order successfully to relate the tags to the repositories the words must be inserted in the text box separated by commas (,) and single space. For example:

```
go, golang
```

On the other hand, in order to not apply any change with regard to the edition of tags, just click the **Cancel** button or the **X** icon in the upper right corner.

In order to describe the edition of tags in more detail, suppose that a given repository does not have any related tags. On other words, the related record from the **user_repositories** table does not contain any value in the **tags** column.

Then, whenever the user writes one or more tags and clicks the **Save** button the related record from the **user_repositories** table will be associated with such words through the API request:

Request:

```
Method: HTTP PUT
```

```
URL: http://{host}:8080/users/icaroribeiro/repositories/MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==
```

```
*application/json*

Body {
    "tags": [
      "go", "golang"
    ]
}
```

Response:

```
Code: 200 OK - In the case of the user repository is successfully updated.
```

```
*application/json*

Body: {
    "id": "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==",
    "name": "go",
    "description": "The Go programming language",
    "url": "https://github.com/golang/go",
    "language": "Go",
    "tags": [
        "go",
        "golang"
    ]
}
```

**Important note**

If the user no longer wants any tag associated to the repository, after clicking **edit** component first he/she must delete all tags from the text box and then click the **Save** button. In this case, the API request body is composed by the **tags** field as an empty array.

Therefore, whenever tags must be created, edited or even removed, the API request body is **only** composed by the **tags** field along with the words (if any) to be associated with the related record in the **user_repositories** table.

### 7.3 Error page

Finally, an error page is presented to the user when the user tries to access an invalid address, that is, an address different from those previously demonstrated. For example:

```
http://{host}:3000/<Invalid address>
```