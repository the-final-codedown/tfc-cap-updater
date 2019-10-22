build:
	protoc -I. --go_out=plugins=micro:. \
		proto/tfc/cap/updater/tfc-cap-updater.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o tfc-cap-updater -a -installsuffix cgo main.go repository.go handler.go datastore.go
	docker build -t tfc/tfc-cap-updater .


run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 tfc-cap-updater