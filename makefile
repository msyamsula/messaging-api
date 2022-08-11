# local
run-local:
	go run .

# cluster
build:
	go build .
	docker build \
	-t syamsuldocker/messaging-api \
	-f ${CURDIR}/env/dev/Dockerfile \
	.
	# docker build \
	# -t syamsuldocker/nginx-api \
	# -f ${CURDIR}/env/dev/Dockerfile.nginx \
	# .

run:
	make build
	docker-compose -f ${CURDIR}/env/dev/docker-compose.yaml up -d --scale messaging-api=${scale}

stop:
	docker-compose -f ${CURDIR}/env/dev/docker-compose.yaml down

ps:
	docker-compose -f ${CURDIR}/env/dev/docker-compose.yaml ps


# ship to production
ship-production:
	go build .
	docker build \
	-t syamsuldocker/messaging-api \
	-f ${CURDIR}/env/prod/Dockerfile \
	.
	docker push syamsuldocker/messaging-api
	scp -i ~/syamsul.pem makefile ubuntu@ec2-3-0-149-232.ap-southeast-1.compute.amazonaws.com:~/
	scp -i ~/syamsul.pem env/prod/default.conf ubuntu@ec2-3-0-149-232.ap-southeast-1.compute.amazonaws.com:~/nginx/

# production
run-production:
	docker pull syamsuldocker/messaging-api
	docker run \
	-itd \
	--name messaging-api \
	--network=host \
	-e GIN_MODE=release \
	syamsuldocker/messaging-api
	docker run \
	-itd \
	--name nginx \
	--network=host \
	-v ${CURDIR}/nginx/:/etc/nginx/conf.d/ \
	-v ${CURDIR}/certbot/conf:/etc/nginx/ssl \
	-v ${CURDIR}/certbot/www:/var/www/certbot/ \
	nginx

stop-production:
	docker stop nginx messaging-api
	docker rm nginx messaging-api

restart-production:
	make stop-production
	make run-production

# ssh
ssh:
	ssh -i ~/syamsul.pem ubuntu@ec2-3-0-149-232.ap-southeast-1.compute.amazonaws.com

# https tools
start-webserver:
	docker run -itd --name nginx --network=host \
	 -v ${CURDIR}/nginx/conf/:/etc/nginx/conf.d/:ro \
	 -v ${CURDIR}/certbot/www:/var/www/certbot/:ro \
	 -v ${CURDIR}/certbot/conf:/etc/nginx/ssl/:ro \
	 nginx:latest
stop-webserver:
	docker stop nginx
	docker rm nginx
certbot-dry-run:
	docker run -it --name certbot --network=host \
	-v ${CURDIR}/certbot/www:/var/www/certbot/:rw \
	-v ${CURDIR}/certbot/conf:/etc/letsencrypt/:rw \
	certbot/certbot:latest certonly --webroot --webroot-path /var/www/certbot/ --dry-run -d syamsulapi.my.id
certbot-create:
	docker run -it --name certbot --network=host \
	-v ${CURDIR}/certbot/www:/var/www/certbot/:rw \
	-v ${CURDIR}/certbot/conf:/etc/letsencrypt/:rw \
	certbot/certbot:latest certonly --webroot --webroot-path /var/www/certbot/ -d syamsulapi.my.id
certbot-stop:
	docker stop certbot
	docker rm certbot
certbot-renew:
	docker run -it --name certbot --network=host \
	-v ${CURDIR}/certbot/www:/var/www/certbot/:rw \
	-v ${CURDIR}/certbot/conf:/etc/letsencrypt/:rw \
	certbot/certbot:latest renew

# kubernetes
kube-build:
	go build .
	docker build -t syamsuldocker/messaging-api-kubernetes:v${version} -f env/kube/Dockerfile .

kube-ship:
	make version=${version} kube-build
	docker push syamsuldocker/messaging-api-kubernetes:v${version}

kube-dev-run:
	make version=${version} kube-build
	docker run -it --name messaging-api -e GIN_MODE=release --network=host syamsuldocker/messaging-api-kubernetes:v${version}

kube-stop:
	docker stop messaging-api
	docker rm messaging-api