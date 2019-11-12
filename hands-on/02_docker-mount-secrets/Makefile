build:
	go build server.go

docker-build: ## builds a container with our go webserver
	docker build -t noris.de/secret:demo .


docker-run:  ## run the server in the container
	docker run -it --name secret --env dockersec=foo --mount type=bind,source="$$(pwd)/secret",target="/var/run/dockersec",readonly,bind-propagation=private --rm -p 127.0.0.1:1323:1323 --rm noris.de/secret:demo

inspect:
	docker inspect secret

reveal-secret1:  ## show that secret is no longer protected
	curl 127.0.0.1:1323/secret1

reveal-secret2:  ## show that environment variables are no longer protected
	curl 127.0.0.1:1323/secret2

docker-bad-build:
	docker build -t noris.de/secret_bad:demo -f Dockerfile_bad .

docker-bad-run:
	docker run -d --name secret --env dockersec=foo -p 127.0.0.1:1323:1323 noris.de/secret_bad:demo

get-bad-secret:
	curl localhost:1323/secret1

secret-in-image:
	docker container stop secret
	docker container rm secret
	docker save noris.de/secret_bad:demo -o secret.tar

	# now use the command dive to browse the image with `dive noris.de/secret_bad:demo`
	# use a graphical archive browser to extract the secret from the layer
