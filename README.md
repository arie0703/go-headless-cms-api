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
