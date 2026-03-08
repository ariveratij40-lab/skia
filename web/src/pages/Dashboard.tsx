import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { useLocation } from "wouter";
import { useEffect, useState } from "react";
import { LogOut, Settings, Bell, Package, Map, Users, CreditCard, FileText } from "lucide-react";

interface UserData {
  id: string;
  email: string;
  full_name: string;
  role: string;
}

interface TenantData {
  id: string;
  name: string;
  plan: string;
  trial_ends_at: string;
}

export default function Dashboard() {
  const [, setLocation] = useLocation();
  const [userData, setUserData] = useState<UserData | null>(null);
  const [tenantData, setTenantData] = useState<TenantData | null>(null);
  const [daysLeft, setDaysLeft] = useState(0);

  useEffect(() => {
    // Verificar autenticación
    const token = localStorage.getItem("accessToken");
    if (!token) {
      setLocation("/login");
      return;
    }

    // Cargar datos del usuario
    const user = localStorage.getItem("userData");
    const tenant = localStorage.getItem("tenantData");

    if (user) setUserData(JSON.parse(user));
    if (tenant) {
      const tenantInfo = JSON.parse(tenant);
      setTenantData(tenantInfo);

      // Calcular días restantes del trial
      if (tenantInfo.trial_ends_at) {
        const trialEnd = new Date(tenantInfo.trial_ends_at);
        const today = new Date();
        const days = Math.ceil((trialEnd.getTime() - today.getTime()) / (1000 * 60 * 60 * 24));
        setDaysLeft(Math.max(days, 0));
      }
    }
  }, [setLocation]);

  const handleLogout = () => {
    if (confirm("¿Estás seguro de que deseas cerrar sesión?")) {
      localStorage.removeItem("accessToken");
      localStorage.removeItem("userData");
      localStorage.removeItem("tenantData");
      setLocation("/login");
    }
  };

  const navigateTo = (path: string) => {
    setLocation(path);
  };

  if (!userData || !tenantData) {
    return (
      <div className="min-h-screen bg-slate-50 dark:bg-slate-950 flex items-center justify-center">
        <div className="text-center">
          <div className="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
          <p className="mt-4 text-slate-600 dark:text-slate-400">Cargando...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-slate-50 dark:bg-slate-950">
      {/* Header */}
      <header className="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4 flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div className="px-3 py-1 bg-blue-600 text-white rounded-lg font-bold text-lg">
              SKIA
            </div>
            <h1 className="text-xl font-bold text-slate-900 dark:text-white">
              {tenantData.name}
            </h1>
          </div>

          <div className="flex items-center gap-4">
            <button className="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-lg transition">
              <Bell className="h-5 w-5 text-slate-600 dark:text-slate-400" />
            </button>
            <button
              onClick={() => navigateTo("/settings")}
              className="p-2 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-lg transition"
            >
              <Settings className="h-5 w-5 text-slate-600 dark:text-slate-400" />
            </button>
            <button
              onClick={handleLogout}
              className="p-2 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition"
            >
              <LogOut className="h-5 w-5 text-red-600 dark:text-red-400" />
            </button>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Welcome Section */}
        <div className="mb-8">
          <h2 className="text-3xl font-bold text-slate-900 dark:text-white mb-2">
            Bienvenido, {userData.full_name}
          </h2>
          <p className="text-slate-600 dark:text-slate-400">
            Aquí está un resumen de tu infraestructura
          </p>
        </div>

        {/* Stats Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
          {/* Plan Card */}
          <Card>
            <CardHeader className="pb-3">
              <CardTitle className="text-sm font-medium text-slate-600 dark:text-slate-400">
                Plan Actual
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold text-slate-900 dark:text-white">
                {tenantData.plan === "trial" ? "Trial" : tenantData.plan}
              </div>
              <p className="text-xs text-slate-600 dark:text-slate-400 mt-1">
                {daysLeft > 0 ? `${daysLeft} días restantes` : "Vencido"}
              </p>
            </CardContent>
          </Card>

          {/* Nodes Card */}
          <Card>
            <CardHeader className="pb-3">
              <CardTitle className="text-sm font-medium text-slate-600 dark:text-slate-400">
                Nodos Activos
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold text-slate-900 dark:text-white">
                0
              </div>
              <p className="text-xs text-slate-600 dark:text-slate-400 mt-1">
                de 100 disponibles
              </p>
            </CardContent>
          </Card>

          {/* Plans Card */}
          <Card>
            <CardHeader className="pb-3">
              <CardTitle className="text-sm font-medium text-slate-600 dark:text-slate-400">
                Planos
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold text-slate-900 dark:text-white">
                0
              </div>
              <p className="text-xs text-slate-600 dark:text-slate-400 mt-1">
                creados
              </p>
            </CardContent>
          </Card>

          {/* Users Card */}
          <Card>
            <CardHeader className="pb-3">
              <CardTitle className="text-sm font-medium text-slate-600 dark:text-slate-400">
                Usuarios
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold text-slate-900 dark:text-white">
                1
              </div>
              <p className="text-xs text-slate-600 dark:text-slate-400 mt-1">
                en tu equipo
              </p>
            </CardContent>
          </Card>
        </div>

        {/* Quick Actions */}
        <div className="mb-8">
          <h3 className="text-lg font-semibold text-slate-900 dark:text-white mb-4">
            Acciones Rápidas
          </h3>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            {/* Inventory Card */}
            <Card className="hover:shadow-lg transition cursor-pointer" onClick={() => navigateTo("/inventory")}>
              <CardHeader>
                <div className="flex items-center justify-between">
                  <CardTitle className="text-lg">Inventario</CardTitle>
                  <Package className="h-5 w-5 text-blue-600 dark:text-blue-400" />
                </div>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-slate-600 dark:text-slate-400">
                  Gestiona nodos, racks y activos
                </p>
              </CardContent>
            </Card>

            {/* Plans Card */}
            <Card className="hover:shadow-lg transition cursor-pointer" onClick={() => navigateTo("/plans")}>
              <CardHeader>
                <div className="flex items-center justify-between">
                  <CardTitle className="text-lg">Planos</CardTitle>
                  <Map className="h-5 w-5 text-green-600 dark:text-green-400" />
                </div>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-slate-600 dark:text-slate-400">
                  Diseña planos de infraestructura
                </p>
              </CardContent>
            </Card>

            {/* Users Card */}
            <Card className="hover:shadow-lg transition cursor-pointer" onClick={() => navigateTo("/users")}>
              <CardHeader>
                <div className="flex items-center justify-between">
                  <CardTitle className="text-lg">Usuarios</CardTitle>
                  <Users className="h-5 w-5 text-purple-600 dark:text-purple-400" />
                </div>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-slate-600 dark:text-slate-400">
                  Invita y gestiona usuarios
                </p>
              </CardContent>
            </Card>
          </div>
        </div>

        {/* Secondary Actions */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {/* Billing Card */}
          <Card className="hover:shadow-lg transition cursor-pointer" onClick={() => navigateTo("/billing")}>
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle>Facturación</CardTitle>
                  <CardDescription>Gestiona tu suscripción y facturas</CardDescription>
                </div>
                <CreditCard className="h-5 w-5 text-orange-600 dark:text-orange-400" />
              </div>
            </CardHeader>
          </Card>

          {/* Documentation Card */}
          <Card className="hover:shadow-lg transition cursor-pointer">
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle>Documentación</CardTitle>
                  <CardDescription>Aprende cómo usar SKIA</CardDescription>
                </div>
                <FileText className="h-5 w-5 text-indigo-600 dark:text-indigo-400" />
              </div>
            </CardHeader>
          </Card>
        </div>

        {/* Recent Activity */}
        <div className="mt-8">
          <h3 className="text-lg font-semibold text-slate-900 dark:text-white mb-4">
            Actividad Reciente
          </h3>
          <Card>
            <CardContent className="pt-6">
              <div className="space-y-4">
                <div className="flex items-center justify-between pb-4 border-b border-slate-200 dark:border-slate-800">
                  <div>
                    <p className="font-medium text-slate-900 dark:text-white">Cuenta creada</p>
                    <p className="text-sm text-slate-600 dark:text-slate-400">
                      {new Date().toLocaleDateString("es-ES")}
                    </p>
                  </div>
                  <span className="px-3 py-1 bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-400 text-xs font-medium rounded-full">
                    Exitoso
                  </span>
                </div>
                <div className="flex items-center justify-between">
                  <div>
                    <p className="font-medium text-slate-900 dark:text-white">Trial iniciado</p>
                    <p className="text-sm text-slate-600 dark:text-slate-400">
                      14 días de acceso completo
                    </p>
                  </div>
                  <span className="px-3 py-1 bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-400 text-xs font-medium rounded-full">
                    Activo
                  </span>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </main>
    </div>
  );
}
