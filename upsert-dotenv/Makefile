ARGS= .env "VAL1" "another value"

all: run

run:
	@go run src/main.go $(ARGS)

build:
	@go build -o upsert-dotenv src/main.go

prod: build
	./main $(ARGS)

clean:
	rm main
