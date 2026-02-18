# Esquema de Base de Datos - SKIA

## Visión General

El esquema de base de datos de SKIA está diseñado para soportar multi-tenancy mediante Row Level Security (RLS) de PostgreSQL. Todos los datos están aislados por `tenant_id`.

## Schema

```sql
CREATE SCHEMA IF NOT EXISTS skia;
```

## Tablas

### tenants

Información de las organizaciones/tenants.

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | UUID | PK, autogenerado |
| name | VARCHAR(255) | Nombre del tenant |
| slug | VARCHAR(100) | Identificador único (URL-friendly) |
| db_schema | VARCHAR(100) | Schema de BD (default: 'skia') |
| is_active | BOOLEAN | Estado del tenant |
| settings | JSONB | Configuración adicional |
| created_at | TIMESTAMP | Fecha de creación |
| updated_at | TIMESTAMP | Última actualización |

```sql
CREATE TABLE skia.tenants (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(100) UNIQUE NOT NULL,
    db_schema VARCHAR(100) DEFAULT 'skia',
    is_active BOOLEAN DEFAULT true,
    settings JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### users

Usuarios del sistema.

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | UUID | PK |
| tenant_id | UUID | FK → tenants(id) |
| email | VARCHAR(255) | Email único por tenant |
| password_hash | VARCHAR(255) | Hash bcrypt |
| role | VARCHAR(50) | admin, technician, viewer |
| first_name | VARCHAR(100) | Nombre |
| last_name | VARCHAR(100) | Apellido |
| phone | VARCHAR(20) | Teléfono |
| is_active | BOOLEAN | Estado |
| last_login | TIMESTAMP | Último acceso |
| created_at | TIMESTAMP | Fecha de creación |
| updated_at | TIMESTAMP | Última actualización |

```sql
CREATE TABLE skia.users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id UUID NOT NULL REFERENCES skia.tenants(id) ON DELETE CASCADE,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'technician',
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    phone VARCHAR(20),
    is_active BOOLEAN DEFAULT true,
    last_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(tenant_id, email)
);
```

### racks

Racks de cableado.

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | UUID | PK |
| tenant_id | UUID | FK → tenants(id) |
| name | VARCHAR(255) | Nombre del rack |
| location | VARCHAR(255) | Ubicación física |
| description | TEXT | Descripción |
| created_at | TIMESTAMP | Fecha de creación |
| updated_at | TIMESTAMP | Última actualización |

```sql
CREATE TABLE skia.racks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id UUID NOT NULL REFERENCES skia.tenants(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### patch_panels

Paneles de parcheo dentro de racks.

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | UUID | PK |
| tenant_id | UUID | FK → tenants(id) |
| rack_id | UUID | FK → racks(id) |
| name | VARCHAR(255) | Nombre del panel |
| port_count | INTEGER | Número de puertos |
| description | TEXT | Descripción |
| created_at | TIMESTAMP | Fecha de creación |
| updated_at | TIMESTAMP | Última actualización |

```sql
CREATE TABLE skia.patch_panels (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id UUID NOT NULL REFERENCES skia.tenants(id) ON DELETE CASCADE,
    rack_id UUID NOT NULL REFERENCES skia.racks(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    port_count INTEGER NOT NULL DEFAULT 24,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### nodes

Nodos/puertos individuales con RFID.

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | UUID | PK |
| tenant_id | UUID | FK → tenants(id) |
| patch_panel_id | UUID | FK → patch_panels(id) |
| rfid_uid | VARCHAR(100) | UID único del tag RFID |
| port_number | INTEGER | Número de puerto en el panel |
| status | VARCHAR(50) | active, inactive, maintenance |
| description | TEXT | Descripción |
| last_scan_at | TIMESTAMP | Último escaneo |
| created_at | TIMESTAMP | Fecha de creación |
| updated_at | TIMESTAMP | Última actualización |

```sql
CREATE TABLE skia.nodes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id UUID NOT NULL REFERENCES skia.tenants(id) ON DELETE CASCADE,
    patch_panel_id UUID NOT NULL REFERENCES skia.patch_panels(id) ON DELETE CASCADE,
    rfid_uid VARCHAR(100) UNIQUE NOT NULL,
    port_number INTEGER NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    description TEXT,
    last_scan_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### cables

Cables conectados a nodos.

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | UUID | PK |
| tenant_id | UUID | FK → tenants(id) |
| node_a_id | UUID | FK → nodes(id) - Origen |
| node_b_id | UUID | FK → nodes(id) - Destino |
| cable_type | VARCHAR(50) | CAT5e, CAT6, CAT6a, fibra |
| length | DECIMAL(10,2) | Longitud en metros |
| color | VARCHAR(50) | Color del cable |
| label | VARCHAR(255) | Etiqueta identificadora |
| status | VARCHAR(50) | active, inactive |
| destination | VARCHAR(255) | Descripción del destino |
| created_at | TIMESTAMP | Fecha de creación |
| updated_at | TIMESTAMP | Última actualización |

```sql
CREATE TABLE skia.cables (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id UUID NOT NULL REFERENCES skia.tenants(id) ON DELETE CASCADE,
    node_a_id UUID NOT NULL REFERENCES skia.nodes(id) ON DELETE CASCADE,
    node_b_id UUID REFERENCES skia.nodes(id) ON DELETE SET NULL,
    cable_type VARCHAR(50) DEFAULT 'CAT6',
    length DECIMAL(10,2),
    color VARCHAR(50),
    label VARCHAR(255),
    status VARCHAR(50) DEFAULT 'active',
    destination VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### scans

Registro de escaneos RFID.

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | UUID | PK |
| tenant_id | UUID | FK → tenants(id) |
| node_id | UUID | FK → nodes(id) |
| user_id | UUID | FK → users(id) |
| scan_type | VARCHAR(50) | check_in, check_out, audit |
| latitude | DECIMAL(10,8) | Latitud GPS |
| longitude | DECIMAL(11,8) | Longitud GPS |
| notes | TEXT | Notas adicionales |
| scanned_at | TIMESTAMP | Fecha del escaneo |
| created_at | TIMESTAMP | Fecha de registro |

```sql
CREATE TABLE skia.scans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id UUID NOT NULL REFERENCES skia.tenants(id) ON DELETE CASCADE,
    node_id UUID NOT NULL REFERENCES skia.nodes(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES skia.users(id) ON DELETE CASCADE,
    scan_type VARCHAR(50) NOT NULL DEFAULT 'check_in',
    latitude DECIMAL(10,8),
    longitude DECIMAL(11,8),
    notes TEXT,
    scanned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### scan_history

Historial de cambios de estado de nodos (auditoría).

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | UUID | PK |
| tenant_id | UUID | FK → tenants(id) |
| node_id | UUID | FK → nodes(id) |
| scan_id | UUID | FK → scans(id) |
| previous_status | VARCHAR(50) | Estado anterior |
| new_status | VARCHAR(50) | Nuevo estado |
| changed_by | UUID | FK → users(id) |
| changed_at | TIMESTAMP | Fecha del cambio |
| reason | TEXT | Razón del cambio |

```sql
CREATE TABLE skia.scan_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id UUID NOT NULL REFERENCES skia.tenants(id) ON DELETE CASCADE,
    node_id UUID NOT NULL REFERENCES skia.nodes(id) ON DELETE CASCADE,
    scan_id UUID REFERENCES skia.scans(id) ON DELETE SET NULL,
    previous_status VARCHAR(50),
    new_status VARCHAR(50) NOT NULL,
    changed_by UUID NOT NULL REFERENCES skia.users(id) ON DELETE CASCADE,
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    reason TEXT
);
```

## Row Level Security (RLS)

### Habilitar RLS

```sql
-- Habilitar RLS en todas las tablas
ALTER TABLE skia.tenants ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.users ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.racks ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.patch_panels ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.nodes ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.cables ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.scans ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.scan_history ENABLE ROW LEVEL SECURITY;
```

### Políticas

```sql
-- Política para tenants (acceso a tenant propio)
CREATE POLICY tenant_isolation ON skia.tenants
    USING (id = current_setting('app.current_tenant')::UUID);

-- Política para users
CREATE POLICY tenant_isolation ON skia.users
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

-- Política para racks
CREATE POLICY tenant_isolation ON skia.racks
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

-- Política para patch_panels
CREATE POLICY tenant_isolation ON skia.patch_panels
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

-- Política para nodes
CREATE POLICY tenant_isolation ON skia.nodes
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

-- Política para cables
CREATE POLICY tenant_isolation ON skia.cables
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

-- Política para scans
CREATE POLICY tenant_isolation ON skia.scans
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

-- Política para scan_history
CREATE POLICY tenant_isolation ON skia.scan_history
    USING (tenant_id = current_setting('app.current_tenant')::UUID);
```

### Función para Establecer Tenant

```sql
CREATE OR REPLACE FUNCTION skia.set_tenant(tenant_id UUID)
RETURNS VOID AS $$
BEGIN
    PERFORM set_config('app.current_tenant', tenant_id::text, false);
END;
$$ LANGUAGE plpgsql;
```

## Índices

```sql
-- Índices para búsquedas frecuentes
CREATE INDEX idx_users_tenant_email ON skia.users(tenant_id, email);
CREATE INDEX idx_users_role ON skia.users(role);

CREATE INDEX idx_racks_tenant ON skia.racks(tenant_id);

CREATE INDEX idx_panels_rack ON skia.patch_panels(rack_id);
CREATE INDEX idx_panels_tenant ON skia.patch_panels(tenant_id);

CREATE INDEX idx_nodes_panel ON skia.nodes(patch_panel_id);
CREATE INDEX idx_nodes_rfid ON skia.nodes(rfid_uid);
CREATE INDEX idx_nodes_status ON skia.nodes(status);
CREATE INDEX idx_nodes_tenant ON skia.nodes(tenant_id);

CREATE INDEX idx_cables_node_a ON skia.cables(node_a_id);
CREATE INDEX idx_cables_node_b ON skia.cables(node_b_id);

CREATE INDEX idx_scans_node ON skia.scans(node_id);
CREATE INDEX idx_scans_user ON skia.scans(user_id);
CREATE INDEX idx_scans_scanned_at ON skia.scans(scanned_at);
CREATE INDEX idx_scans_tenant ON skia.scans(tenant_id);

CREATE INDEX idx_scan_history_node ON skia.scan_history(node_id);
CREATE INDEX idx_scan_history_changed_at ON skia.scan_history(changed_at);
```

## Diagrama ER

```
┌─────────────────┐
│    tenants      │
├─────────────────┤
│ id (PK)         │
│ name            │
│ slug            │
│ ...             │
└────────┬────────┘
         │
         │ 1:N
         │
         ▼
┌─────────────────┐       ┌─────────────────┐
│     users       │       │     racks       │
├─────────────────┤       ├─────────────────┤
│ id (PK)         │       │ id (PK)         │
│ tenant_id (FK)  │       │ tenant_id (FK)  │
│ email           │       │ name            │
│ role            │       │ location        │
│ ...             │       │ ...             │
└────────┬────────┘       └────────┬────────┘
         │                         │
         │                         │ 1:N
         │                         │
         │                         ▼
         │                ┌─────────────────┐
         │                │  patch_panels   │
         │                ├─────────────────┤
         │                │ id (PK)         │
         │                │ rack_id (FK)    │
         │                │ name            │
         │                │ port_count      │
         │                │ ...             │
         │                └────────┬────────┘
         │                         │
         │                         │ 1:N
         │                         │
         │                         ▼
         │                ┌─────────────────┐       ┌─────────────────┐
         │                │     nodes       │◄──────│     cables      │
         │                ├─────────────────┤       ├─────────────────┤
         │                │ id (PK)         │       │ id (PK)         │
         │                │ patch_panel_id  │       │ node_a_id (FK)  │
         │                │ rfid_uid (UQ)   │       │ node_b_id (FK)  │
         │                │ port_number     │       │ cable_type      │
         │                │ status          │       │ length          │
         │                │ ...             │       │ ...             │
         │                └────────┬────────┘       └─────────────────┘
         │                         │
         │                         │ 1:N
         │                         │
         └────────────────┐        │
                          │        │
                          ▼        ▼
                   ┌─────────────────┐
                   │     scans       │
                   ├─────────────────┤
                   │ id (PK)         │
                   │ node_id (FK)    │
                   │ user_id (FK)    │
                   │ scan_type       │
                   │ latitude        │
                   │ longitude       │
                   │ scanned_at      │
                   │ ...             │
                   └────────┬────────┘
                            │
                            │ 1:N
                            │
                            ▼
                   ┌─────────────────┐
                   │  scan_history   │
                   ├─────────────────┤
                   │ id (PK)         │
                   │ node_id (FK)    │
                   │ scan_id (FK)    │
                   │ previous_status │
                   │ new_status      │
                   │ changed_by (FK) │
                   │ ...             │
                   └─────────────────┘
```

---

**Versión**: 1.0.0  
**Última actualización**: 2024
