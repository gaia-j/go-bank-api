@echo off

chcp 65001 > nul

IF "%1"=="" (
    echo Por favor, use o parametro up ou down para realizar a migrate desejada!
    exit /b 1
)

set MIGRATE_TYPE=%1

set POSTGRES_URL=postgres://user:password@localhost:5432/db?sslmode=disable

migrate -database %POSTGRES_URL% -path scripts/db/migrations %MIGRATE_TYPE%
