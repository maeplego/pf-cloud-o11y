# P02 observability — Kubernetes（連携デモ最小）

Collector + Tempo + Grafana。フルスタック（Prometheus / Loki / Promtail）は Compose 単体デモで確認する。

Tempo 設定は `tempo-config.yaml`（Compose 版 `deploy/tempo/tempo.yaml` と同等）を k8s 配下に置く。
