# refresh the vault container

docker-compose pull vault && docker-compose up -d --no-deps --build vault
