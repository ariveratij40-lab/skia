# Guía de Despliegue - SKIA

## Requisitos del VPS

### Hardware Mínimo
- **CPU**: 2 cores
- **RAM**: 4 GB
- **Disco**: 50 GB SSD
- **Red**: Conexión estable, puertos 80/443 abiertos

### Software Requerido
- Docker 24.0+
- Docker Compose 2.20+
- Git 2.30+
- PostgreSQL 15+ (existente)

### Infraestructura Existente (NO MODIFICAR)

Según la Directiva Técnica Oficial, el VPS ya cuenta con:

| Componente | Nombre | Estado |
|------------|--------|--------|
| PostgreSQL | `global_postgres_db` | ✅ Existente |
| Red Docker | `infra_network` | ✅ Existente |
| Nginx | `nginx` | ✅ Existente |
| pgAdmin | Configurado | ✅ Existente |

> ⚠️ **ADVERTENCIA**: Estos componentes son inmutables. NO crear, modificar ni eliminar.

---

## Configuración Inicial

### 1. Preparar Base de Datos

Conectarse a PostgreSQL existente:

```bash
# Ejecutar en el contenedor PostgreSQL
docker exec -it global_postgres_db psql -U postgres

# Crear base de datos para SKIA
CREATE DATABASE skia_db;

# Verificar creación
\l

# Salir
\q
```

### 2. Configurar Variables de Entorno

```bash
cd /opt/skia
cp .env.example .env
nano .env
```

**Configuración mínima:**

```bash
# Database (usar contenedor existente)
DB_HOST=global_postgres_db
DB_PORT=5432
DB_NAME=skia_db
DB_USER=postgres
DB_PASSWORD=tu_password_seguro
DATABASE_URL=postgresql://postgres:tu_password_seguro@global_postgres_db:5432/skia_db?sslmode=disable

# JWT
JWT_SECRET=genera_un_secreto_largo_y_aleatorio_32_chars
JWT_EXPIRATION=15m
JWT_REFRESH_EXPIRATION=7d

# App
APP_ENV=production
APP_PORT=8080
LOG_LEVEL=info

# Domain (para configuración nginx)
DOMAIN=tu-dominio.com
```

### 3. Ejecutar Migraciones

```bash
cd /opt/skia/infra

# Instalar golang-migrate (primera vez)
if ! command -v migrate &> /dev/null; then
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/
fi

# Aplicar migraciones
./migrate.sh up

# Verificar
./migrate.sh version
```

---

## Despliegue con Script Automático

### Script Principal

```bash
cd /opt/skia/infra
./deploy.sh
```

Este script:
1. Construye las imágenes Docker
2. Inicia los contenedores
3. Verifica la salud de los servicios

### Verificar Despliegue

```bash
# Estado de contenedores
docker-compose ps

# Logs
docker-compose logs -f skia-backend
docker-compose logs -f skia-web

# Health check
curl http://localhost:8080/health
```

---

## Despliegue Manual

### Paso 1: Construir Imágenes

```bash
cd /opt/skia

# Backend
docker build -t skia-backend:latest ./backend

# Web
docker build -t skia-web:latest ./web
```

### Paso 2: Iniciar Servicios

```bash
cd /opt/skia/infra
docker-compose up -d
```

### Paso 3: Verificar

```bash
# Contenedores en ejecución
docker ps | grep skia

# Logs
docker logs skia-backend
docker logs skia-web

# Test endpoint
curl -s http://localhost:8080/health | jq
```

---

## Configuración Nginx (Reverse Proxy)

Agregar al nginx existente:

```nginx
# /etc/nginx/conf.d/skia.conf

upstream skia_backend {
    server localhost:8080;
}

upstream skia_web {
    server localhost:3000;
}

# API Server
server {
    listen 443 ssl http2;
    server_name api.tu-dominio.com;

    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;

    location / {
        proxy_pass http://skia_backend;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
        
        # Timeouts
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }
}

# Web Dashboard
server {
    listen 443 ssl http2;
    server_name app.tu-dominio.com;

    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;

    location / {
        proxy_pass http://skia_web;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }
}

# Redirección HTTP a HTTPS
server {
    listen 80;
    server_name api.tu-dominio.com app.tu-dominio.com;
    return 301 https://$server_name$request_uri;
}
```

**Recargar nginx:**

```bash
sudo nginx -t
sudo systemctl reload nginx
```

---

## SSL/TLS con Let's Encrypt

```bash
# Instalar certbot
sudo apt install certbot python3-certbot-nginx

# Obtener certificados
sudo certbot --nginx -d api.tu-dominio.com -d app.tu-dominio.com

# Auto-renewal
sudo systemctl enable certbot.timer
```

---

## Actualizaciones

### Proceso de Update

```bash
cd /opt/skia

# 1. Backup
cd infra
./backup.sh

# 2. Pull cambios
git pull origin main

# 3. Ejecutar migraciones
./migrate.sh up

# 4. Re-desplegar
./deploy.sh

# 5. Verificar
curl -s http://localhost:8080/health
```

### Rollback

```bash
cd /opt/skia/infra

# Revertir migración
./migrate.sh down

# Restaurar backup (si es necesario)
# ...

# Re-desplegar versión anterior
git checkout <version-anterior>
./deploy.sh
```

---

## Monitoreo

### Health Checks

```bash
# Backend
curl -f http://localhost:8080/health || echo "Backend DOWN"

# Web
curl -f http://localhost:3000 || echo "Web DOWN"
```

### Logs

```bash
# Ver logs en tiempo real
docker-compose logs -f

# Logs específicos
docker-compose logs -f skia-backend

# Logs históricos
docker-compose logs --tail=100 skia-backend
```

### Recursos

```bash
# Uso de recursos
docker stats

# Espacio en disco
df -h

# Memoria
free -h
```

---

## Troubleshooting

### Problema: Contenedor no inicia

```bash
# Ver logs
docker-compose logs skia-backend

# Verificar variables de entorno
docker-compose config

# Reconstruir
docker-compose down
docker-compose up --build -d
```

### Problema: Error de conexión a BD

```bash
# Verificar red
docker network inspect infra_network

# Probar conexión
docker exec -it skia-backend ping global_postgres_db

# Verificar credenciales
docker exec -it global_postgres_db psql -U postgres -d skia_db -c "\dt"
```

### Problema: Migraciones fallan

```bash
# Verificar versión
./migrate.sh version

# Forzar versión (con cuidado)
./migrate.sh force 1

# Revertir y re-aplicar
./migrate.sh down
./migrate.sh up
```

### Problema: Permisos denegados

```bash
# Verificar permisos de archivos
ls -la /opt/skia/

# Corregir
sudo chown -R $USER:$USER /opt/skia
```

---

## Backup y Recuperación

### Backup Automático

```bash
# Agregar a crontab
crontab -e

# Backup diario a las 2 AM
0 2 * * * /opt/skia/infra/backup.sh >> /var/log/skia-backup.log 2>&1
```

### Backup Manual

```bash
cd /opt/skia/infra
./backup.sh
```

### Restaurar Backup

```bash
# Detener servicios
docker-compose down

# Restaurar base de datos
docker exec -i global_postgres_db psql -U postgres -d skia_db < backup_YYYYMMDD_HHMMSS.sql

# Iniciar servicios
docker-compose up -d
```

---

## Checklist de Despliegue

- [ ] Variables de entorno configuradas
- [ ] Base de datos creada
- [ ] Migraciones aplicadas
- [ ] Imágenes Docker construidas
- [ ] Contenedores en ejecución
- [ ] Health check exitoso
- [ ] Nginx configurado
- [ ] SSL configurado
- [ ] DNS apuntando al VPS
- [ ] Backup configurado
- [ ] Monitoreo activo

---

**Versión**: 1.0.0  
**Última actualización**: 2024
