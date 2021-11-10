echo 'running unity tests...'
cd ..
go test ./src/...

cd tests/integration
sh integration.sh

