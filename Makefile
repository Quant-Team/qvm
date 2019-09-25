GIT_SHA?=$(shell git log -n 1 | head -n 1 | awk '{print $$2}')
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-extldflags "-static" -X main.GitSHA=$(GIT_SHA) -X main.BuildTime=$(BUILD_TIME)"

build:
	go build $(LDFLAGS) -o ./bin/qvm -v ./cmd/qvm/.

clean:
	rm -rf ./bin
