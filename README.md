# Swipe

Swipe (code: `swipe-api`) is a api server for [Swipe App](https://github.com/swclabs/swipe-app) & [Swipe Admin](https://github.com/swclabs/swipe-admin). `swipe-api` provides functions, and services through API and microservices. Designed for server, `swipe-api` provides cli commands to run api services and redis-based distributed systems cluster.

## Installing

Before installing, you must install make (Makefile) if you use windows operating system

- Go v1.22.5 or higher
- PostgreSQL
- Redis

Update your environment variables. see [.env.example](./.env.example)

### Monolithic

If you want to use makefile, see Other Command Below

Build applications

```bash
make
```

Run api server

```bash
./bin/swipe s
```

Run worker server

```bash
./bin/swipe w
```

with Docker compose

```bash
make all
```

### Other command

`make s` : run api server in dev mode

`make w` : run worker server in dev mode

`make m` : migrate database

`make d` : generate api documentation

`make dev` : run application on docker but not build

`make dev-b` : build and run application on docker in dev mode & no database

`make dev-down` : remove all application containers from docker

`make all` : run application on docker but not build

`make all-b` : build and run application & database on docker

`make all-down` : remove all application containers from docker

`make db` : start database on docker

`make db-down` : remove database container on docker
