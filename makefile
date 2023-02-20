ship:
	go build -o messaging-api binary/http.go
	docker build -t syamsuldocker/syamsulapp-http:${TAG} .
	docker push syamsuldocker/syamsulapp-http:${TAG}