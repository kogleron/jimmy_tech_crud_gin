## How to...

### Run the project

For starting the service in the project dir exec `cp .env.example .env && make docker`.

### Read the docs and test

<http://localhost:8085/swagger/index.html>
or
<http://localhost:YOUR_PORT/swagger/index.html>

Auth token for requests: auth_token_xxx

### Run examples

Examples of http calls can be found in **./examples**

### Generate Swagger docs

1. Pull swagger libraries.
```
make install
```
2. Generate the Swagger Specification.
```
make swagger
```

## Assignment
Headless (no UI) CMS for managing articles (CRUD).

Each article has its ID, title, body and timestamps (on creation and update).

Managing articles is the typical set of CRUD calls including reading one and all the articles. Creating/updating/deleting data is possible only if a secret token is provided (can be just one static token).

For reading all the articles, the endpoint must allow specifying a field to sort by including whether in an ascending or descending order + basic limit/offset pagination.

The whole client-server communication must be in a JSON format and be ready to be extended with other formats (eg. XML).

Keep in mind the best practices for building server applications including automated testing.

## Technical Requirements
* GoLang
* Automated tests
* REST API + documentation
* Relational Database (MSSQL, PostgreSQL, ...) * README
* Docker

## TODO

- [x] CRUD
- [x] Authorization.  Creating/updating/deleting data is possible only if a secret token is provided (can be just one static token).
- [x] For reading all the articles, the endpoint must allow specifying a field to sort by including whether in an ascending or descending order + basic limit/offset pagination.
- [x] The whole client-server communication must be in a JSON format and be ready to be extended with other formats (eg. XML).
- [x] Fix TODOs
- [x] Format
- [x] Lint
- [x] Docker
- [x] Instructions in README.md
- [x] Swagger
- [x] Automated tests
- [x] Squash commits
- [ ] Fix models' examples in XML format
