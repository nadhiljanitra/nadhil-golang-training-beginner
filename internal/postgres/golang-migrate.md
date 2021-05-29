# Golang Migrate

This is a readme on how to add golang-migare CLI in your local based on golang-migrate package

Reference: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#unversioned

## Installation

### Download pre-built binary (Windows, MacOS, or Linux)

[Release Downloads](https://github.com/golang-migrate/migrate/releases)

```bash
$ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
```

### MacOS

```bash
$ brew install golang-migrate
```

### Windows

Using [scoop](https://scoop.sh/)

```bash
$ scoop install migrate
```

### Linux (*.deb package)

```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```
---

# How to migrate

Reference: https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md

##### Migration-up
```bash
$ migrate -path=internal/postgres/migrations -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DATABASE?sslmode=disable" up
```

##### Migration-down
```bash
$ migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_ADDRESS/$POSTGRES_DATABASE?sslmode=disable" -path=internal/postgres/migrations down
```