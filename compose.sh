docker-compose down
docker build -t postgres_db ./src/infra/postgres/.
docker-compose build postgres
docker-compose up -d --force-recreate postgres
