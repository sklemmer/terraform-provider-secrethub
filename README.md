<p align="center">
  <a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" alt="Terraform" width="330px">
  </a>
  <img width="30px"/>
  <a href="https://secrethub.io">
    <img src="https://secrethub.io/img/secrethub-logo.svg" alt="SecretHub" width="360px">
  </a>
</p>
<h1 align="center">
  <i>Provider<sup><a href="#beta">BETA</a></sup></i>
</h1>

[![GoDoc](https://godoc.org/github.com/secrethub/terraform-provider-secrethub?status.svg)][godoc]
[![CircleCI](https://circleci.com/gh/secrethub/terraform-provider-secrethub.svg)][circleci]
[![Go Report Card](https://goreportcard.com/badge/github.com/secrethub/terraform-provider-secrethub)][goreportcard]
[![Version]( https://img.shields.io/github/release/secrethub/terraform-provider-secrethub.svg)][latest-version]
[![Discord](https://img.shields.io/badge/chat-on%20discord-7289da.svg?logo=discord)][discord]

The Terraform SecretHub Provider lets you manage your secrets using Terraform.

> [SecretHub](https://secrethub.io) is a developer tool to help you keep database passwords, API tokens, and other secrets out of IT automation scripts.

## Usage

```hcl
provider "secrethub" {
  # pass in credential or set SECRETHUB_CREDENTIAL enviroment variable
  credential = "${file("~/.secrethub/credential")}" 
}

resource "secrethub_secret" "db_password" {
  path = "my-org/my-repo/db/password"

  generate {
    length  = 22
    use_symbols = true
  }
}

resource "secrethub_secret" "db_username" {
  path = "my-org/my-repo/db/username"
  value = "db-user"
}

resource "aws_db_instance" "default" {
  allocated_storage    = 10
  storage_type         = "gp2"
  engine               = "mysql"
  engine_version       = "5.7"
  instance_class       = "db.t2.micro"
  name                 = "mydb"
  username             = "${secrethub_secret.db_username.value}"
  password             = "${secrethub_secret.db_password.value}"
  parameter_group_name = "default.mysql5.7"
}
```

Have a look at the [reference docs](/website/docs) for more information on the supported resources and data sources.

## [Get Started]((https://secrethub.io/docs/terraform/))

Check out the [step-by-step integration guide](https://secrethub.io/docs/terraform/) to get started.

A detailed use case is described in the [beta announcement](https://secrethub.io/blog/secret-management-for-terraform/).
There are also some [examples](/examples) in this repo.

## BETA

This project is [currently in beta](https://secrethub.io/blog/secret-management-for-terraform/) and we'd love your feedback! Check out the [issues](https://github.com/secrethub/terraform-provider-secrethub/issues) and feel free suggest cool ideas, use cases, or improvements.

Because it's still in beta, you can expect to see some changes introduced. Pull requests are very welcome.

For support, send us a message on the `#terraform` channel on [<img src="https://discordapp.com/assets/2c21aeda16de354ba5334551a883b481.png" alt="Discord" width="20px"> Discord](https://discord.gg/wcxV5RD) or send an email to [terraform@secrethub.io](mailto:terraform@secrethub.io)

## Development

### Building

Get the source code:

```
git clone https://github.com/secrethub/terraform-provider-secrethub
```

Build it using:

```
make build
```

### Testing

To run the [acceptance tests](https://www.terraform.io/docs/extend/testing/acceptance-tests/index.html), the following environment variables need to be set up.

* `SECRETHUB_CREDENTIAL` - a SecretHub credential.
* `SECRETHUB_TF_ACC_NAMESPACE` - a namespace registered on SecretHub. Make sure `SECRETHUB_CREDENTIAL` has admin access.
* `SECRETHUB_TF_ACC_REPOSITORY` - a repository within `SECRETHUB_TF_ACC_NAMESPACE` to be used in the acceptance tests. Make sure `SECRETHUB_CREDENTIAL` has admin access.

With the environment variables properly set up, run:

```
make testacc
```

[godoc]: https://godoc.org/github.com/secrethub/terraform-provider-secrethub
[circleci]: https://circleci.com/gh/secrethub/terraform-provider-secrethub
[discord]: https://discord.gg/wcxV5RD
[latest-version]: https://github.com/secrethub/terraform-provider-secrethub/releases/latest
[goreportcard]: https://goreportcard.com/report/github.com/secrethub/terraform-provider-secrethub
