run:
	docker-compose up -d
stop:
	docker-compose down
build:
	go build .
	docker build -t syamsuldocker/messaging-api .
	docker image prune -f
push:
	docker push syamsuldocker/messaging-api