echo 'running unit tests...'
cd ..
go test ./src/...

if [ $? = 1 ]; then
  exit 1
fi

echo 'running integration tests...'
cd tests/integration
sh integration.sh

