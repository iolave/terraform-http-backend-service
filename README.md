# Terraform: HTTP backend service

Service used to initialize terraform and store `terraform.tfstate` files in a desired storage by specifying a project id and an environment.

> [!WARNING]
> All versions released prior to `v1.0.0` are to be considered [breaking changes](https://semver.org/#how-do-i-know-when-to-release-100).

This service exposes the `/tf-state/{projectId}/{env}` with the following behaviour:
- `GET`: Retrieves `terraform.tfstate` content from the storage.
- `POST`: Replaces `terraform.tfstate` content with the one of the request body.

## Requirements

- [Go](https://golang.org/doc/install) >= 1.22.2

## Building the service

1. Clone the repository
2. Enter the repository directory
3. Install go dependencies by running the go `install` command:
```shell
go install
```
5. Build the service using the makefile `build` recipe:
```shell
make build
```

## Running the service
Once the service is built, run the executable from the `dist` dir that matches your platform. Alternatively use the makefile `serve` recipe.
```shell
make serve
``` 

### Enviroment variables

| Variable       | Description |
|----------------|-------------|
| `STORAGE_PATH` | Local path used to store tfstate files. If not defined, the service will respond with an internal server error |