echo 'running docker compose...'
cd ../..
sh compose.sh

echo 'exporting environments variables...'
cd src
export $(xargs < .env)

echo 'building application...'
go build main.go

echo 'running application...'
sleep 5
./main &
sleep 2

echo 'running integration tests...'
cd ../tests/integration
java -cp karate.jar com.intuit.karate.Main .