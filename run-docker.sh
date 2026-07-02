#!/usr/bin/env bash
# Start iotex-analyser-api in Docker against a remote analyzer Postgres.
# Prereqs: image built (docker build -t iotex-analyser-api:local .) and a
# filled .env.docker (see .env.docker.example).
set -euo pipefail

cd "$(dirname "$0")"

if [ ! -f .env.docker ]; then
  echo "missing .env.docker — copy .env.docker.example and fill in DB creds" >&2
  exit 1
fi

# Ports: HTTP 8889 (what iotex-kit hits), gRPC 8888.
docker rm -f iotex-analyser-api 2>/dev/null || true
docker run -d \
  --name iotex-analyser-api \
  --env-file .env.docker \
  -p 8889:8889 \
  -p 8888:8888 \
  iotex-analyser-api:local

echo "started. logs: docker logs -f iotex-analyser-api"
echo "health: curl http://localhost:8889/healthz"
