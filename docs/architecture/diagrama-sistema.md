# Arquitectura del Sistema SKIA

## Visión General

SKIA sigue una arquitectura de microservicios ligera con separación clara de responsabilidades y multi-tenancy mediante Row Level Security (RLS) en PostgreSQL.

## Diagrama de Componentes

```
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                                    CLIENTES                                         │
├─────────────────────────────────────────────────────────────────────────────────────┤
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐ │
│  │  Técnico de     │  │  Administrador  │  │   Operador de   │  │   Videowall     │ │
│  │  Campo          │  │  de Sistema     │  │   Monitoreo     │  │   (Pantalla)    │ │
│  │                 │  │                 │  │                 │  │                 │ │
│  │  Zebra MC22U    │  │  Laptop/PC      │  │   Workstation   │  │   TV/Monitor    │ │
│  │  (Android)      │  │  (Navegador)    │  │   (Navegador)   │  │   (Navegador)   │ │
│  └────────┬────────┘  └────────┬────────┘  └────────┬────────┘  └────────┬────────┘ │
│           │                    │                    │                    │          │
│           │  HTTP/HTTPS        │  HTTP/HTTPS        │  HTTP/HTTPS        │          │
│           │  (REST API)        │  (REST API)        │  (REST API)        │          │
│           │                    │                    │                    │          │
└───────────┼────────────────────┼────────────────────┼────────────────────┼──────────┘
            │                    │                    │                    │
            └────────────────────┴────────────────────┴────────────────────┘
                                      │
                                      ▼
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                                API GATEWAY                                          │
│  ┌─────────────────────────────────────────────────────────────────────────────┐    │
│  │                         Nginx Reverse Proxy                                  │    │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐ │    │
│  │  │  SSL/TLS    │  │ Rate Limit  │  │   CORS      │  │  Load Balancing     │ │    │
│  │  │  Termination│  │  Config     │  │  Headers    │  │  (si aplica)        │ │    │
│  │  └─────────────┘  └─────────────┘  └─────────────┘  └─────────────────────┘ │    │
│  └─────────────────────────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────────────────────────┘
                                      │
                                      ▼
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                              BACKEND API (Go)                                       │
│  ┌─────────────────────────────────────────────────────────────────────────────┐    │
│  │                         HTTP Router (Gin/Echo)                               │    │
│  └─────────────────────────────────────────────────────────────────────────────┘    │
│                                       │                                             │
│  ┌────────────────────────────────────┼────────────────────────────────────────┐    │
│  │                                    │                                        │    │
│  │  ┌─────────────────┐    ┌─────────▼──────────┐    ┌─────────────────────┐  │    │
│  │  │   Middleware    │    │      Handlers      │    │      Services       │  │    │
│  │  │                 │    │                    │    │                     │  │    │
│  │  │ • JWT Auth      │◄──►│ • Auth Handler     │◄──►│ • Auth Service      │  │    │
│  │  │ • RLS Context   │    │ • User Handler     │    │ • User Service      │  │    │
│  │  │ • Logging       │    │ • Node Handler     │    │ • Node Service      │  │    │
│  │  │ • Rate Limit    │    │ • Rack Handler     │    │ • Rack Service      │  │    │
│  │  │ • Validation    │    │ • Scan Handler     │    │ • Scan Service      │  │    │
│  │  │                 │    │ • Report Handler   │    │ • Report Service    │  │    │
│  │  └─────────────────┘    └────────────────────┘    └─────────────────────┘  │    │
│  │                                    │                                        │    │
│  │                                    ▼                                        │    │
│  │                         ┌─────────────────────┐                            │    │
│  │                         │     Repository      │                            │    │
│  │                         │                     │                            │    │
│  │                         │ • User Repository   │                            │    │
│  │                         │ • Node Repository   │                            │    │
│  │                         │ • Rack Repository   │                            │    │
│  │                         │ • Scan Repository   │                            │    │
│  │                         └──────────┬──────────┘                            │    │
│  └────────────────────────────────────┼────────────────────────────────────────┘    │
│                                       │                                             │
└───────────────────────────────────────┼─────────────────────────────────────────────┘
                                        │
                                        │ SQL + RLS
                                        │ (pgx driver)
                                        ▼
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                           DATA LAYER                                                │
│                                                                                     │
│  ┌─────────────────────────────┐    ┌─────────────────────┐    ┌─────────────────┐  │
│  │      PostgreSQL 15+         │    │       Redis         │    │     MinIO       │  │
│  │                             │    │                     │    │                 │  │
│  │  ┌─────────────────────┐    │    │  ┌─────────────┐    │    │  ┌───────────┐  │  │
│  │  │   skia Schema       │    │    │  │  Sessions   │    │    │  │  Planos   │  │  │
│  │  │                     │    │    │  │  JWT Cache  │    │    │  │  Fotos    │  │  │
│  │  │ • tenants           │    │    │  │  Rate Limit │    │    │  │  Docs     │  │  │
│  │  │ • users             │    │    │  └─────────────┘    │    │  └───────────┘  │  │
│  │  │ • racks             │    │    └─────────────────────┘    └─────────────────┘  │
│  │  │ • patch_panels      │    │                                                      │
│  │  │ • nodes             │    │                                                      │
│  │  │ • cables            │    │                                                      │
│  │  │ • scans             │    │                                                      │
│  │  │ • scan_history      │    │                                                      │
│  │  │                     │    │                                                      │
│  │  │  RLS Policies:      │    │                                                      │
│  │  │  tenant_isolation   │    │                                                      │
│  │  └─────────────────────┘    │                                                      │
│  └─────────────────────────────┘                                                      │
│                                                                                     │
└─────────────────────────────────────────────────────────────────────────────────────┘
```

## Flujo de Autenticación

```
┌─────────┐                    ┌─────────────┐                    ┌─────────────┐
│ Cliente │                    │   Backend   │                    │  PostgreSQL │
└────┬────┘                    └──────┬──────┘                    └──────┬──────┘
     │                                │                                  │
     │  1. POST /auth/login           │                                  │
     │  {email, password}             │                                  │
     │ ─────────────────────────────>│                                  │
     │                                │                                  │
     │                                │  2. Validar credenciales         │
     │                                │  SELECT * FROM users WHERE...    │
     │                                │ ────────────────────────────────>│
     │                                │                                  │
     │                                │  3. Retornar usuario             │
     │                                │ <────────────────────────────────│
     │                                │                                  │
     │                                │  4. Generar JWT                  │
     │                                │  {user_id, tenant_id, role}      │
     │                                │                                  │
     │  5. Retornar token             │                                  │
     │  {access_token, refresh_token} │                                  │
     │ <─────────────────────────────│                                  │
     │                                │                                  │
     │  6. Request con Authorization: │                                  │
     │     Bearer <token>             │                                  │
     │ ─────────────────────────────>│                                  │
     │                                │                                  │
     │                                │  7. Validar JWT                  │
     │                                │  Extraer tenant_id del token     │
     │                                │                                  │
     │                                │  8. SET app.current_tenant       │
     │                                │  SET SESSION...                  │
     │                                │ ────────────────────────────────>│
     │                                │                                  │
     │                                │  9. Ejecutar query               │
     │                                │  RLS filtra automáticamente      │
     │                                │ ────────────────────────────────>│
     │                                │                                  │
     │                                │  10. Retornar resultados         │
     │                                │ <────────────────────────────────│
     │                                │                                  │
     │  11. Response filtrada         │                                  │
     │      por tenant                │                                  │
     │ <─────────────────────────────│                                  │
     │                                │                                  │
```

## Flujo de Escaneo RFID

```
┌─────────────┐         ┌─────────────┐         ┌─────────────┐         ┌─────────────┐
│ Zebra MC22U │         │  App Flutter│         │   Backend   │         │  PostgreSQL │
└──────┬──────┘         └──────┬──────┘         └──────┬──────┘         └──────┬──────┘
       │                       │                       │                       │
       │  1. Escanear RFID     │                       │                       │
       │  (Botón físico)       │                       │                       │
       │ ─────────────────────>│                       │                       │
       │                       │                       │                       │
       │                       │  2. DataWedge API     │                       │
       │                       │  Captura código RFID  │                       │
       │                       │                       │                       │
       │                       │  3. POST /api/scans   │                       │
       │                       │  {rfid_code, lat, lng}│                       │
       │                       │ ─────────────────────>│                       │
       │                       │                       │                       │
       │                       │                       │  4. Buscar nodo       │
       │                       │                       │  SELECT * FROM nodes  │
       │                       │                       │  WHERE rfid_uid = ?   │
       │                       │                       │ ─────────────────────>│
       │                       │                       │                       │
       │                       │                       │  5. Retornar nodo     │
       │                       │                       │ <─────────────────────│
       │                       │                       │                       │
       │                       │                       │  6. Crear registro    │
       │                       │                       │  INSERT INTO scans... │
       │                       │                       │  INSERT INTO scan_... │
       │                       │                       │ ─────────────────────>│
       │                       │                       │                       │
       │                       │  7. Retornar info     │                       │
       │                       │  del nodo escaneado   │                       │
       │                       │ <─────────────────────│                       │
       │                       │                       │                       │
       │                       │  8. Mostrar datos     │                       │
       │                       │  del nodo en UI       │                       │
       │ <─────────────────────│                       │                       │
       │                       │                       │                       │
```

## Modelo de Datos (Simplificado)

```
┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐
│    tenants      │       │     users       │       │     racks       │
├─────────────────┤       ├─────────────────┤       ├─────────────────┤
│ id (PK)         │◄──────│ tenant_id (FK)  │       │ tenant_id (FK)  │
│ name            │       │ id (PK)         │       │ id (PK)         │
│ slug            │       │ email           │       │ name            │
│ db_schema       │       │ password_hash   │       │ location        │
│ is_active       │       │ role            │       │ description     │
│ created_at      │       │ tenant_id       │       │ created_at      │
└─────────────────┘       │ created_at      │       └────────┬────────┘
                          └─────────────────┘                │
                                                             │
                          ┌─────────────────┐                │
                          │  patch_panels   │◄───────────────┘
                          ├─────────────────┤
                          │ tenant_id (FK)  │
                          │ id (PK)         │
                          │ rack_id (FK)    │
                          │ name            │
                          │ port_count      │
                          └────────┬────────┘
                                   │
                                   │
                          ┌────────▼────────┐       ┌─────────────────┐
                          │     nodes       │       │     cables      │
                          ├─────────────────┤       ├─────────────────┤
                          │ tenant_id (FK)  │       │ tenant_id (FK)  │
                          │ id (PK)         │◄──────│ node_a_id (FK)  │
                          │ patch_panel_id  │       │ node_b_id (FK)  │
                          │ rfid_uid (UQ)   │       │ cable_type      │
                          │ port_number     │       │ status          │
                          │ status          │       │ length          │
                          └────────┬────────┘       └─────────────────┘
                                   │
                                   │
                          ┌────────▼────────┐
                          │     scans       │
                          ├─────────────────┤
                          │ tenant_id (FK)  │
                          │ id (PK)         │
                          │ node_id (FK)    │
                          │ user_id (FK)    │
                          │ scan_type       │
                          │ latitude        │
                          │ longitude       │
                          │ scanned_at      │
                          └─────────────────┘
```

## Multi-tenancy con RLS

### Configuración

```sql
-- Habilitar RLS en todas las tablas
ALTER TABLE skia.users ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.racks ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.nodes ENABLE ROW LEVEL SECURITY;
ALTER TABLE skia.scans ENABLE ROW LEVEL SECURITY;

-- Crear política de aislamiento
CREATE POLICY tenant_isolation ON skia.users
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

-- Aplicar a todas las tablas
CREATE POLICY tenant_isolation ON skia.racks
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

CREATE POLICY tenant_isolation ON skia.nodes
    USING (tenant_id = current_setting('app.current_tenant')::UUID);

CREATE POLICY tenant_isolation ON skia.scans
    USING (tenant_id = current_setting('app.current_tenant')::UUID);
```

### Uso en Backend

```go
// Middleware para establecer tenant
func RLSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tenantID := c.GetString("tenant_id") // del JWT
        
        // Establecer en contexto de PostgreSQL
        db.Exec("SET LOCAL app.current_tenant = $1", tenantID)
        
        c.Next()
    }
}
```

## Escalabilidad

### Horizontal
- **Stateless Backend**: Múltiples instancias del API Go
- **Load Balancer**: Nginx distribuye carga
- **Database**: Read replicas para consultas

### Vertical
- **Caching**: Redis para sesiones y datos frecuentes
- **CDN**: Assets estáticos y planos
- **Object Storage**: MinIO para archivos

## Monitoreo

```
┌─────────────────┐
│   Prometheus    │────┐
│   (Métricas)    │    │
└─────────────────┘    │    ┌─────────────────┐
                       └───►│    Grafana      │
┌─────────────────┐    │    │  (Dashboards)   │
│    Loki         │────┘    └─────────────────┘
│    (Logs)       │
└─────────────────┘
```

## Consideraciones de Seguridad

1. **Autenticación**: JWT con expiración corta (15 min)
2. **Refresh Tokens**: Rotación automática
3. **RLS**: Aislamiento estricto por tenant
4. **HTTPS**: Todo el tráfico encriptado
5. **Rate Limiting**: 100 req/min por IP
6. **Validación**: Input sanitization en todos los endpoints

---

**Versión**: 1.0.0  
**Última actualización**: 2024
