# SKIA Backend API

**SKIA** - Plataforma Industrial de Gestión de Infraestructura (DCIM)

Versión: 1.0.0 | Fases: 0 (Multi-tenancy & Billing) + 1 (DCIM Inventory)

## 📋 Descripción

Backend API completo para la plataforma SKIA, construido con Go y Fiber. Proporciona endpoints para:

- **Fase 0:** Autenticación, gestión de usuarios, multi-tenancy, facturación y auditoría
- **Fase 1:** Gestión de infraestructura (sitios, edificios, pisos, salas, racks, dispositivos)

## 🛠️ Tecnologías

- **Framework:** Fiber (Go web framework)
- **Database:** PostgreSQL con GORM
- **Authentication:** JWT (JSON Web Tokens)
- **Containerization:** Docker
- **Dependencies:** go.mod / go.sum

## 📁 Estructura del Proyecto

```
backend/
├── cmd/
│   └── api/
│       └── main.go              # Punto de entrada
├── internal/
│   ├── config/
│   │   └── database.go          # Configuración de BD
│   ├── handlers/
│   │   ├── auth_handler.go      # Endpoints de autenticación
│   │   ├── user_handler.go      # Endpoints de usuarios
│   │   └── dcim_handler.go      # Endpoints de DCIM
│   ├── middleware/
│   │   └── auth.go              # JWT middleware
│   ├── models/
│   │   └── models.go            # Modelos GORM
│   └── services/
│       ├── auth_service.go      # Lógica de autenticación
│       └── dcim_service.go      # Lógica de DCIM
├── Dockerfile                   # Configuración Docker
├── go.mod                       # Dependencias
├── go.sum                       # Checksums
└── README.md                    # Este archivo
```

## 🚀 Inicio Rápido

### Requisitos Previos

- Go 1.21+
- PostgreSQL 12+
- Docker (opcional)

### Instalación Local

1. **Clonar el repositorio**
   ```bash
   git clone https://github.com/ariveratij40-lab/skia.git
   cd skia/backend
   ```

2. **Instalar dependencias**
   ```bash
   go mod download
   ```

3. **Configurar variables de entorno**
   ```bash
   cp .env.example .env
   ```

   **Variables necesarias:**
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=skia_db
   DB_SSL_MODE=disable
   JWT_SECRET=your_jwt_secret_min_32_chars
   PORT=8080
   CORS_ORIGINS=http://localhost:3000
   ```

4. **Ejecutar el servidor**
   ```bash
   go run ./cmd/api/main.go
   ```

   El servidor estará disponible en `http://localhost:8080`

### Despliegue con Docker

1. **Construir la imagen**
   ```bash
   docker build -t skia-api:latest .
   ```

2. **Ejecutar el contenedor**
   ```bash
   docker run -p 8080:8080 \
     -e DB_HOST=postgres \
     -e DB_PORT=5432 \
     -e DB_USER=postgres \
     -e DB_PASSWORD=password \
     -e DB_NAME=skia_db \
     -e JWT_SECRET=your_secret \
     skia-api:latest
   ```

## 📚 API Endpoints

### Autenticación (Públicos)

```
POST   /api/v1/auth/register       # Crear cuenta
POST   /api/v1/auth/login          # Iniciar sesión
POST   /api/v1/auth/refresh        # Renovar token
```

### Usuarios (Protegidos)

```
GET    /api/v1/users               # Listar usuarios
GET    /api/v1/users/:id           # Obtener usuario
PUT    /api/v1/users/:id           # Actualizar usuario
DELETE /api/v1/users/:id           # Eliminar usuario
```

### Sitios (Protegidos)

```
GET    /api/v1/sites               # Listar sitios
POST   /api/v1/sites               # Crear sitio
GET    /api/v1/sites/:id           # Obtener sitio
PUT    /api/v1/sites/:id           # Actualizar sitio
DELETE /api/v1/sites/:id           # Eliminar sitio
```

### Edificios (Protegidos)

```
GET    /api/v1/buildings           # Listar edificios
POST   /api/v1/buildings           # Crear edificio
GET    /api/v1/buildings/:id       # Obtener edificio
PUT    /api/v1/buildings/:id       # Actualizar edificio
DELETE /api/v1/buildings/:id       # Eliminar edificio
```

### Pisos (Protegidos)

```
GET    /api/v1/floors              # Listar pisos
POST   /api/v1/floors              # Crear piso
GET    /api/v1/floors/:id          # Obtener piso
PUT    /api/v1/floors/:id          # Actualizar piso
DELETE /api/v1/floors/:id          # Eliminar piso
```

### Salas (Protegidos)

```
GET    /api/v1/rooms               # Listar salas
POST   /api/v1/rooms               # Crear sala
GET    /api/v1/rooms/:id           # Obtener sala
PUT    /api/v1/rooms/:id           # Actualizar sala
DELETE /api/v1/rooms/:id           # Eliminar sala
```

### Racks (Protegidos)

```
GET    /api/v1/racks               # Listar racks
POST   /api/v1/racks               # Crear rack
GET    /api/v1/racks/:id           # Obtener rack
PUT    /api/v1/racks/:id           # Actualizar rack
DELETE /api/v1/racks/:id           # Eliminar rack
```

### Dispositivos (Protegidos)

```
GET    /api/v1/devices             # Listar dispositivos
POST   /api/v1/devices             # Crear dispositivo
GET    /api/v1/devices/:id         # Obtener dispositivo
PUT    /api/v1/devices/:id         # Actualizar dispositivo
DELETE /api/v1/devices/:id         # Eliminar dispositivo
```

### Dashboard (Protegido)

```
GET    /api/v1/dashboard           # Obtener estadísticas
```

### Health Check

```
GET    /health                     # Verificar estado del servidor
```

## 🔐 Autenticación

Todos los endpoints protegidos requieren un token JWT en el header:

```bash
Authorization: Bearer <token>
```

### Ejemplo de Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

Respuesta:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "name": "User Name",
    "tenant_id": "uuid"
  }
}
```

## 🗄️ Modelos de Datos

### Fase 0: Multi-Tenancy

- **Tenant:** Organización cliente
- **User:** Usuario del sistema
- **Role:** Rol de usuario (admin, user, viewer)
- **Permission:** Permisos del sistema
- **AuditLog:** Registro de cambios
- **Subscription:** Información de suscripción

### Fase 1: DCIM

- **Site:** Ubicación física (data center, oficina)
- **Building:** Edificio dentro de un sitio
- **Floor:** Piso dentro de un edificio
- **Room:** Sala (sala de servidores)
- **Rack:** Rack de servidores
- **Device:** Dispositivo físico (servidor, switch, etc.)

## 📊 Características de Seguridad

- ✅ Autenticación JWT
- ✅ Multi-tenancy con aislamiento de datos
- ✅ CORS configurable
- ✅ Validación de entrada
- ✅ Auditoría de cambios
- ✅ Encriptación de contraseñas (bcrypt)
- ✅ Rate limiting (configurable)

## 🧪 Testing

```bash
# Ejecutar tests
go test ./...

# Con cobertura
go test -cover ./...
```

## 📝 Logging

El servidor registra:
- Solicitudes HTTP
- Errores
- Cambios en la base de datos
- Acciones de usuarios

## 🚀 Despliegue en Producción

### Requisitos

- PostgreSQL 12+ (centralizado)
- Nginx (reverse proxy)
- Let's Encrypt (SSL/TLS)
- Docker & Docker Compose

### Pasos

1. Configurar variables de entorno con valores reales
2. Construir imagen Docker
3. Desplegar con docker-compose
4. Configurar Nginx como reverse proxy
5. Crear certificados SSL
6. Ejecutar migraciones

Ver `VPS_DEPLOYMENT_GUIDE.md` para instrucciones detalladas.

## 🐛 Troubleshooting

### Error: "connection refused"
- Verificar que PostgreSQL está corriendo
- Verificar credenciales en .env

### Error: "invalid token"
- Verificar que JWT_SECRET es el mismo en cliente y servidor
- Verificar que el token no ha expirado

### Error: "permission denied"
- Verificar que el usuario tiene los permisos necesarios
- Verificar rol del usuario

## 📞 Soporte

Para reportar bugs o solicitar features, abrir un issue en GitHub.

## 📄 Licencia

MIT License - Ver LICENSE file

---

**Versión:** 1.0.0  
**Última actualización:** Marzo 2026  
**Autor:** SKIA Development Team
