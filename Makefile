build:
	go build -o ./bin/env -v ./cmd/qvm/.

clean:
	rm -rf ./bin
