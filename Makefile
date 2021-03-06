build:
	protoc -I. --gofast_out=plugins=grpc:. \
		src/proto/tfc-cap-updater.proto
	#GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o tfc-cap-updater -a -installsuffix cgo main.go repository.go handler.go datastore.go reader.go

local:
	protoc -I. --gofast_out=plugins=grpc:. \
		proto/tfc/cap/updater/tfc-cap-updater.proto
	go build -o tfc-cap-updater -a -installsuffix cgo main.go repository.go handler.go datastore.go reader.go

run:
	docker run -p 50051:50051 -e SERVER_PORT=:50051 tfc/tfc-cap-updater