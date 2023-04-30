

deploy:
	prisma1 deploy

gen:
	prisma1 generate

grpc:
	buf generate

user-dev:
	make -C services/user dev

user-proxy-dev:
	make -C services/user dev-proxy