# 計装ガイドライン（pf-* 共通）

P02 の契約。アプリはベンダー SDK を直接 Grafana / Jaeger / Loki に送らず、**OpenTelemetry Collector** 経由にする。

## 必須エンドポイント

| パス | 用途 |
| --- | --- |
| `GET /health` | liveness（プロセス生存） |
| `GET /ready` | readiness（DB 等の依存。MVP では 200 固定でも可） |

## 環境変数

```env
OTEL_SERVICE_NAME=my-service
OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4318
```

gRPC を使う場合は `:4317`。アプリ README にどちらかを明記する。

## ログ（JSON）

stdout に 1 行 JSON。最低限のキー:

| キー | 内容 |
| --- | --- |
| `service` | `OTEL_SERVICE_NAME` と同じ |
| `trace_id` | 有効 span があるとき |
| `span_id` | 有効 span があるとき |
| `msg` | イベント名 |
| `http.route` | 正規化済みルート（`/users/:id`） |

**禁止**: 生の user id / order id を Prometheus ラベルや Loki ラベルに入れない。

## メトリクス

- HTTP は `otelhttp` 等の公式計装を使う
- ルートはパラメータを `:id` に正規化して `http.route` に載せる
- カスタムカウンタのラベルは 10 種類未満・有限集合に抑える

## トレース

- 入口 HTTP で span を開始
- DB / 外部 HTTP は子 span（ライブラリ計装があればそれを使う）
- Collector → Tempo。アプリから Tempo へ直接送らない

## ローカル開発

1. `pf-cloud-o11y` の Compose を起動
2. アプリの Compose に `OTEL_EXPORTER_OTLP_ENDPOINT=http://host.docker.internal:4318`（Windows/macOS Docker）
3. Grafana http://localhost:3000 で RED とトレースを確認

## pf-identity への適用（将来）

`apps/server` に OTel middleware と JSON slog を足すときは、上記キー名を揃える。ログイン失敗回数などはメトリクスラベルに email を入れない。
