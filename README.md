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

**Prior** to run the project, it is first necessary to configure the IP address (*host*) configured by Docker so that the **front-end** application can communicate with the **back-end** application using HTTP.

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
$ go test -v -run=TestGetAllUserGithubStarredRepositories
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













## 7 - Project Dynamics

Traduzir:

Logo abaixo é evidenciado um breve relato de como a solução funciona na prática.

### 7.1 - Página Inicial

A página inicial do projeto é apresentada pela aplicação **front-end** ao acessar o endereço no seguinte formato em que o *host* corresponde ao endereço do servidor de acesso:

```
http://{host}:3000
```

Para acessar a página com a tabela de repositórios estrelados, primeiramente, é necessário inserir um *username* de usuário válido do Github na caixa de texto demonstrada e, em seguida, clicar no botão **get repositories**.

Ao clicar no botão **get repositories** é desenvolvida uma **série de avaliações** envolvendo repositórios que resultam na criação, edição ou até mesmo a deleção de repositórios (como também suas tags relacionadas, caso existam) do banco de dados.

### 7.1.1 - Avaliação dos repositórios associados a um usuário do Github

Inicialmente, **duas** listas de repositórios são formuladas e, posteriormente, seus repositórios são comparados entre si considerando **somente** o valor do identificador único de cada repositório, o campo **id**.

Uma lista de repositórios é composta **apenas** pelos repositórios atualmente estrelados pelo usuário do Github, obtidos por meio da seguinte requisição da API:

Requisição:

```
Método: HTTP GET
```

```
URL: http://{host}:8080/users/{userId}/githubStarredRepositories
```

A outra lista de repositórios é composta **apenas** dos repositórios anteriormente criados na tabela **user_repository_tag** que estão associados ao mesmo usuário, obtidos por meio da seguinte requisição da API:

Requisição:

```
Método: HTTP GET
```

```
URL: http://{host}:8080/users/{userId}/repositories
```

A lista de repositórios acima é composta por registros do banco de dados que podem ter sido inseridos em um momento passado ao acessar o projeto usando o mesmo nome do usuário do Github.

Conforme indicado, após a obtenção destas **duas** listas iniciais de repositórios, seus repositórios são comparados entre si considerando **somente** o campo **id**.

#### Criação de repositórios

Primeiramente, é analisado **somente** os repositórios estrelados pelo usuário do Github que não estão registrados na tabela **user_repository_tag**. (Isto é, os repositórios que estão na lista primeira lista, mas não estão na segudna lista).

Cada um destes repositórios vai ser associados ao usuário na tabela **user_repository_tag** por meio da seguinte requisição da API:

Requisição:

```
Método: HTTP POST
```

```
URL: http://{host}:8080/users/{userId}/repository
```

```
Body: {
    "id": <ID do repositório>,
    "name": <Nome do repositório>,
    "description": <Descrição do repositório>,
    "url": <URL do repositório>,
    "language": <Linguagem do repositório>
}
```

Ainda sobre este cenário, caso um determinado repositório recentemente associado ao usuário do Github na tabela **user_repository_tags** ainda não tenha sido criado na tabela **repository**, então ele será registrado na tabela **repository**. Do contrário, isto significa que o repositório foi anteriormente criado na tabela **repository** por também estar associado a algum outro usuário já registrado na tabela **user_repository_tag** no passado.

#### Remoção de repositório

Posteriormente, é avaliado **somente** os repositórios associados ao usuário na tabela **user_repository_tag** que não estão mais estrelados pelo usuário no Github. (Isto é, os repositórios que estão na segunda lista, mas não estão na primeira lista).

Estes repositórios vão ser deletados da tabela **user_repository_tag** por meio da seguinte requisição da API:

Requisição:

```
Método: HTTP DELETE
```

```
URL: http://{host}:8080/users/{userId}/repositories/{repositoryId}
```

Durante a remoção de repositórios associados ao usuário na tabela **user_repository_tag** que não estão mais estrelados no Github é desenvolvida uma análise para cada um deles que pode resultar na deleção de outros registros do banco de dados.

**Remoção de tags**

Inicialmente, é identificado se o repositório a ser deletado possui tags relacionadas com base no campo **tag_id** dos registros correspondentes da tabela **user_repository_tag**, isto é, aqueles registros que possuem o campo **user_id** igual o nome do usuário do Github e o campo **repository_id** igual o identificador único do repositório avaliado.

Caso o repositório tenha tags relacionadas, é preciso verificar se cada uma delas também está relacionada a alguma outra associação entre um usuário do Github e um repositório na tabela **user_repository_tag**.

Se uma determinada tag não está relacionada a outra associação, além da remoção do registro correspondente da tabela **user_repository_tag**, a tag em questão também é removida da tabela **tag**.

**Deleção de Repositório**

Em continuidade, caso um repositório que foi recentemente deletado da tabela **user_repository_tag** não esteja associado a outros usuários na mesma tabela, ele também será removido da tabela **repository**.

#### Atualização de repositório

Finalmente, é avaliado **somente** os repositórios associados ao usuário na tabela **user_repository_tag** que **também** estão estrelados pelo usuário no Github. (Isto é, os repositórios que estão em ambas as listas).

Estes repositórios vão ser atualizados na tabela **repository** por meio da seguinte requisição da API:

Requisição:

```
Método: HTTP PUT
```

```
URL: http://{host}:8080/repositories/{repositoryID}
```

```
Body: {
    "name": <Nome do repositório>,
    "description": <Descrição do repositório>,
    "url": <URL do repositório>,
    "language": <Linguagem do repositório>
}
```

### 7.2 Página com a tabela de repositórios

Após o acesso a página inicial do projeto, a inserção de um *username* de um usuário válido do Github e o clique no botão **get repositories**, a página com os repositórios estrelados associados ao usuário é demonstrada com os componentes descritos abaixo:

No canto superior a esquerda existe uma caixa de texto com a mensagem *search by tag* que é utilizada para filtrar repositórios pelo nome de uma de suas tags relacionadas e a busca por tags funciona com palavras incompletas.

Assim, caso um repositório possua uma tag nomeada **go** e outro repositório tenha uma tag chamada **golang**, ao realizar uma busca por repositórios utilizando a palavra **go**, ou ainda somente a letra **g**, ambos repositórios serão retornados.

No canto superior a direita existe um *link* chamado **Home** utilizado para voltar a página inicial do projeto.

Abaixo destes componentes existe uma tabela de repositórios provida com os seguintes dados: nome do repositório, descrição e linguagem.

A respeito do preenchimento da tabela de repositórios, caso o *username* inserido na página inicial não pertença a um usuário válido do Github, ou ainda, o *username* pertença a usuário válido, porém que não possui repositórios estrelados associados, a tabela de repositórios é ilustrada sem qualquer registro.

Em continuidade, caso a tabela seja preenchida com repositórios, as tags relacionadas a um determinado repositório registradas em um momento anterior serão mostradas abaixo da coluna **Tags**, na linha corresponde ao repositório relacionado.

A direita da tabela existe uma última coluna sem título e cada uma de suas linhas apresenta um *link* chamado *edit* usado para a edição de tags relacionadas aos repositórios correspondentes de acordo com a linha selecionada. Com isso, ao clicar no componente *edit*, uma janela para a edição de tags é demonstrada.

A janela contém uma caixa de texto para a edição de tags que é **automaticamente** preenchida com palavras de acordo com duas situações.

Caso o repositório tenha uma ou mais tags relacionadas de acordo com registros correspondentes criados na tabela **user_repository_tags** em um momento passado, tais palavras são ilustradas dentro da caixa de texto.

Por outro lado, e **somente** nesta circunstância, caso o repositório não tenha qualquer tag relacionada, a caixa de texto é preenchida com sugestões de tags para identificar o repositório.

O mecanismo de sugestão de tags proposto é a busca por tags anteriormente atribuídas por outros usuários associados ao mesmo repositório na tabela **user_repository_tag** que são recuperadas do repositório por meio da seguinte requisição da API:

Requisição:

```
Método: HTTP GET
```

```
URL: http://{host}:8080/repositories/{repositoryId}
```

Conforme indicado previamente, o número máximo de tags sugeridas para cada repositório é determinado por meio da variável de ambiente **MAX_TAGS_NBR** inserida no arquivo **front-end/.env**. Por exemplo:

```
MAX_TAGS_NBR=<Número máximo de tags>
```

Para que as tags sejam relacionadas aos repositórios com sucesso elas devem ser inseridas na caixa de texto e separadas por vírgula (,) e espaço simples. Por exemplo: 

```
go, golang
```

Em contrapartida, para que nenhuma alteração seja realizada com relação a edição de tags, basta clicar no botão **Cancelar** ou ainda no ícone X do canto superior a direita.

Para descrever a edição de tags em maiores detalhes, suponha que um determinado repositório não tenha qualquer tag relacionada. Como consequência o campo **tag_id** da associação entre usuário e o repositório na tabela **user_repository_tag** é vazio.

A criação de uma tag ocorre por meio da seguinte requisição da API:

Requisição:

```
Método: HTTP POST
```

```
URL: http://{host}:8080/users/{userId}/repositories/{repositoryId}/tag
```

```
Body {
    name: <Nome da tag>
}
```

Durante a criação de uma tag, caso uma palavra inserida na caixa de texto para edição de tags não esteja relacionada a qualquer repositório na tabela **user_repository_tag**, ou seja, ela é uma nova palavra que ainda não foi registrada na tabela **tag**, primeiramente é gerado um identificador único para tal palavra.

Em seguida, para identificar o relacionamento entre a tag e o repositório associado ao usuário do Github, um registro é criado na tabela **user_repository_tag** e o campo **tag_id** é preenchido com o identificar único da tag. Além disso, uma vez que ela é uma nova palavra, também é criado um registro na tabela **tag** utilizando o identificar único e o nome da tag.

Por outro lado, caso uma tag adicionada na caixa de texto para edição de tags já esteja relacionada a qualquer outro repositório na tabela **user_repository_tag**, independentemente do usuário do Github associado, um registro é criado na tabela **user_repository_tag** e o campo **tag_id** é preenchido com o identificador único da tag já existente na tabela **tag**. 

Finalmente, caso uma tag que já estava relacionada a um repositório é removida da caixa de texto para edição de tags, isto significa que tal palavra não deve estar mais relacionada ao repositório e, por conta disso, tal relacionamento precisa ser removido.

A remoção de uma tag ocorre por meio da seguinte requisição da API:

Requisição:

```
Método: HTTP DELETE
```

```
URL: http://{host}:8080/users/{userId}/repositories/{repositoryId}/tags/{tagId}
```

Durante a remoção de uma tag, o registro que representa o relacionamento entre a tag e o repositório associado ao usuário do Github é removido da tabela **user_repository_tag**. Posteriormente, com o intuito de evitar o armazenamento de dados que não são utilizados pelo projeto, caso a tag não esteja relacionada a outro repositório ela também é removida da tabela **tag**.

Neste contexto, quando todas as tags relacionadas a um repositório precisam ser removidas, isto é, todas as palavras que aparecem na caixa de texto para edição de tags são apagadas, todos os registros de relacionamento entre elas e o repositório associado ao usuário do Github são removidos da tabela **user_repository_tag**.

Contudo, apesar da remoção dos registros anteriores, ao final **apenas** um registro é criado na tabela **user_repository_tag** com o campo **tag_id** vazio identificando que o repositório **ainda** está associado ao usuário do Github, porém não existe tag alguma relacionada.

### 7.3 Página de erro

Finalmente, uma página de erro é apresentada ao usuário quando o mesmo tenta acessar um endereço inválido, ou seja, um endereço diferente daqueles demonstrados anteriormente. Por exemplo:

```
http://{host}:3000/<Endereço inválido>
```