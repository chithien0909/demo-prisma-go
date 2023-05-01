

deploy:
	prisma1 deploy

gen-1:
	prisma1 generate

gen:
	go run github.com/prisma/prisma-client-go generate

grpc:
	buf generate

user-dev:
	make -C services/user dev

user-proxy-dev:
	make -C services/user dev-proxy