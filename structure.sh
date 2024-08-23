#!/bin/bash

# Корневая директория проекта
PROJECT_ROOT="."

# Создание структуры директорий
mkdir -p $PROJECT_ROOT/{cmd/notesapp,internal/{app,config,service/notes,storage/postgres,logger},pkg/{utils,errors},configs,migrations,scripts,deployments,docs/api}

# Создание файлов
touch $PROJECT_ROOT/cmd/notesapp/main.go
touch $PROJECT_ROOT/internal/app/{app.go,routes.go,middleware.go}
touch $PROJECT_ROOT/internal/config/config.go
touch $PROJECT_ROOT/internal/service/notes/{service.go,validation.go,auth.go}
touch $PROJECT_ROOT/internal/storage/{models.go,postgres/repo.go}
touch $PROJECT_ROOT/internal/logger/logger.go
touch $PROJECT_ROOT/pkg/utils/utils.go
touch $PROJECT_ROOT/pkg/errors/errors.go
touch $PROJECT_ROOT/configs/{dev.yaml,prod.yaml,test.yaml}
touch $PROJECT_ROOT/migrations/{001_create_notes_table.up.sql,001_create_notes_table.down.sql}
touch $PROJECT_ROOT/scripts/{build.sh,test.sh,deploy.sh}
touch $PROJECT_ROOT/deployments/docker-compose.yaml
touch $PROJECT_ROOT/docs/architecture.md
touch $PROJECT_ROOT/docs/api/notes-api.md
touch $PROJECT_ROOT/Dockerfile
touch $PROJECT_ROOT/go.mod
touch $PROJECT_ROOT/go.sum
touch $PROJECT_ROOT/README.md

echo "Проектная структура успешно создана в директории $PROJECT_ROOT"
