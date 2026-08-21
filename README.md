# pf-cloud-o11y

学習用の可観測性スタックです。アプリは OpenTelemetry で Collector に送り、メトリクス・ログ・トレースを Grafana で見ます。**本番 SRE 基盤の置き換えではありません。**

| ディレクトリ | 役割 |
| --- | --- |
| `apps/demo-api` | 計装の見本（ヘルス、負荷、障害注入） |
| `deploy/` | Collector、Prometheus、Loki、Tempo、Grafana |

Kubernetes 連携は [pf-cloud-k8s](https://github.com/maeplego/pf-cloud-k8s) です。Terraform は [pf-cloud-aws](https://github.com/maeplego/pf-cloud-aws)（**AWS へ apply しません**）。

## 起動

```powershell
copy deploy\.env.example deploy\.env
docker compose -f deploy/compose.yaml --env-file deploy/.env up --build
```

| URL | 用途 |
| --- | --- |
| http://localhost:8088 | demo-api |
| http://localhost:3020 | Grafana（ユーザー `admin`。パスワードは `.env`） |
| http://localhost:9090 | Prometheus |

Grafana には **Demo API RED** ダッシュボードが入っています。

## 障害を入れて見る

`ENABLE_DEBUG=true` のときだけ有効です。

```powershell
curl -X POST "http://localhost:8088/debug/slow?ms=800"
curl -X POST http://localhost:8088/debug/fail
curl http://localhost:8088/work/abc-123
```

Grafana で p95 の上昇、5xx、トレースを確認します。データはローカルだけです。保持は数日分の想定です。

## 他アプリへの接続

- `GET /health` と `GET /ready`
- `OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4318`
- JSON ログに `service`、`trace_id`、`span_id`

設計の詳細は [portfolio-plan](https://github.com/maeplego/portfolio-plan) の `portfolio-plan/cloud-platform/docs/` です。

## ライセンスと利用条件

本リポジトリは **デモ・学習・社内評価用** です。現状品質に **保証はありません**。

- 許可: クローン、ローカル実行、学習、非本番の評価
- 別契約が必要: 本番運用、有償サービスへの組込み、再販・托管の提供

詳細は [LICENSE](./LICENSE) と [licensing.md](https://github.com/maeplego/portfolio-plan/blob/master/portfolio-plan/licensing.md) を参照してください。

