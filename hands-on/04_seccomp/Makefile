build:
	go build server.go

docker-build: ## builds a container with our go webserver
	docker build -t noris.de/seccomp:demo .

docker-run:
	docker run -d --name seccomp --rm -p 127.0.0.1:1323:1323 --rm noris.de/seccomp:demo
	curl localhost:1323
	curl localhost:1323/hack
	curl localhost:1323/hack2
	docker rm --force seccomp

docker-run-with-seccomp:
	# getcwd syscall wird nicht mehr erlaubt!
	docker run -d --rm --name seccomp -p 127.0.0.1:1323:1323 --security-opt seccomp="$$(pwd)/profile_bad.json" --rm noris.de/seccomp:demo
	curl localhost:1323/hack2
	# no way to stop open or openat with seccomp, because this is needed to read every file, even the binary program
	curl localhost:1323/hack
	docker rm --force seccomp
