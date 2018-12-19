APPLICATION := pace
PROJECT := github.com/242617/$(APPLICATION)
SERVER_ADDRESS := ":8080"

clear:
	clear

run: build
	docker run \
		-p 8080:8080 \
		pace:latest

build: clear
	docker build \
		--build-arg APPLICATION=${APPLICATION} \
		--build-arg PROJECT=${PROJECT} \
		--build-arg SERVER_ADDRESS=${SERVER_ADDRESS} \
		--rm \
		-t pace:latest \
		.

all: run