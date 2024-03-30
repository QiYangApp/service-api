# service-api

## Start

### conf

CP ./config.example.yaml ./config.yaml

### Using Default

1. edit `./config.toml`

```bash
openssl genpkey -algorithm RSA -out private_key.pem -aes256
openssl rsa -pubout -in private_key.pem -out public_key.pem
```

2. run script update database schema

```bash
cd .\pkg\ent
go generate .\gen.go
```

3. start

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

## Bug

[ ] framework global config 

