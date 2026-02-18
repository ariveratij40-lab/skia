-- SKIA Database Schema
-- Migración inicial: Creación de tablas con RLS

-- Habilitar extensión UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Crear schema
CREATE SCHEMA IF NOT EXISTS skia;

-- Tabla de tenants (organizaciones)
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

-- Tabla de usuarios
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

-- Tabla de racks
CREATE TABLE skia.racks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id UUID NOT NULL REFERENCES skia.tenants(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de patch panels
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

-- Tabla de nodos (puertos con RFID)
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

-- Tabla de cables
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

-- Tabla de escaneos
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

-- Tabla de historial de escaneos
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

-- Habilitar Row Level Security
ALTER TABLE skia.tenants ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.users ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.racks ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.patch_panels ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.nodes ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.cables ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.scans ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.scan_history ENABLE ROW LEVEL SECURITY;

-- Crear políticas de aislamiento por tenant
CREATE POLICY tenant_isolation ON skia.tenants
    USING (id = current_setting('app.current_tenant')::UUID);

CREATE POLICY tenant_isolation ON skia.users
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

CREATE POLICY tenant_isolation ON skia.racks
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

CREATE POLICY tenant_isolation ON skia.patch_panels
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

CREATE POLICY tenant_isolation ON skia.nodes
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

CREATE POLICY tenant_isolation ON skia.cables
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

CREATE POLICY tenant_isolation ON skia.scans
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

CREATE POLICY tenant_isolation ON skia.scan_history
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

-- Crear función para establecer tenant
CREATE OR REPLACE FUNCTION skia.set_tenant(tenant_id UUID)
RETURNS VOID AS $$
BEGIN
    PERFORM set_config('app.current_tenant', tenant_id::text, false);
END;
$$ LANGUAGE plpgsql;

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
