# Headless CMS (Go + MySQL)

GoとMySQLを使用して構築されたヘッドレスCMS API

---

## 📝 MEMO

### 前提条件

- Go 1.24.1
- Docker / Docker Compose
- (任意) `goose` インストール: `go install github.com/pressly/goose/v3/cmd/goose@latest`

---

### `.env` ファイルの作成

`.env.example`の内容をコピーして`.env`を作成してください

### サービス起動

```bash
make up
```

### マイグレーション

```bash
make create-migration
```

### docker compose build with logs

`COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker compose build --progress=plain`

## API Request Examples

```bash
# get all users
curl -X GET http://localhost:8080/api/users

# create user
curl -X POST http://localhost:8080/api/users \\n  -H "Content-Type: application/json" \\n  -d '{\n    "name": "Alice",\n    "email": "alice@example.com",\n    "password_hash": "hashedpassword123",\n    "role": "editor"\n  }'

# update user
curl -X PUT http://localhost:8080/api/users/767f3e16-f800-4888-825d-2b7a15740b62 \\n  -H "Content-Type: application/json" \\n  -d '{\n    "name": "Alice Updated",\n    "email": "alice.new@example.com",\n    "password_hash": "newhash",\n    "role": "admin"\n  }'

# delete user
curl -X DELETE http://localhost:8080/api/users/767f3e16-f800-4888-825d-2b7a15740b62

```
