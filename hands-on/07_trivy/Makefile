
docker-build-pypi: clone-pypi
	docker build -t noris.de/pypi:0.1 -f Dockerfile-pypi pypiserver

docker-build-mariadb:
	docker build -t noris.de/mariadb:0.1 -f Dockerfile .

clone-pypi:
	mkdir pypiserver; \
	cd pypiserver; \
	git init; \
	git remote add origin https://github.com/pypiserver/pypiserver.git; \
	git fetch --depth 1 origin afafd0ae509972d5d81530f2a37920359d777637; \
	git checkout FETCH_HEAD
