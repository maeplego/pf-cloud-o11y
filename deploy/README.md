# deploy

Local observability stack. Not for production.

```powershell
copy .env.example .env
docker compose -f compose.yaml --env-file .env up --build
```

Grafana binds to `127.0.0.1:3000` only. Change password in `.env` before sharing a machine.

Retention: Loki/Tempo ~72h in config. Disk use grows with traffic.
