CMD_ARGS=$(filter-out $@,$(MAKECMDGOALS))

start:
	docker-compose up -d ${CMD_ARGS}

stop:
	docker-compose stop ${CMD_ARGS}
	docker-compose rm -v -f

restart:
	$(MAKE) stop ${CMD_ARGS}
	$(MAKE) start ${CMD_ARGS}

shell:
	docker-compose exec ${CMD_ARGS} sh

build:
	docker-compose build ${CMD_ARGS}

status:
	docker-compose ps

logs:
	docker-compose logs -f --tail 100 ${CMD_ARGS}

%:
	@: