.PHONY: build clean

# make build
build:
	env GOOS=linux go build -o bin/main main.go
	zip bin/main.zip bin/main

# make clean
clean:
	rm -rf ./bin

deploy:
	sls deploy
