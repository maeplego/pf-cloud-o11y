# pf-cloud-o11y

P02 のローカル可観測性スタックです。**本番 SRE 基盤の置き換えではありません。**

アプリは OpenTelemetry SDK で Collector に送り、メトリクス・ログ・トレースを Grafana で相関します。Terraform / Kubernetes は別リポジトリ（`pf-cloud-aws`, `pf-cloud-k8s`）で後フェーズから載せます。

```
apps/demo-api/     計装サンプル（/health, /ready, /work, 障害注入）
deploy/            Compose（Collector, Prometheus, Loki, Tempo, Grafana）
deploy/k8s/        連携デモ最小（Collector + Tempo + Grafana）。束ね役は pf-cloud-k8s
docs/              計装ガイドライン（他 pf-* アプリ向け）
```

## 起動

```powershell
copy deploy\.env.example deploy\.env
docker compose -f deploy/compose.yaml --env-file deploy/.env up --build
```

| URL | 用途 |
| --- | --- |
| http://localhost:8080 | demo-api |
| http://localhost:3000 | Grafana（既定 `admin` / `.env` のパスワード） |
| http://localhost:9090 | Prometheus |

Grafana には **Demo API RED** ダッシュボードが provision されます。

## 障害注入デモ

`ENABLE_DEBUG=true` のときだけ有効です。

```powershell
curl -X POST "http://localhost:8080/debug/slow?ms=800"
curl -X POST http://localhost:8080/debug/fail
curl http://localhost:8080/work/abc-123
```

Grafana で p95 上昇・5xx・トレースを確認します。

## 他アプリへの接続

`docs/instrumentation.md` を参照。最低限:

- `GET /health`, `GET /ready`
- `OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4318`
- JSON ログに `service`, `trace_id`, `span_id`

## 保持とコスト

ローカルのみ。Loki / Tempo / Prometheus は数日分の保持想定。ディスクはホスト依存です。
