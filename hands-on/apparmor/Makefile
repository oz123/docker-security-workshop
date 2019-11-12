build:
	go build server.go

docker-build: ## builds a container with our go webserver
	docker build -t noris.de/apparmor:demo .


docker-run:  ## run the server in the container
	docker run -it --rm -p 127.0.0.1:1323:1323 --rm noris.de/apparmor:demo


hello:  ## show that docker and the program work
	curl 127.0.0.1:1323


oops:  ## show that the hack works
	cat /etc/shadow || echo "Good! we can't see stuff we should not!"
	curl 127.0.0.1:1323/hack && echo "OOPS!!! We *CAN* see stuff we should NOT!"


oops-i-did-it-again:
	curl 127.0.0.1:1323/hack2 && echo "\nOOPS!!! Something happend, which should not"


genrate-apparmor:
	sudo aa-genprof ./server
	# in a different window
	# curl localhost:1323 - allow
	# curl localhost:1323/hack deny
	# finish the server, answer the questions to create apparmor profile
	# now start the server again, / should work, /hack should not

docker-run-with-apparmor:
	docker run -it --name apparmor --rm -p 127.0.0.1:1323:1323 --rm --security-opt apparmor="/home/vagrant/go/src/apparmor/server" noris.de/apparmor:demo

	# now exec a shell in the container
	# docker exec -it apparmor /bin/bash
	# nothing should work here! except the webserver
	# curl to / or /hack still works though
