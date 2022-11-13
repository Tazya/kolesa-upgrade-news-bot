SHELL=CMD

# собрать контейнер
build: 
	docker build -t hello:v1 .
# запустить контейнер с проброской порта 
run:
	docker run --env PORT=$(port) -it --rm -p $(port):$(port) hello:v1

