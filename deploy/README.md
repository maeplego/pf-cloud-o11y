# ローカル Compose（観測）

Collector、Prometheus、Loki、Tempo、Grafana、demo-api です。本番向けではありません。

```powershell
copy .env.example .env
docker compose -f compose.yaml --env-file .env up --build
```

Grafana は `3020` にバインドします（Windows の Hyper-V 予約で `3000` が使えない環境向け）。共有マシンでは `.env` のパスワードを変えてください。保持は設定上おおよそ 72 時間です。
