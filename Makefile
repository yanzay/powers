build:
	go build -v -i .
dev: build
	./powers --local --log-level trace
test:
	go test -v
	rm .test.db
linuxbuild:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -i -v .
docker: linuxbuild
	docker build -t yanzay/powers .
push: docker
	docker push yanzay/powers
deploy: push
	ssh root@yanzay.com "cd infra; docker-compose pull bot; docker-compose up -d"
clean:
	rm powers
	rm powers.db
