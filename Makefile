
include .env

APP_NAME=feedback-api-v1
APP_POST=3333

CONTAINER=feedback-db

POSTGRES_VERSION=postgres:12-alpine
POSTGRES_VOLUME=pgdata:/var/run/postgresql/data


db.cotainer:
	docker run --name ${CONTAINER} -p ${DB_PORT}:${DB_PORT} -e POSTGRES_PASSWORD=${DB_PASSWORD} -e POSTGRES_USER=${DB_USER} -d ${POSTGRES_VERSION}

db.create:
	docker exec -it ${CONTAINER} createdb --username=root --owner=root ${DB_NAME}

container.start:
	docker start ${CONTAINER}

container.stop:
	@echo "Stop ${CONTAINER} container..."
	if [ $$(docker ps -q) ]; then \
		echo "Found and stop ${CONTAINER} container..."; \
		docker stop $$(docker ps -q); \
	else \
		echo "No active ${CONTAINER} containers!"; \
	fi

migrations.create:
	sqlx migrate add -r init

migrations.up:
	sqlx migrate run --database-url ${SUPPORT_SERVICE_CONNECT_DATABASE_URL}

migrations.down:
	sqlx migrate revert --database-url ${SUPPORT_SERVICE_CONNECT_DATABASE_URL}

binary.build:
	if [ -f "./build/${APP_NAME}" ]; then \
		rm "./build/${APP_NAME}"; \
		echo "Deleted ${APP_NAME}"; \
	fi
	@echo "Building ${APP_NAME}..."
	go build -o ./build/${APP_NAME} cmd/main.go
	@echo "App ${APP_NAME} build!"

binary.run: binary.build container.stop container.start
	./build/${APP_NAME}

start: binary.run

stop:
	@echo "Stopping ${APP_NAME}..."
	@-pkill -SIGTERM -f "./build/${APP_NAME}"
	@echo "${APP_NAME} stoped!"
