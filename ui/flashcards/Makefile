default: build

build: build-latest build-alpine

build-latest:
	docker build -t johnmcdnl/flashcards:latest .

build-alpine:
	docker build --build-arg node_version=alpine -t johnmcdnl/flashcards:alpine .

prod:
	docker build -t johnmcdnl/flashcards:latest .

run:
	docker run -p 3000:3000 johnmcdnl/flashcards:latest

run-alpine:
	docker run -p 3000:3000 johnmcdnl/flashcards:alpine
