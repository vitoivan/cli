ARGS= .env "VAL1" "another value"
NAME=upsert-dotenv.out
OUT=./

all: run

run:
	@go run src/main.go $(ARGS)

build:
	@go build -o $(NAME) src/main.go
	@mv $(NAME) $(OUT)

prod: build
	./main $(ARGS)

clean:
	rm main
