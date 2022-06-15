migrate_up:
ifdef dbUrl
	migrate --path migrations --database $(dbUrl) --verbose up
else
	migrate --path migrations --database "postgres://postgres:postgres@localhost:5432/link_management?sslmode=disable" --verbose up
endif

migrate_down:
ifdef dbUrl
	migrate --path migrations --database $(dbUrl) --verbose down
else
	migrate --path migrations --database "postgres://postgres:postgres@localhost:5432/link_management?sslmode=disable" --verbose down
endif

create_migration:
ifdef name
	migrate create -ext sql -dir migrations -seq $(name)
else
	@echo "Must provide migration name"
endif

.PHONY: migrate_up migrate_down create_migration