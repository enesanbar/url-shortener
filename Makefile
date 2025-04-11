PROJECT_NAME=url-shortener
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)
BINARY=$(GOBIN)/$(PROJECT_NAME)

clear:
	rm -f $(GOBIN)/$(PROJECT_NAME)

build: clear
build:
	CGO_ENABLED=0 GOOS=linux go build -o $(GOBIN)/$(PROJECT_NAME) $(GOBASE)/cmd/rest/main.go || exit

test-ci: export DEPLOY_TYPE=dev
test-ci: export CONFIG_SOURCE=consul
test-ci:
	gotestsum --junitfile report.xml --format testname -- -coverprofile=coverage.out -coverpkg=./... ./... 2>&1
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out
	sed -e 's%github.com/enesanbar/url-shortener/%%g' coverage.out | grep -v mocks > coverage.new.out
	gocover-cobertura < coverage.new.out > cobertura.xml


start: build
start:
	$(GOBIN)/$(PROJECT_NAME)

graph:
	godepgraph -maxlevel 5 -ignoreprefixes github,go.uber,github.com/enesanbar/go-service  -nostdlib github.com/enesanbar/url-shortener/cmd/rest  | dot -Tpng -o dependency-graph.png
