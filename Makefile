COMPOSE_FILE="./development/compose.yaml"

.PHONY: up

up:
	@docker compose -f ${COMPOSE_FILE} up --watch

down:
	@docker compose -f ${COMPOSE_FILE} down
	@docker container prune -f
	@docker image prune -f

prune:
	@docker system prune -af