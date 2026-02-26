# SKIA - Sistema de Gestión de Infraestructura Física

[![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/ariveratij40-lab/skia)
[![Flutter](https://img.shields.io/badge/Flutter-3.16+-02569B.svg)](https://flutter.dev)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-336791.svg)](https://postgresql.org)
[![License](https://img.shields.io/badge/license-Proprietary-red.svg)]()

> **Solución integral para gestión de cableado estructurado con escaneo RFID en dispositivos Zebra MC22U**

![SKIA Banner](docs/assets/skia-banner.png)

## 📋 Tabla de Contenidos

- [Descripción](#descripción)
- [Arquitectura](#arquitectura)
- [Stack Tecnológico](#stack-tecnológico)
- [Requisitos Previos](#requisitos-previos)
- [Instalación](#instalación)
- [Despliegue](#despliegue)
- [Migraciones de Base de Datos](#migraciones-de-base-de-datos)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Documentación](#documentación)
- [Soporte](#soporte)

## 🎯 Descripción

SKIA es un sistema empresarial de gestión de infraestructura física que permite:

- **Escaneo RFID**: Captura de nodos de cableado mediante dispositivos Zebra MC22U
- **Gestión de Inventario**: Control completo de racks, patch panels, cables y nodos
- **Trazabilidad**: Seguimiento del historial completo de cada componente
- **Dashboard en Tiempo Real**: Videowall con planos interactivos y métricas
- **Multi-tenancy**: Soporte para múltiples organizaciones con aislamiento de datos

### Características Principales

| Módulo | Descripción |
|--------|-------------|
| 📱 **Mobile** | App Flutter para escaneo RFID y gestión de campo |
| 🔧 **Backend** | API REST en Go con autenticación JWT y RLS |
| 🌐 **Web** | Dashboard Next.js con planos interactivos |
| 🗄️ **Database** | PostgreSQL con Row Level Security |

## 🏗️ Arquitectura

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              SKIA PLATFORM                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐  │
│  │   Mobile    │    │    Web      │    │   Videowall │    │   Admin     │  │
│  │  (Flutter)  │    │  (Next.js)  │    │  (Next.js)  │    │  (Next.js)  │  │
│  │  Zebra MC22U│    │   Dashboard │    │   Planos    │    │   Panel     │  │
│  └──────┬──────┘    └──────┬──────┘    └──────┬──────┘    └──────┬──────┘  │
│         │                  │                  │                  │         │
│         └──────────────────┴──────────────────┴──────────────────┘         │
│                                    │                                         │
│                         ┌──────────▼──────────┐                             │
│                         │   API Gateway       │                             │
│                         │   (Go + Gin/Echo)   │                             │
│                         └──────────┬──────────┘                             │
│                                    │                                         │
│                         ┌──────────▼──────────┐                             │
│                         │   Backend API       │                             │
│                         │   (Go + JWT + RLS)  │                             │
│                         └──────────┬──────────┘                             │
│                                    │                                         │
│         ┌──────────────────────────┼──────────────────────────┐             │
│         │                          │                          │             │
│  ┌──────▼──────┐          ┌────────▼────────┐        ┌───────▼──────┐      │
│  │  PostgreSQL │          │   Redis Cache   │        │   MinIO      │      │
│  │   (RLS)     │          │   (Sesiones)    │        │  (Archivos)  │      │
│  └─────────────┘          └─────────────────┘        └──────────────┘      │
│                                                                            │
└────────────────────────────────────────────────────────────────────────────┘
```

> 📊 Ver [Diagrama Detallado](docs/architecture/diagrama-sistema.md)

## 🛠️ Stack Tecnológico

### Backend
- **Go 1.21+** - Lenguaje principal
- **Gin/Echo** - Framework HTTP
- **pgx** - Driver PostgreSQL
- **golang-migrate** - Migraciones de BD
- **JWT** - Autenticación

### Frontend Web
- **Next.js 14** - Framework React
- **TypeScript** - Tipado estático
- **Tailwind CSS** - Estilos
- **shadcn/ui** - Componentes UI
- **React Query** - Gestión de estado

### Mobile
- **Flutter 3.16+** - Framework multiplataforma
- **Dart** - Lenguaje
- **DataWedge API** - Integración Zebra
- **HTTP Client** - Comunicación API

### Infraestructura
- **Docker** - Contenerización
- **PostgreSQL 15+** - Base de datos
- **Redis** - Cache y sesiones
- **Nginx** - Proxy inverso

## 📋 Requisitos Previos

### VPS (Producción)
- Docker 24.0+
- Docker Compose 2.20+
- Red Docker: `infra_network` (existente)
- PostgreSQL: `global_postgres_db` (existente)

### Desarrollo
- Go 1.21+
- Node.js 18+
- Flutter 3.16+
- PostgreSQL 15+

## 🚀 Instalación

### 1. Clonar Repositorio

```bash
git clone https://github.com/ariveratij40-lab/skia.git
cd skia
```

### 2. Configurar Variables de Entorno

```bash
cp .env.example .env
# Editar .env con tus configuraciones
```

### 3. Configurar Base de Datos

```bash
# Crear base de datos (ejecutar en PostgreSQL existente)
CREATE DATABASE skia_db;

# Ejecutar migraciones
cd infra
./migrate.sh up
```

### 4. Iniciar Servicios

```bash
# Backend y Web
docker-compose -f infra/docker-compose.yml up -d

# Verificar estado
docker-compose -f infra/docker-compose.yml ps
```

## 📦 Despliegue

### Despliegue en VPS (Producción)

```bash
# 1. Conectar al VPS
ssh usuario@vps

# 2. Clonar/Actualizar repositorio
cd /opt/skia
git pull origin main

# 3. Ejecutar script de despliegue
cd infra
./deploy.sh

# 4. Verificar logs
docker-compose logs -f skia-backend
docker-compose logs -f skia-web
```

### Despliegue Manual

```bash
# Construir imágenes
docker-compose -f infra/docker-compose.yml build

# Iniciar servicios
docker-compose -f infra/docker-compose.yml up -d

# Verificar salud
curl http://localhost:8080/health
```

## 🗄️ Migraciones de Base de Datos

### Comandos Disponibles

```bash
cd infra

# Aplicar todas las migraciones pendientes
./migrate.sh up

# Revertir última migración
./migrate.sh down

# Crear nueva migración
./migrate.sh create nombre_de_la_migracion

# Verificar versión actual
./migrate.sh version

# Forzar versión específica
./migrate.sh force 1
```

### Estructura de Migraciones

```
migrations/
├── 000001_create_schema.up.sql      # Crear tablas
├── 000001_create_schema.down.sql    # Eliminar tablas
├── 000002_seed_tenants.up.sql       # Datos iniciales
├── 000002_seed_tenants.down.sql     # Eliminar datos
└── ...
```

> ⚠️ **IMPORTANTE**: Nunca modificar migraciones ya aplicadas. Crear nueva migración para cambios.

## 📁 Estructura del Proyecto

```
skia/
├── 📁 infra/                    # Infraestructura Docker
│   ├── docker-compose.yml       # Orquestación de servicios
│   ├── deploy.sh               # Script de despliegue
│   └── migrate.sh              # Script de migraciones
│
├── 📁 backend/                  # API Backend (Go)
│   ├── cmd/                    # Entry points
│   ├── internal/               # Código interno
│   │   ├── config/            # Configuración
│   │   ├── handlers/          # HTTP handlers
│   │   ├── middleware/        # Middleware (auth, RLS)
│   │   ├── models/            # Modelos de datos
│   │   ├── repository/        # Acceso a BD
│   │   └── services/          # Lógica de negocio
│   ├── pkg/                    # Librerías compartidas
│   ├── migrations/             # Migraciones (referencia)
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
│
├── 📁 mobile/                   # App Flutter
│   ├── lib/
│   │   ├── main.dart
│   │   ├── config/            # Configuración
│   │   ├── models/            # Modelos
│   │   ├── screens/           # Pantallas
│   │   ├── services/          # API client
│   │   ├── widgets/           # Componentes
│   │   └── utils/             # Utilidades
│   ├── android/               # Config Android
│   ├── ios/                   # Config iOS
│   ├── pubspec.yaml
│   └── Dockerfile
│
├── 📁 web/                      # Dashboard Next.js
│   ├── app/                    # App Router
│   ├── components/             # Componentes React
│   ├── lib/                    # Utilidades
│   ├── hooks/                  # Custom hooks
│   ├── types/                  # TypeScript types
│   ├── public/                 # Assets estáticos
│   ├── Dockerfile
│   ├── package.json
│   └── next.config.js
│
├── 📁 migrations/               # Migraciones SQL
│   ├── 000001_create_schema.up.sql
│   ├── 000001_create_schema.down.sql
│   ├── 000002_seed_tenants.up.sql
│   └── 000002_seed_tenants.down.sql
│
├── 📁 docs/                     # Documentación
│   ├── architecture/           # Arquitectura
│   ├── api/                    # Documentación API
│   ├── deployment/             # Guías de despliegue
│   └── assets/                 # Imágenes y recursos
│
├── .env.example                 # Variables de entorno template
├── .gitignore                   # Archivos ignorados
└── README.md                    # Este archivo
```

## 📚 Documentación

| Documento | Descripción |
|-----------|-------------|
| [Arquitectura del Sistema](docs/architecture/diagrama-sistema.md) | Diagramas y flujos |
| [API Reference](docs/api/README.md) | Documentación endpoints |
| [Guía de Despliegue](docs/deployment/README.md) | Instrucciones detalladas |
| [Mobile Development](docs/mobile/README.md) | Guía app Flutter |
| [Database Schema](docs/database/schema.md) | Esquema de base de datos |

## 🔐 Seguridad

- **Row Level Security (RLS)**: Aislamiento de datos por tenant
- **JWT**: Tokens con expiración configurable
- **HTTPS**: Comunicación encriptada
- **Rate Limiting**: Protección contra abuso
- **Input Validation**: Validación en todos los endpoints

## 🛟 Soporte

### Canales de Comunicación

- **Issues**: [GitHub Issues](https://github.com/ariveratij40-lab/skia/issues)
- **Email**: soporte@skia-lab.com

### Reportar Problemas

1. Verificar [issues existentes](https://github.com/ariveratij40-lab/skia/issues)
2. Crear nuevo issue con:
   - Descripción del problema
   - Pasos para reproducir
   - Logs relevantes
   - Entorno (OS, versión, etc.)

## 📄 Licencia

Este proyecto es propietario y confidencial. Todos los derechos reservados.

---

<p align="center">
  <strong>SKIA</strong> - Gestión Inteligente de Infraestructura<br>
  <sub>Desarrollado con ❤️ por el equipo SKIA</sub>
</p>
# trigger
# trigger-2
# redeploy
