-- SKIA Database Schema Rollback
-- Revertir migración inicial

-- Eliminar políticas RLS
DROP POLICY IF EXISTS tenant_isolation ON skia.scan_history;
DROP POLICY IF EXISTS tenant_isolation ON skia.scans;
DROP POLICY IF EXISTS tenant_isolation ON skia.cables;
DROP POLICY IF EXISTS tenant_isolation ON skia.nodes;
DROP POLICY IF EXISTS tenant_isolation ON skia.patch_panels;
DROP POLICY IF EXISTS tenant_isolation ON skia.racks;
DROP POLICY IF EXISTS tenant_isolation ON skia.users;
DROP POLICY IF EXISTS tenant_isolation ON skia.tenants;

-- Eliminar función
DROP FUNCTION IF EXISTS skia.set_tenant(UUID);

-- Eliminar tablas
DROP TABLE IF EXISTS skia.scan_history;
DROP TABLE IF EXISTS skia.scans;
DROP TABLE IF EXISTS skia.cables;
DROP TABLE IF EXISTS skia.nodes;
DROP TABLE IF EXISTS skia.patch_panels;
DROP TABLE IF EXISTS skia.racks;
DROP TABLE IF EXISTS skia.users;
DROP TABLE IF EXISTS skia.tenants;

-- Eliminar schema
DROP SCHEMA IF EXISTS skia;
