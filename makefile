db-start:
	docker run -itd --name postgres \
	--network=host \
	-e POSTGRES_PASSWORD=admin \
	-e POSTGRES_USER=admin \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	postgres

db-stop:
	docker stop postgres
	docker rm postgres

run:
	~/go1.18/go/bin/go build .
	docker build -t syamsuldocker/messaging-api .
	docker run -itd --name messaging-api --network=host -v ${CURDIR}/dev:/app/dev -v ${CURDIR}/prod:/app/prod syamsuldocker/messaging-api ./messaging-api

prod-run:
	cp .env prod/.env
	docker pull syamsuldocker/messaging-api
	docker run -itd --name messaging-api --network=host -v ${CURDIR}/prod:/app/prod syamsuldocker/messaging-api env GIN_MODE=release ./messaging-api

stop:
	docker stop messaging-api
	docker rm messaging-api

vm-start:
	docker run -it --name vm --network=host -v ${CURDIR}/dev:/app/dev syamsuldocker/messaging-api bash

vm-stop:
	docker stop vm
	docker rm vm

ssh:
	ssh -i ~/syamsul.pem ubuntu@ec2-3-0-149-232.ap-southeast-1.compute.amazonaws.com

ship:
	~/go1.18/go/bin/go build .
	docker build -t syamsuldocker/messaging-api .
	scp -i ~/syamsul.pem makefile ubuntu@ec2-3-0-149-232.ap-southeast-1.compute.amazonaws.com:~
	scp -i ~/syamsul.pem prod/.env ubuntu@ec2-3-0-149-232.ap-southeast-1.compute.amazonaws.com:~/.env
	scp -i ~/syamsul.pem nginx/conf/nginx.conf ubuntu@ec2-3-0-149-232.ap-southeast-1.compute.amazonaws.com:~/nginx/conf
	docker push syamsuldocker/messaging-api


# https tools
webserver-start:
	docker run -itd --name nginx --network=host \
	 -v ${CURDIR}/nginx/conf/:/etc/nginx/conf.d/:ro \
	 -v ${CURDIR}/certbot/www:/var/www/certbot/:ro \
	 -v ${CURDIR}/certbot/conf:/etc/nginx/ssl/:ro \
	 nginx:latest
webserver-stop:
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
