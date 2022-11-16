SHELL=CMD

# собрать контейнер
build: 
	docker build -t hello:v1 .

# запустить контейнер с проброской порта 
run:
	docker run --env PORT=$(port) --env CONFIG=$(config) --env HOST=$(host) -it --rm -p $(port):$(port) hello:v1 

run_def:
	docker run --env PORT=8888 --env CONFIG="config/local.toml" --env HOST="host.docker.internal" -it --rm -p 8888:8888 hello:v1 