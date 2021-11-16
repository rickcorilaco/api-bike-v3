docker-compose down
if [ $? = 1 ]; then
  exit 1
fi

docker build -t postgres_db ./src/infra/postgres/.
if [ $? = 1 ]; then
  exit 1
fi

docker-compose build postgres
if [ $? = 1 ]; then
  exit 1
fi

docker-compose up -d --force-recreate postgres
if [ $? = 1 ]; then
  exit 1
fi