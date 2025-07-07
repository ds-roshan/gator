# gator


This document provides instructions on how to install and configure the Gator CLI tool. Please follow the steps below to get started.

## Prerequisites

Before running the program, ensure you have the following installed:

- **Postgres**: Gator uses Postgres as its database. You can download Postgres from [the official website](https://www.postgresql.org/download/) and follow the installation instructions for your operating system.
- **Go**: The Gator CLI is written in Go. Install Go by following the instructions on the [official Go website](https://golang.org/dl/).

## Installing the Gator CLI

Once you have Go installed, you can install the Gator CLI using the `go install` command. Open your terminal and run:

```sh
go install github.com/ds-roshan/gator@latest
```

This command downloads the latest version of the Gator CLI and installs it to your Go binary path.

## Configuring the Gator CLI

The Gator CLI requires a configuration file in order to connect to your Postgres database. Create a JSON configuration file named `.gatorconfig.json` in the root directory with the following format:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

Make sure you update the `db_url` with your database details.
