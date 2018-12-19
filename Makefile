APPLICATION := pace
PROJECT := github.com/242617/$(APPLICATION)
SERVER_ADDRESS := ":8080"

clear:
	clear

run: build
	docker run \
		-p 8080:8080 \
		242617/pace:latest

build: clear
	docker build \
		--build-arg APPLICATION=${APPLICATION} \
		--build-arg PROJECT=${PROJECT} \
		--build-arg SERVER_ADDRESS=${SERVER_ADDRESS} \
		--build-arg DB_CONN=${DB_CONN} \
		--rm \
		-t 242617/pace:latest \
		.

push: build
	docker push 242617/pace:latest

all: run