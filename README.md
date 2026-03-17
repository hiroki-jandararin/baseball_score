# baseball-score-app

野球スコア入力を素早く行うための Web アプリです。

## 技術構成

- Frontend: React
- Backend: Go
- Database: SQLite
- Infra (MVP): Vercel + Lightsail or EC2
- Local Dev: Docker Compose

## ディレクトリ構成

```text
baseball-score-app/
  ├─ frontend/
  ├─ backend/
  ├─ docker-compose.yml
  ├─ .gitignore
  ├─ .env.example
  └─ README.md
```

## Docker 開発時の基本操作

### 起動

初回起動またはイメージを作り直したい場合:

```bash
docker compose up --build
```

通常起動:

```bash
docker compose up
```

- Frontend: `http://localhost:5173`
- Backend: `http://localhost:8080`
- Backend health check: `http://localhost:8080/health`

### バックグラウンド起動

```bash
docker compose up -d
```

### 停止

```bash
docker compose down
```

コンテナと一緒に匿名ボリュームも削除したい場合:

```bash
docker compose down -v
```

### ログ確認

全体ログ:

```bash
docker compose logs -f
```

Frontend のみ:

```bash
docker compose logs -f frontend
```

Backend のみ:

```bash
docker compose logs -f backend
```

### コンテナ状態確認

```bash
docker compose ps
```

### 再ビルド

イメージを再ビルドして起動:

```bash
docker compose up --build
```

キャッシュを使わずにビルドしたい場合:

```bash
docker compose build --no-cache
docker compose up
```

### コンテナ内でコマンド実行

Frontend コンテナでシェルを開く:

```bash
docker compose exec frontend sh
```

Backend コンテナでシェルを開く:

```bash
docker compose exec backend sh
```

### データの保存先

この開発環境では Database に SQLite を使っており、DB ファイルは次の場所に保存されます。

```text
./backend/data/app.db
```

`docker compose down` だけではこのファイルは消えません。DB を初期化したい場合は、コンテナ停止後に `backend/data/app.db` を削除してください。
