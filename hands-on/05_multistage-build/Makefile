
build:
	echo "No build here! Use the docker-build target!"

docker-build:
	docker build -t noris.de/multi:demo .


docker-run:
	# the server works without all the un-necessary stuff
	# reduces attack surface
	# you can still combine all other techniques like seccomp and apparmor
	docker run -d --name multi --rm -p 127.0.0.1:1323:1323 --rm noris.de/multi:demo
	curl localhost:1323
	docker rm --force multi
