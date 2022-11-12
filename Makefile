SHELL=CMD

# собрать контейнер с проброской порта
build: 
	docker build --build-arg PORT=$(port) -t hello:v1 .
# запустить контейнер с проброской порта 
run:
	docker run --env PORT=$(port) -it --rm -p $(port):$(port) hello:v1

