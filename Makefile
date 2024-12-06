COMPOSE_FILE="./development/compose.yaml"

.PHONY: up

up:
	@docker compose -f ${COMPOSE_FILE} up

up-dev:
	@docker compose -f ${COMPOSE_FILE} up --watch

down:
	@docker compose -f ${COMPOSE_FILE} down
	@docker container prune -f
	@docker image prune -f
	@docker rmi atehere-app

prune:
	@docker system prune -af

prune-volume:
	@docker volume prune -af