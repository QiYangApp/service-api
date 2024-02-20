# service-api
 
## generate token rsa

```bash
openssl genpkey -algorithm RSA -out private_key.pem -aes256
openssl rsa -pubout -in private_key.pem -out public_key.pem
```

## Start

## conf

CP ./config.example.toml ./config.toml

### Using Default

1. edit `./config.toml`

2. generate token ras

```bash
openssl genpkey -algorithm RSA -out private_key.pem -aes256
openssl rsa -pubout -in private_key.pem -out public_key.pem
```

3. generate api

```bash
go run ./pkg/framework/script/router/main.go
go run -mod=mod entgo.io/ent/cmd/ent generate "./ent/schema"

```

3. start

```bash
go run ./main.go
```

### Useing VsCode Docker

1. edit `.devcontainer/.env`
2. vscode install docker plugin
3. start vscode docker server
4. go run ./src/main.go


### Testing

open `localhost:3000/admin/test/welcome`


### TODO

[] 调整底层框架结构
[] 调整返回方法
[] 路由中间件验证

