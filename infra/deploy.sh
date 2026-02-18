#!/bin/bash
# SKIA Deployment Script
# Script de despliegue para producción

set -e

# Colores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuración
COMPOSE_FILE="docker-compose.yml"
PROJECT_NAME="skia"

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

# Verificar Docker
if ! command -v docker &> /dev/null; then
    log_error "Docker no está instalado"
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    log_error "Docker Compose no está instalado"
    exit 1
fi

# Verificar red externa
if ! docker network inspect infra_network &> /dev/null; then
    log_error "La red 'infra_network' no existe. Crear antes de desplegar."
    exit 1
fi

log_info "Iniciando despliegue de SKIA..."

# Pull de imágenes base
docker-compose -f "$COMPOSE_FILE" pull

# Construir imágenes
log_info "Construyendo imágenes..."
docker-compose -f "$COMPOSE_FILE" build --no-cache

# Detener servicios existentes
log_info "Deteniendo servicios existentes..."
docker-compose -f "$COMPOSE_FILE" down

# Iniciar servicios
log_info "Iniciando servicios..."
docker-compose -f "$COMPOSE_FILE" up -d

# Esperar a que los servicios estén listos
log_info "Esperando a que los servicios estén listos..."
sleep 10

# Verificar salud
log_info "Verificando salud de los servicios..."

# Backend
if curl -s http://localhost:8080/health > /dev/null; then
    log_info "✓ Backend está saludable"
else
    log_error "✗ Backend no responde"
    docker-compose -f "$COMPOSE_FILE" logs skia-backend --tail=50
    exit 1
fi

# Web
if curl -s http://localhost:3000 > /dev/null; then
    log_info "✓ Web está saludable"
else
    log_error "✗ Web no responde"
    docker-compose -f "$COMPOSE_FILE" logs skia-web --tail=50
    exit 1
fi

log_info "Despliegue completado exitosamente!"
log_info "Backend: http://localhost:8080"
log_info "Web: http://localhost:3000"
