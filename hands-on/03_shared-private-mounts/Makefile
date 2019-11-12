
docker-run-permissive-mounts:
	docker run -d --rm --name mount-permissive -v "$$(pwd)/parent":"/parent" ubuntu:18.04 bash -c "sleep 6000"
	# Now mount something at host (while container is still running):
	mount --bind -o ro /etc parent/child
	# echo foo > /parent/bla
	# echo foo > /parent/child/bla  
	docker exec -it mount-permissive sh
	umount --lazy $$(pwd)/parent/child
	docker rm --force mount-permissive

docker-run-with-mounts:
	docker run -d --rm --name mount --mount type=bind,source="$$(pwd)/parent",target="/parent",readonly ubuntu:18.04 bash -c "sleep 6000" 
	mount --bind -o ro /etc parent/child
	# echo "foo"> /parent/bla should not work
	# echo "foo"> /parent/child/bla should not work
	docker exec -it mount sh
	docker rm --force mount
	umount --lazy $$(pwd)/parent/child


