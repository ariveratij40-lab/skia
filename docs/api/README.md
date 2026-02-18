# API Reference - SKIA Backend

## Base URL

```
Producción:  https://api.tu-dominio.com/v1
Desarrollo:  http://localhost:8080/v1
```

## Autenticación

Todas las peticiones (excepto login/register) requieren header:

```
Authorization: Bearer <access_token>
```

## Endpoints

### 🔐 Autenticación

#### POST /auth/login

Autenticar usuario y obtener tokens.

**Request:**
```json
{
  "email": "tecnico@empresa.com",
  "password": "contraseña123"
}
```

**Response (200):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "expires_in": 900,
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "tecnico@empresa.com",
    "role": "technician",
    "tenant_id": "550e8400-e29b-41d4-a716-446655440001"
  }
}
```

#### POST /auth/refresh

Renovar access token.

**Request:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Response (200):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "expires_in": 900
}
```

#### POST /auth/logout

Invalidar tokens.

**Response (204):**
```
No Content
```

---

### 👤 Usuarios

#### GET /users

Listar usuarios del tenant.

**Query Params:**
- `page` (int): Página (default: 1)
- `limit` (int): Items por página (default: 20)
- `role` (string): Filtrar por rol

**Response (200):**
```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "tecnico@empresa.com",
      "role": "technician",
      "first_name": "Juan",
      "last_name": "Pérez",
      "is_active": true,
      "created_at": "2024-01-15T10:30:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 45,
    "total_pages": 3
  }
}
```

#### GET /users/:id

Obtener usuario por ID.

**Response (200):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "tecnico@empresa.com",
  "role": "technician",
  "first_name": "Juan",
  "last_name": "Pérez",
  "phone": "+1234567890",
  "is_active": true,
  "created_at": "2024-01-15T10:30:00Z",
  "last_login": "2024-01-20T08:15:00Z"
}
```

#### POST /users

Crear nuevo usuario (Admin only).

**Request:**
```json
{
  "email": "nuevo@empresa.com",
  "password": "temporal123",
  "role": "technician",
  "first_name": "María",
  "last_name": "García",
  "phone": "+1234567891"
}
```

**Response (201):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440002",
  "email": "nuevo@empresa.com",
  "role": "technician",
  "first_name": "María",
  "last_name": "García",
  "is_active": true,
  "created_at": "2024-01-20T14:30:00Z"
}
```

#### PUT /users/:id

Actualizar usuario.

**Request:**
```json
{
  "first_name": "Juan Carlos",
  "phone": "+1234567899"
}
```

**Response (200):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "tecnico@empresa.com",
  "first_name": "Juan Carlos",
  "phone": "+1234567899",
  "updated_at": "2024-01-20T15:00:00Z"
}
```

#### DELETE /users/:id

Desactivar usuario (soft delete).

**Response (204):**
```
No Content
```

---

### 📦 Racks

#### GET /racks

Listar racks.

**Query Params:**
- `page` (int)
- `limit` (int)
- `location` (string): Filtrar por ubicación
- `search` (string): Búsqueda por nombre

**Response (200):**
```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440010",
      "name": "Rack Principal - Piso 1",
      "location": "Sala de Servidores A",
      "description": "Rack principal de telecomunicaciones",
      "patch_panel_count": 4,
      "node_count": 96,
      "created_at": "2024-01-10T09:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 8
  }
}
```

#### GET /racks/:id

Obtener rack con sus patch panels.

**Response (200):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440010",
  "name": "Rack Principal - Piso 1",
  "location": "Sala de Servidores A",
  "description": "Rack principal de telecomunicaciones",
  "patch_panels": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440011",
      "name": "PP-01",
      "port_count": 24,
      "occupied_ports": 18
    }
  ],
  "created_at": "2024-01-10T09:00:00Z"
}
```

#### POST /racks

Crear rack.

**Request:**
```json
{
  "name": "Rack Secundario - Piso 2",
  "location": "Sala de Servidores B",
  "description": "Rack secundario para expansión"
}
```

**Response (201):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440020",
  "name": "Rack Secundario - Piso 2",
  "location": "Sala de Servidores B",
  "description": "Rack secundario para expansión",
  "created_at": "2024-01-20T16:00:00Z"
}
```

---

### 🔌 Patch Panels

#### GET /patch-panels

Listar patch panels.

**Query Params:**
- `rack_id` (uuid): Filtrar por rack

**Response (200):**
```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440011",
      "rack_id": "550e8400-e29b-41d4-a716-446655440010",
      "name": "PP-01",
      "port_count": 24,
      "occupied_ports": 18,
      "available_ports": 6
    }
  ]
}
```

#### POST /patch-panels

Crear patch panel.

**Request:**
```json
{
  "rack_id": "550e8400-e29b-41d4-a716-446655440010",
  "name": "PP-02",
  "port_count": 24
}
```

---

### 🏷️ Nodos

#### GET /nodes

Listar nodos.

**Query Params:**
- `patch_panel_id` (uuid)
- `status` (string): active, inactive, maintenance
- `search` (string): Búsqueda por RFID o nombre

**Response (200):**
```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440100",
      "patch_panel_id": "550e8400-e29b-41d4-a716-446655440011",
      "rfid_uid": "E200341502001080",
      "port_number": 1,
      "status": "active",
      "description": "Conexión oficina 101",
      "last_scan_at": "2024-01-20T10:30:00Z",
      "created_at": "2024-01-15T09:00:00Z"
    }
  ]
}
```

#### GET /nodes/:id

Obtener nodo con historial.

**Response (200):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440100",
  "patch_panel": {
    "id": "550e8400-e29b-41d4-a716-446655440011",
    "name": "PP-01",
    "rack": {
      "id": "550e8400-e29b-41d4-a716-446655440010",
      "name": "Rack Principal - Piso 1"
    }
  },
  "rfid_uid": "E200341502001080",
  "port_number": 1,
  "status": "active",
  "description": "Conexión oficina 101",
  "cable": {
    "id": "550e8400-e29b-41d4-a716-446655440200",
    "type": "CAT6",
    "length": 15.5,
    "destination": "Oficina 101 - Escritorio 3"
  },
  "scan_history": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440300",
      "scan_type": "check_in",
      "user": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "first_name": "Juan",
        "last_name": "Pérez"
      },
      "latitude": 19.4326,
      "longitude": -99.1332,
      "scanned_at": "2024-01-20T10:30:00Z"
    }
  ],
  "created_at": "2024-01-15T09:00:00Z"
}
```

#### POST /nodes

Crear nodo.

**Request:**
```json
{
  "patch_panel_id": "550e8400-e29b-41d4-a716-446655440011",
  "rfid_uid": "E200341502001080",
  "port_number": 1,
  "description": "Conexión oficina 101"
}
```

**Response (201):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440100",
  "patch_panel_id": "550e8400-e29b-41d4-a716-446655440011",
  "rfid_uid": "E200341502001080",
  "port_number": 1,
  "status": "active",
  "description": "Conexión oficina 101",
  "created_at": "2024-01-20T16:30:00Z"
}
```

#### GET /nodes/by-rfid/:rfid

Buscar nodo por RFID UID.

**Response (200):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440100",
  "rfid_uid": "E200341502001080",
  "port_number": 1,
  "status": "active",
  "patch_panel": {
    "name": "PP-01",
    "rack": {
      "name": "Rack Principal - Piso 1"
    }
  }
}
```

---

### 📡 Escaneos

#### POST /scans

Registrar escaneo RFID.

**Request:**
```json
{
  "rfid_code": "E200341502001080",
  "scan_type": "check_in",
  "latitude": 19.4326,
  "longitude": -99.1332,
  "notes": "Inspección rutinaria"
}
```

**Response (201):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440300",
  "node": {
    "id": "550e8400-e29b-41d4-a716-446655440100",
    "rfid_uid": "E200341502001080",
    "port_number": 1,
    "patch_panel": {
      "name": "PP-01",
      "rack": {
        "name": "Rack Principal - Piso 1"
      }
    }
  },
  "scan_type": "check_in",
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "first_name": "Juan",
    "last_name": "Pérez"
  },
  "latitude": 19.4326,
  "longitude": -99.1332,
  "scanned_at": "2024-01-20T16:45:00Z"
}
```

#### GET /scans

Listar escaneos.

**Query Params:**
- `node_id` (uuid)
- `user_id` (uuid)
- `from` (date): Fecha inicio
- `to` (date): Fecha fin
- `page` (int)
- `limit` (int)

**Response (200):**
```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440300",
      "node": {
        "id": "550e8400-e29b-41d4-a716-446655440100",
        "rfid_uid": "E200341502001080"
      },
      "scan_type": "check_in",
      "user": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "first_name": "Juan"
      },
      "latitude": 19.4326,
      "longitude": -99.1332,
      "scanned_at": "2024-01-20T16:45:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 150
  }
}
```

---

### 📊 Reportes

#### GET /reports/dashboard

Obtener métricas para dashboard.

**Response (200):**
```json
{
  "summary": {
    "total_racks": 8,
    "total_panels": 32,
    "total_nodes": 768,
    "active_nodes": 654,
    "inactive_nodes": 45,
    "maintenance_nodes": 69
  },
  "recent_scans": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440300",
      "node_rfid": "E200341502001080",
      "user_name": "Juan Pérez",
      "scanned_at": "2024-01-20T16:45:00Z"
    }
  ],
  "scans_by_day": [
    {
      "date": "2024-01-20",
      "count": 45
    },
    {
      "date": "2024-01-19",
      "count": 38
    }
  ]
}
```

#### GET /reports/activity

Reporte de actividad de técnicos.

**Query Params:**
- `from` (date)
- `to` (date)
- `user_id` (uuid)

**Response (200):**
```json
{
  "data": [
    {
      "user": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "name": "Juan Pérez"
      },
      "total_scans": 125,
      "unique_nodes": 45,
      "first_scan": "2024-01-01T08:00:00Z",
      "last_scan": "2024-01-20T16:45:00Z"
    }
  ]
}
```

---

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Bad Request - Datos inválidos |
| 401 | Unauthorized - Token inválido o expirado |
| 403 | Forbidden - Sin permisos |
| 404 | Not Found - Recurso no existe |
| 409 | Conflict - Conflicto de datos (ej: RFID duplicado) |
| 422 | Unprocessable Entity - Validación falló |
| 429 | Too Many Requests - Rate limit excedido |
| 500 | Internal Server Error |

### Formato de Error

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Datos de entrada inválidos",
    "details": [
      {
        "field": "email",
        "message": "Email es requerido"
      }
    ]
  }
}
```

## Rate Limiting

- **Límite**: 100 requests por minuto por IP
- **Headers**:
  - `X-RateLimit-Limit`: 100
  - `X-RateLimit-Remaining`: 95
  - `X-RateLimit-Reset`: 1642681200

## Paginación

Todas las listas soportan paginación:

```
GET /nodes?page=2&limit=50
```

**Headers de respuesta:**
- `X-Total-Count`: 150
- `X-Total-Pages`: 3
- `X-Current-Page`: 2

---

**Versión API**: 1.0.0
