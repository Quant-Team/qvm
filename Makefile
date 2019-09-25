build:
	go build -o ./bin/env -v ./cmd/qvm/.
	chmod $(UID):$(GID) -R ./bin go.mod go.sum

clean:
	rm -rf ./bin
