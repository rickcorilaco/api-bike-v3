echo 'running docker compose...'
cd ../..
sh compose.sh
if [ $? = 1 ]; then
  exit 1
fi

echo 'exporting environments variables...'
cd src
export $(xargs < .env)

echo 'building application...'
go build main.go
if [ $? = 1 ]; then
  exit 1
fi
sleep 5

echo 'running application...'
./main &
if [ $? = 0 ]; then
  pid=$!
  sleep 5
else
  exit 1
fi

echo 'checking application...'
kill -0 $pid
if [ $? = 1 ]; then
  exit 1
fi

echo 'running integration tests...'
cd ../tests/integration
java -cp karate.jar com.intuit.karate.Main .

echo 'closing application...'
kill $pid