build:
	go build -o ./bin/qvm -v ./cmd/qvm/.

clean:
	rm -rf ./bin
