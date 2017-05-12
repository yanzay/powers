build:
	go build -v .
linuxbuild:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v .
docker: linuxbuild
	docker build -t yanzay/powers .

