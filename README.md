# service-api

## Start

### conf

CP ./config.example.toml ./config.toml

### Using Default

1. edit `./config.toml`

2. generate token ras

```bash
openssl genpkey -algorithm RSA -out private_key.pem -aes256
openssl rsa -pubout -in private_key.pem -out public_key.pem
```

3. run script update database schema

* win

```bash
go generate .\pkg\ent\gen.go
```

* linux or mac

```bash
go generate ./pkg/ent/gen.go
```

4. start

* default

```bash
go run ./main.go
```

* use air

```bash
air
```

### Useing VsCode Docker

1. edit `.devcontainer/.env`
2. vscode install docker plugin
3. start vscode docker server
4. go run ./src/main.go

