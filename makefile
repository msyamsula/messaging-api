run:
	docker-compose up -d
stop:
	docker-compose down
build:
	go build .
	docker build -t syamsuldocker/messaging-api .
	docker image prune -f
build-prod:
	go build .
	docker build -t syamsuldocker/messaging-api:v0.0.0 .
	docker image prune -f
push:
	docker push syamsuldocker/messaging-api:v0.0.0

#non-container run
non-container-run:
	gin -p 3000 -a 5000 -b messaging-api run main.go