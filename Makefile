CHECK_ENV = test -e .env || (cp .env.example .env  && echo "Você deve preencher seu .env" && exit 1)
ARTISAN_CMD = go run . artisan
GENERATE_KEY = $(ARTISAN_CMD) key:generate
GENERATE_JWT = $(ARTISAN_CMD) jwt:secret

build:
	@${CHECK_ENV}
	@${GENERATE_KEY}
	@${GENERATE_JWT}
	docker compose build

run:
	@${CHECK_ENV}
	@echo "Running => Starting application"
	docker compose up -d
	docker logs -f zooapi01 -n 30
	
stop:
	docker compose down -v
