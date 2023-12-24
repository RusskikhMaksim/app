.PHONY: dev-up
dev-up:
	docker compose -d --file=./deployments/development/docker-compose.yaml up

.PHONY: dev-down
dev-down:
	docker compose --file=./deployments/development/docker-compose.yaml down

.PHONY: dev
dev:
	docker compose --file=./deployments/development/docker-compose.yaml up

.PHONY: dev-build
dev-build:
	docker compose --file=./deployments/development/docker-compose.yaml build


.PHONY: dev-env

define ENV_SAMPLE
$(shell sed 's/=.*//' ./deployments/development/.env.example)
endef

export ENV_SAMPLE
env:
	echo "$$ENV_SAMPLE" > ./deployments/development/.env;\
