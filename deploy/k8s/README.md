# Kubernetes マニフェスト（観測）

連携デモ用の最小構成です。Collector、Tempo、Grafana です。Prometheus / Loki のフルスタックは Compose で確認してください。

起動は [pf-cloud-k8s](https://github.com/maeplego/pf-cloud-k8s) からです。このフォルダだけを apply しないでください。Collector イメージは `0.88.0` です（新しい distroless が一部の Docker Desktop で起動しないため）。
