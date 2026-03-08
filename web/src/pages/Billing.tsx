import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { useLocation } from "wouter";
import { ArrowLeft, Check, X } from "lucide-react";

const plans = [
  {
    name: "Free",
    price: "$0",
    description: "Para empezar",
    features: [
      { name: "100 nodos", included: true },
      { name: "1 edificio", included: true },
      { name: "Planos básicos", included: true },
      { name: "5 usuarios", included: false },
      { name: "API pública", included: false },
      { name: "Soporte prioritario", included: false },
    ],
  },
  {
    name: "Starter",
    price: "$99",
    period: "/mes",
    description: "Para pequeños equipos",
    popular: true,
    features: [
      { name: "1,000 nodos", included: true },
      { name: "5 edificios", included: true },
      { name: "Planos avanzados", included: true },
      { name: "25 usuarios", included: true },
      { name: "API pública", included: false },
      { name: "Soporte prioritario", included: false },
    ],
  },
  {
    name: "Core",
    price: "$299",
    period: "/mes",
    description: "Para empresas",
    features: [
      { name: "Nodos ilimitados", included: true },
      { name: "Edificios ilimitados", included: true },
      { name: "Planos multicapa", included: true },
      { name: "Usuarios ilimitados", included: true },
      { name: "API pública", included: true },
      { name: "Soporte prioritario", included: true },
    ],
  },
];

export default function Billing() {
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
            Facturación
          </h1>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Current Plan */}
        <Card className="mb-8 bg-blue-50 dark:bg-blue-900/20 border-blue-200 dark:border-blue-800">
          <CardHeader>
            <CardTitle>Plan Actual</CardTitle>
            <CardDescription>
              Estás en el plan Trial con acceso completo
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="flex items-center justify-between">
              <div>
                <p className="text-2xl font-bold text-slate-900 dark:text-white">
                  Trial
                </p>
                <p className="text-sm text-slate-600 dark:text-slate-400 mt-1">
                  Vence en 14 días
                </p>
              </div>
              <Button>Cambiar Plan</Button>
            </div>
          </CardContent>
        </Card>

        {/* Plans */}
        <div className="mb-8">
          <h2 className="text-2xl font-bold text-slate-900 dark:text-white mb-4">
            Planes Disponibles
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {plans.map((plan) => (
              <Card
                key={plan.name}
                className={`relative ${
                  plan.popular
                    ? "border-blue-600 dark:border-blue-400 shadow-lg"
                    : ""
                }`}
              >
                {plan.popular && (
                  <div className="absolute -top-3 left-1/2 transform -translate-x-1/2">
                    <span className="bg-blue-600 text-white px-3 py-1 rounded-full text-xs font-semibold">
                      Popular
                    </span>
                  </div>
                )}
                <CardHeader>
                  <CardTitle>{plan.name}</CardTitle>
                  <CardDescription>{plan.description}</CardDescription>
                  <div className="mt-4">
                    <span className="text-3xl font-bold text-slate-900 dark:text-white">
                      {plan.price}
                    </span>
                    {plan.period && (
                      <span className="text-slate-600 dark:text-slate-400">
                        {plan.period}
                      </span>
                    )}
                  </div>
                </CardHeader>
                <CardContent className="space-y-4">
                  <Button
                    className="w-full"
                    variant={plan.popular ? "default" : "outline"}
                  >
                    {plan.name === "Free" ? "Degradar" : "Cambiar a este plan"}
                  </Button>
                  <div className="space-y-3">
                    {plan.features.map((feature) => (
                      <div key={feature.name} className="flex items-center gap-3">
                        {feature.included ? (
                          <Check className="h-5 w-5 text-green-600 dark:text-green-400" />
                        ) : (
                          <X className="h-5 w-5 text-slate-400 dark:text-slate-600" />
                        )}
                        <span
                          className={
                            feature.included
                              ? "text-slate-900 dark:text-white"
                              : "text-slate-500 dark:text-slate-500"
                          }
                        >
                          {feature.name}
                        </span>
                      </div>
                    ))}
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>

        {/* Invoices */}
        <Card>
          <CardHeader>
            <CardTitle>Historial de Facturas</CardTitle>
            <CardDescription>
              Descarga tus facturas anteriores
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="space-y-4">
              <div className="flex items-center justify-between p-4 bg-slate-50 dark:bg-slate-900 rounded-lg">
                <div>
                  <p className="font-medium text-slate-900 dark:text-white">
                    Factura #001
                  </p>
                  <p className="text-sm text-slate-600 dark:text-slate-400">
                    Marzo 5, 2026
                  </p>
                </div>
                <Button variant="outline" size="sm">
                  Descargar
                </Button>
              </div>
              <div className="text-center py-8 text-slate-600 dark:text-slate-400">
                <p>No hay más facturas disponibles</p>
              </div>
            </div>
          </CardContent>
        </Card>

        {/* Payment Method */}
        <Card className="mt-6">
          <CardHeader>
            <CardTitle>Método de Pago</CardTitle>
            <CardDescription>
              Actualiza tu método de pago
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="flex items-center justify-between p-4 bg-slate-50 dark:bg-slate-900 rounded-lg">
              <div>
                <p className="font-medium text-slate-900 dark:text-white">
                  Visa •••• 4242
                </p>
                <p className="text-sm text-slate-600 dark:text-slate-400">
                  Vence en 12/2026
                </p>
              </div>
              <Button variant="outline">Cambiar</Button>
            </div>
          </CardContent>
        </Card>
      </main>
    </div>
  );
}
