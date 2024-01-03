migrate_option= -path "./migrations" -database "postgresql://postgres:postgres@localhost:5432/nebula_corp?sslmode=disable"

.PHONY: migrate/up
migrate/up:
	@echo "appliying migrations..."
	@migrate ${migrate_option} up
	@echo "done!"

.PHONY: migrate/down
migrate/down:
	@echo "undoing migrations"
	@migrate ${migrate_option} down
	@echo "done!"

.PHONY: migrate/force
migrate/force:
	@echo "forcing migrations"
	@migrate ${migrate_option} force $(version)
	@echo "done"