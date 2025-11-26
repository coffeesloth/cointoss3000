.PHONY: *

build:
	go build -o ct3 main.go

docker-build: build
	docker build --no-cache=true -t cointoss3000:latest .

docker-run: docker-build
	docker run -p 8080:8080 cointoss3000:latest

docker-save: docker-build
	docker image save cointoss3000:latest | gzip > cointoss3000-docker.tgz

clean:
	rm -f ct3 cointoss3000-docker.tgz
