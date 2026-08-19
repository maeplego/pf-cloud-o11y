# ローカル Compose（観測）

Collector、Prometheus、Loki、Tempo、Grafana、demo-api です。本番向けではありません。

```powershell
copy .env.example .env
docker compose -f compose.yaml --env-file .env up --build
```

Grafana は `127.0.0.1:3000` だけにバインドします。共有マシンでは `.env` のパスワードを変えてください。保持は設定上おおよそ 72 時間です。
