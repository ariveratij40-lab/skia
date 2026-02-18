#!/bin/bash
# SKIA Database Migration Script
# Script de migraciones usando golang-migrate

set -e

# Colores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Configuración
MIGRATIONS_DIR="../migrations"

# Leer variables de entorno
if [ -f ../.env ]; then
    export $(cat ../.env | grep -v '^#' | xargs)
fi

# Construir DATABASE_URL
DB_HOST=${DB_HOST:-global_postgres_db}
DB_PORT=${DB_PORT:-5432}
DB_NAME=${DB_NAME:-skia_db}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-}

if [ -z "$DB_PASSWORD" ]; then
    echo -e "${RED}Error: DB_PASSWORD no configurado${NC}"
    exit 1
fi

DATABASE_URL="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# Funciones
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Verificar golang-migrate
if ! command -v migrate &> /dev/null; then
    log_info "Instalando golang-migrate..."
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
    sudo mv migrate /usr/local/bin/
fi

# Comandos
case "${1:-up}" in
    up)
        log_info "Aplicando migraciones..."
        migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" up
        log_info "Migraciones aplicadas exitosamente"
        ;;
    
    down)
        log_warn "Revirtiendo última migración..."
        migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" down 1
        log_info "Migración revertida"
        ;;
    
    version)
        log_info "Versión actual:"
        migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" version
        ;;
    
    force)
        if [ -z "$2" ]; then
            log_error "Uso: $0 force <version>"
            exit 1
        fi
        log_warn "Forzando versión a $2..."
        migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" force "$2"
        ;;
    
    create)
        if [ -z "$2" ]; then
            log_error "Uso: $0 create <nombre_migracion>"
            exit 1
        fi
        log_info "Creando migración: $2"
        migrate create -ext sql -dir "$MIGRATIONS_DIR" -seq "$2"
        log_info "Archivos creados en $MIGRATIONS_DIR"
        ;;
    
    status)
        log_info "Estado de migraciones:"
        migrate -path "$MIGRATIONS_DIR" -database "$DATABASE_URL" version 2>&1 || true
        echo ""
        echo "Archivos de migración disponibles:"
        ls -la "$MIGRATIONS_DIR"
        ;;
    
    *)
        echo "Uso: $0 {up|down|version|force|create|status}"
        echo ""
        echo "Comandos:"
        echo "  up              - Aplicar todas las migraciones pendientes"
        echo "  down            - Revertir última migración"
        echo "  version         - Mostrar versión actual"
        echo "  force <version> - Forzar versión específica"
        echo "  create <nombre> - Crear nueva migración"
        echo "  status          - Mostrar estado de migraciones"
        exit 1
        ;;
esac
