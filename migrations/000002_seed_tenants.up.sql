-- SKIA Seed Data
-- Datos iniciales para desarrollo

-- Insertar tenant de desarrollo
INSERT INTO skia.tenants (name, slug, settings) VALUES
('SKIA Development', 'skia-dev', '{"timezone": "America/Mexico_City"}');

-- Obtener ID del tenant
DO $$
DECLARE
    tenant_id UUID;
BEGIN
    SELECT id INTO tenant_id FROM skia.tenants WHERE slug = 'skia-dev';
    
    -- Insertar usuario admin
    INSERT INTO skia.users (tenant_id, email, password_hash, role, first_name, last_name, is_active)
    VALUES (
        tenant_id,
        'admin@skia.local',
        '$2a$10$YourHashedPasswordHere', -- Cambiar por hash real
        'admin',
        'Administrador',
        'SKIA',
        true
    );
    
    -- Insertar racks de ejemplo
    INSERT INTO skia.racks (tenant_id, name, location, description)
    VALUES 
        (tenant_id, 'Rack Principal - Piso 1', 'Sala de Servidores A', 'Rack principal de telecomunicaciones'),
        (tenant_id, 'Rack Secundario - Piso 2', 'Sala de Servidores B', 'Rack secundario para expansión');
    
END $$;
