# Makefile for MaxMind

.PHONY: vendor
vendor:
	@echo "==> Go vendor. Gathering all exteranl dependencies in vendor folder 🎲 <=="
	go mod vendor -v

.PHONY: mock
mock:
	@echo "==> Generating mocks for unit test 🇪🇸 <=="
	go generate ./...

.PHONE: rmmock
rmmock:
	@echo "==> Removing all mock files  <=="
	find . -name 'mock_*' -type f -delete

.PHONY: test
test:
	@echo "==> Running Unit Tests 🎮 <=="
	go test ./... -cover

.PHONY: testmock
testmock:
	@echo "==> Generating mocks and then run unit tests 🏀 <=="
	make mock
	make test

.PHONY: cover
cover:
	@echo "==> Visual coverage for $(FOLDER)"
	go test $(FOLDER) -coverprofile=coverage/lastCoverage.out
	go tool cover -html=coverage/lastCoverage.out

.PHONY: build
build:
	@echo "==> Build: Generate binary on /bin folder <=="
	go build -o ./bin/maxmind -ldflags "-s -w" ./cmd/maxmind.go

.PHONY: cover
cover:
	@echo "==> Visual coverage for $(FOLDER)"
	go test $(FOLDER) -coverprofile=.coverage/lastCoverage.out
	go tool cover -html=.coverage/lastCoverage.out
