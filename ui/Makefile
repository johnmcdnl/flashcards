default: build

build: build-latest build-alpine build-prod build-prod-alpine

build-latest:
	docker build -t johnmcdnl/flashcards:latest .

build-alpine:
	docker build --build-arg node_version=alpine -t johnmcdnl/flashcards:alpine .

build-prod:
	docker build -t johnmcdnl/flashcards:prod -f Dockerfile.prod .

build-prod-alpine:
	docker build --build-arg nginx_version=alpine -t johnmcdnl/flashcards:prod-alpine -f Dockerfile.prod .

run:
	docker run -p 3000:3000 johnmcdnl/flashcards:latest

run-alpine:
	docker run -p 3001:3000 johnmcdnl/flashcards:alpine

prod:
	docker run -p 5000:80 johnmcdnl/flashcards:prod

prod-alpine:
	docker run -p 5001:80 johnmcdnl/flashcards:prod-alpine
