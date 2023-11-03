build:
	go build -o ./bin/silkania .

run: build
	./bin/silkania