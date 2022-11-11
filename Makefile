# собрать контейнер
build:
	docker build -t hello:v1 .
# запустить контейнер
run:
	docker run --env PORT=8888 -it --rm -p 8888:8888 hello:v1
