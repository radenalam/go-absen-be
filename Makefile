include .env

.PHONY: up down migrate seed seeder

# Ambil DATABASE_URL dari environment
DB_URL   := "$(DATABASE_URL)"
MIG_DIR  := db/migrations
SEED_DIR := db/seeder

# Ambil argumen ke-2 di MAKECMDGOALS (bisa dipakai untuk migrate maupun seeder)
NAME := $(word 2, $(MAKECMDGOALS))

# ----------------------------
# Target migrasi
# ----------------------------
up:
	@echo ">> Migrating up (migrations)..."
	migrate -path $(MIG_DIR) -database "$(DB_URL)" up

down:
	@echo ">> Migrating down (migrations)..."
	migrate -path $(MIG_DIR) -database "$(DB_URL)" down

migrate:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make migrate <migration_name>"; exit 1; \
	fi
	@echo ">> Creating new migration: $(NAME)"
	migrate create -ext sql -dir $(MIG_DIR) $(NAME)

force:
	@if [ "$(words $(MAKECMDGOALS))" -lt 2 ]; then \
		echo "Usage: make force <version>"; exit 1; \
	fi
	@VERS=$$(lastword $(MAKECMDGOALS)); \
	echo ">> Forcing migration version to $$VERS"; \
	migrate -path $(MIG_DIR) -database "$(DB_URL)" force $$VERS

drop:
	@echo ">> Dropping all tables (fresh start)..."; \
	migrate -path $(MIG_DIR) -database "$(DB_URL)" drop	

# ----------------------------
# Target seeder
# ----------------------------

# make seed
#   → Menjalankan semua file *.up.sql di folder db/seeder
seed:
	@echo ">> Running all seeders (*.up.sql) in $(SEED_DIR)..."
	@if [ ! -d "$(SEED_DIR)" ]; then \
		echo "Folder '$(SEED_DIR)' tidak ditemukan."; exit 1; \
	fi
	@for f in $(SEED_DIR)/*.up.sql; do \
		if [ -f $$f ]; then \
			echo "  → exec $$f"; \
			psql "$(DB_URL)" -f $$f; \
		fi \
	done

# make seeder <name>
#   → Membuat satu file seeder <timestamp>_<name>.up.sql di folder db/seeder
seeder:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make seeder <name>"; exit 1; \
	fi
	@FILENAME=$(SEED_DIR)/$$(date +%Y%m%d%H%M%S)_$(NAME).up.sql; \
	echo "-- Seeder: $(NAME) (created at $$(date +%Y%m%d%H%M%S))" > $$FILENAME; \
	echo ">> Created seeder file: $$FILENAME"

# Rule wildcard supaya Make tidak error kalau ada goal kedua (misalnya “admin”)
# dan agar Make tidak mencari file/target “admin”
%::
	@:
