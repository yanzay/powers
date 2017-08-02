build:
	go build -v -i .
dev: build
	./powers --local --log-level trace
test:
	go test -v
	rm .test.db
clean:
	rm powers
	rm powers.db
