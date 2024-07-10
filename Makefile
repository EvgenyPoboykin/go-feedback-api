
include .env

APP_NAME=feedback-api-v1

CONTAINER=feedback-db

POSTGRES_VERSION=postgres:12-alpine
POSTGRES_VOLUME=pgdata:/var/run/postgresql/data

DSN=postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL}

swagger:
	swag init -g ./cmd/main.go -o ./docs

db.container:
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
	migrate create -ext sql -dir migration/ -seq init

migrations.up:
	migrate -path migrations/ -database ${DSN} -verbose up

migrations.down:
	migrate -path migrations/ -database ${DSN} -verbose down

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

dev: 
	go run  cmd/main.go
stop:
	@echo "Stopping ${APP_NAME}..."
	@-pkill -SIGTERM -f "./build/${APP_NAME}"
	@echo "${APP_NAME} stoped!"
