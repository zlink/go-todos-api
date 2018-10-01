# Makefile for goalng-todo-api


APP_NAME = "todos"

default:
	go build -o ${APP_NAME} main.go

install:
	govendor sync -v

dev:
	fresh -c configs/development.conf

clean:
	if [ -f ${APP_NAME} ]; then rm ${APP_NAME}; fi