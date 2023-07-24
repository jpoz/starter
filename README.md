# Starter

## Introduction

This project is a robust starter pack for a back-end service built with Go, backed by a PostgreSQL database. It employs GraphQL to facilitate efficient data transfer, using `gqlgen` for schema generation.

## Key Packages

- `github.com/99designs/gqlgen` : For generating GraphQL servers in Go.
- `github.com/charmbracelet/log` : A charming logger for all your logging needs.
- `github.com/go-chi/chi/v5` : Lightweight and feature-rich router for building Go HTTP services.
- `github.com/redis/go-redis/v9` : A Redis client for Golang.
- `gorm.io` : A developer-friendly ORM for handling interactions with your PostgreSQL database.

## Getting Started

### Prerequisites

- Go 1.18 or higher
- Docker

### Installing

1. Fork the repository, and clone it to your machine

```sh
git clone https://github.com/jpoz/starter
```

2. Move into the project directory.

```sh
cd starter
```

3. Rename the original package name to your new name.

```sh
grep -rl 'github.com/jpoz/starter' ./ | LC_ALL=C xargs sed -i '' 's/github\.com\/jpoz\/starter/github.com\/you\/your_new_project/g'
```

4. Download the required Go dependencies.

```sh
make install
```

4. Setup your database and fill the required information in the `.env` file. Look at the `.env.example`.

5. Run the server locally.

```sh
make dev
```

Now, your server should be running at `localhost:8080`. (or what ever ADDR you set in your `.env` file.

## Deployment

You can build the project using the standard Go build tool. This will create a binary file that can be executed.

```sh
go build -o main .
```

## License

This project is licensed under the MIT License - see the `LICENSE.md` file for details.

## Acknowledgments

This project wouldn't be possible without these wonderful projects and their contributors:

- [GQLGen](https://github.com/99designs/gqlgen)
- [Charm Log](https://github.com/charmbracelet/log)
- [Chi](https://github.com/go-chi/chi)
- [Go-Redis](https://github.com/redis/go-redis)
- [GORM](https://gorm.io)

Please feel free to contribute to this project, report bugs and issues, and suggest improvements.
