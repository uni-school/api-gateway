# Uni School: Api Gateway

api gateway is a gatway that calling the microservices.

## Installation

- Step 1: Install all GRPC prerequisites on this [link](https://grpc.io/docs/languages/go/quickstart/#prerequisites)
- Step 2: Install Wire by running

```bash
go install github.com/google/wire/cmd/wire@latest
```

- Step 3: Install Depedencies by Running

```bash
go mod tidy
```

- Step 4: Copy `<dev|stag|prod|test>`.application.yaml.example to `<dev|stag|prod|test>`.application.yaml. `NOTE`: choose type dev

- Step 5: Running the api-gateway

```bash
make start-dev
```

## Documentation

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/10344918-59f09a1f-f08f-4507-a9bc-bccf4cf2ed67?action=collection%2Ffork&collection-url=entityId%3D10344918-59f09a1f-f08f-4507-a9bc-bccf4cf2ed67%26entityType%3Dcollection%26workspaceId%3D43df7931-feec-460c-8889-25210781dc3f)

## NOTE

- Number 1: make sure you do all step in installation guide
- Number 2: make sure in `<env>`.application.yaml variable passwordHashing.saltHash must have same value in file `<env>`.application.yaml on user-microservice
- Number 3: if you specify port of this api gateway, the server should be listening on `<server-host>`:`<server-port>`
