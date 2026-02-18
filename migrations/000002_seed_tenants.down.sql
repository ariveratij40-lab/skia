-- SKIA Seed Data Rollback
-- Eliminar datos iniciales

-- Eliminar datos (en orden por dependencias)
DELETE FROM skia.scan_history WHERE tenant_id IN (SELECT id FROM skia.tenants WHERE slug = 'skia-dev');
DELETE FROM skia.scans WHERE tenant_id IN (SELECT id FROM skia.tenants WHERE slug = 'skia-dev');
DELETE FROM skia.cables WHERE tenant_id IN (SELECT id FROM skia.tenants WHERE slug = 'skia-dev');
DELETE FROM skia.nodes WHERE tenant_id IN (SELECT id FROM skia.tenants WHERE slug = 'skia-dev');
DELETE FROM skia.patch_panels WHERE tenant_id IN (SELECT id FROM skia.tenants WHERE slug = 'skia-dev');
DELETE FROM skia.racks WHERE tenant_id IN (SELECT id FROM skia.tenants WHERE slug = 'skia-dev');
DELETE FROM skia.users WHERE tenant_id IN (SELECT id FROM skia.tenants WHERE slug = 'skia-dev');
DELETE FROM skia.tenants WHERE slug = 'skia-dev';
