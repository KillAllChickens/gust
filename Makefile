BINARY := gust
CMD    := ./cmd/gust

.PHONY: build run test vet fmt tidy clean

build:
	go build -o bin/$(BINARY) $(CMD)

run:
	go run $(CMD) $(ARGS)

test:
	go test -race ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

tidy:
	go mod tidy

clean:
	rm -rf bin
