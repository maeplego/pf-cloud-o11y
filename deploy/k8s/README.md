# P02 observability — Kubernetes（連携デモ最小）

Collector + Tempo + Grafana。フルスタック（Prometheus / Loki / Promtail）は Compose 単体デモで確認する。

Grafana のホームは **Media API traces**（Tempo Explore への案内）。Collector イメージは `0.88.0`（`0.116.x` distroless が一部 Docker Desktop 環境で起動不能なため）。

Collector イメージは `0.88.0`（`0.116.x` distroless が一部 Docker Desktop 環境で起動不能なため）。

Tempo 設定は `tempo-config.yaml`（Compose 版 `deploy/tempo/tempo.yaml` と同等）を k8s 配下に置く。
