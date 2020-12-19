# refresh the redis container

docker-compose pull redis && docker-compose up -d --no-deps --build redis
