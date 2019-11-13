
docker-build-pypi: clone-pypi
	docker build -t noris.de/pypi:0.1 -f Dockerfile-pypi pypiserver

docker-build-mariadb:
	docker build -t noris.de/mariadb:0.1 -f Dockerfile-pypi .

clone-pypi:
	git clone --depth 1 https://github.com/pypiserver/pypiserver.git
