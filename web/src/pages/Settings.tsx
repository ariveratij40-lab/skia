import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useLocation } from "wouter";
import { ArrowLeft } from "lucide-react";

export default function Settings() {
  const [, setLocation] = useLocation();

  return (
    <div className="min-h-screen bg-slate-50 dark:bg-slate-950">
      {/* Header */}
      <header className="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4 flex items-center gap-4">
          <button
            onClick={() => setLocation("/dashboard")}
            className="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-lg transition"
          >
            <ArrowLeft className="h-5 w-5" />
          </button>
          <h1 className="text-2xl font-bold text-slate-900 dark:text-white">
            Configuración
          </h1>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Tenant Settings */}
        <Card className="mb-6">
          <CardHeader>
            <CardTitle>Información de la Empresa</CardTitle>
            <CardDescription>
              Personaliza los datos de tu empresa
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="company">Nombre de la Empresa</Label>
                <Input
                  id="company"
                  defaultValue="Demo Company"
                  placeholder="Mi Empresa S.A."
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="industry">Industria</Label>
                <Input
                  id="industry"
                  placeholder="Selecciona una industria"
                />
              </div>
            </div>
            <div className="space-y-2">
              <Label htmlFor="description">Descripción</Label>
              <textarea
                id="description"
                className="w-full px-3 py-2 border border-slate-300 dark:border-slate-700 rounded-lg bg-white dark:bg-slate-900 text-slate-900 dark:text-white"
                rows={3}
                placeholder="Describe tu empresa"
              />
            </div>
            <Button>Guardar Cambios</Button>
          </CardContent>
        </Card>

        {/* Appearance Settings */}
        <Card className="mb-6">
          <CardHeader>
            <CardTitle>Personalización</CardTitle>
            <CardDescription>
              Personaliza la apariencia de SKIA
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="logo">Logo de la Empresa</Label>
              <div className="flex items-center gap-4">
                <div className="w-20 h-20 bg-slate-200 dark:bg-slate-800 rounded-lg flex items-center justify-center">
                  <span className="text-2xl font-bold text-slate-600 dark:text-slate-400">
                    DC
                  </span>
                </div>
                <Button variant="outline">Cambiar Logo</Button>
              </div>
            </div>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="primaryColor">Color Primario</Label>
                <div className="flex items-center gap-2">
                  <div className="w-10 h-10 bg-blue-600 rounded-lg"></div>
                  <Input
                    id="primaryColor"
                    type="color"
                    defaultValue="#2563eb"
                    className="w-20 h-10 p-1"
                  />
                </div>
              </div>
              <div className="space-y-2">
                <Label htmlFor="accentColor">Color de Acento</Label>
                <div className="flex items-center gap-2">
                  <div className="w-10 h-10 bg-green-600 rounded-lg"></div>
                  <Input
                    id="accentColor"
                    type="color"
                    defaultValue="#16a34a"
                    className="w-20 h-10 p-1"
                  />
                </div>
              </div>
            </div>
            <Button>Guardar Personalización</Button>
          </CardContent>
        </Card>

        {/* User Settings */}
        <Card className="mb-6">
          <CardHeader>
            <CardTitle>Mi Perfil</CardTitle>
            <CardDescription>
              Actualiza tu información personal
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="fullName">Nombre Completo</Label>
                <Input
                  id="fullName"
                  defaultValue="Demo User"
                  placeholder="Tu nombre"
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="email">Correo Electrónico</Label>
                <Input
                  id="email"
                  type="email"
                  defaultValue="demo@skia.com"
                  placeholder="tu@email.com"
                />
              </div>
            </div>
            <div className="space-y-2">
              <Label htmlFor="phone">Teléfono</Label>
              <Input
                id="phone"
                type="tel"
                placeholder="+1 (555) 000-0000"
              />
            </div>
            <Button>Guardar Perfil</Button>
          </CardContent>
        </Card>

        {/* Security Settings */}
        <Card className="mb-6">
          <CardHeader>
            <CardTitle>Seguridad</CardTitle>
            <CardDescription>
              Gestiona tu seguridad y privacidad
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label>Cambiar Contraseña</Label>
              <Button variant="outline">Cambiar Contraseña</Button>
            </div>
            <div className="space-y-2">
              <Label>Autenticación de Dos Factores</Label>
              <Button variant="outline">Habilitar 2FA</Button>
            </div>
            <div className="space-y-2">
              <Label>Sesiones Activas</Label>
              <Button variant="outline">Cerrar Todas las Sesiones</Button>
            </div>
          </CardContent>
        </Card>

        {/* Notifications Settings */}
        <Card className="mb-6">
          <CardHeader>
            <CardTitle>Notificaciones</CardTitle>
            <CardDescription>
              Configura cómo deseas recibir notificaciones
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-3">
              <div className="flex items-center justify-between">
                <div>
                  <p className="font-medium text-slate-900 dark:text-white">
                    Alertas por Email
                  </p>
                  <p className="text-sm text-slate-600 dark:text-slate-400">
                    Recibe alertas importantes por email
                  </p>
                </div>
                <input type="checkbox" defaultChecked className="w-4 h-4" />
              </div>
              <div className="flex items-center justify-between">
                <div>
                  <p className="font-medium text-slate-900 dark:text-white">
                    Notificaciones Push
                  </p>
                  <p className="text-sm text-slate-600 dark:text-slate-400">
                    Recibe notificaciones en tu navegador
                  </p>
                </div>
                <input type="checkbox" defaultChecked className="w-4 h-4" />
              </div>
              <div className="flex items-center justify-between">
                <div>
                  <p className="font-medium text-slate-900 dark:text-white">
                    Newsletter
                  </p>
                  <p className="text-sm text-slate-600 dark:text-slate-400">
                    Recibe actualizaciones y consejos
                  </p>
                </div>
                <input type="checkbox" className="w-4 h-4" />
              </div>
            </div>
            <Button>Guardar Preferencias</Button>
          </CardContent>
        </Card>

        {/* Danger Zone */}
        <Card className="border-red-200 dark:border-red-900">
          <CardHeader>
            <CardTitle className="text-red-600 dark:text-red-400">
              Zona de Peligro
            </CardTitle>
            <CardDescription>
              Acciones irreversibles
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <p className="text-sm text-slate-600 dark:text-slate-400 mb-2">
                Eliminar tu cuenta y todos los datos asociados
              </p>
              <Button variant="destructive">Eliminar Cuenta</Button>
            </div>
          </CardContent>
        </Card>
      </main>
    </div>
  );
}
