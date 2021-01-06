GIT_SHA?=$(shell git log -n 1 | head -n 1 | awk '{print $$2}')
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-extldflags "-static" -X main.GitSHA=$(GIT_SHA) -X main.BuildTime=$(BUILD_TIME)"

build:
	go build $(LDFLAGS) -o ./bin/qvm -v ./cmd/qvm/.

clean:
	rm -rf ./bin

qasm3:
	$(MAKE) -C ./internal $@

qasm3-go:
	go build -o ./bin/qasm3 -v ./cmd/qasm3/.

.PHONY: cmd-qasm
gen-qasm: qasm3 qasm3-go
