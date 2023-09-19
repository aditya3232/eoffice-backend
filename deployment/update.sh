#!/bin/bash
cd /home/bit/build-eoffice-v2/eoffice-backend/
git pull origin main
docker build . --file .docker/Dockerfile --tag eoffice-v2-backend
cd /home/bit/docker-eoffice-v2/eoffice-backend/
docker compose up -d --build -d --force-recreate
docker rmi -f $(docker images -f "dangling=true" -q) || true
