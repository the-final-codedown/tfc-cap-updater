build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o tfc-cap-cli -a -installsuffix cgo cli.go
	docker build -t tfc-cap-cli .

run:
	docker run -e MICRO_SERVER_ADDRESS=:50051 tfc-cap-cli