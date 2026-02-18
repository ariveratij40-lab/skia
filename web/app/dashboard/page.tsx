'use client'

import { useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { useAuthStore } from '@/lib/store/auth-store'
import { DashboardStats } from '@/components/dashboard/stats'
import { RecentScans } from '@/components/dashboard/recent-scans'
import { Sidebar } from '@/components/layout/sidebar'
import { Header } from '@/components/layout/header'

export default function DashboardPage() {
  const router = useRouter()
  const { isAuthenticated, user } = useAuthStore()

  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/login')
    }
  }, [isAuthenticated, router])

  if (!isAuthenticated) {
    return null
  }

  return (
    <div className="min-h-screen bg-gray-100">
      <Sidebar />
      <div className="ml-64">
        <Header user={user} />
        <main className="p-6">
          <div className="mb-6">
            <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
            <p className="text-gray-600">Resumen de la infraestructura</p>
          </div>

          <DashboardStats />

          <div className="mt-8 grid grid-cols-1 lg:grid-cols-2 gap-6">
            <RecentScans />
            <div className="bg-white rounded-lg shadow p-6">
              <h2 className="text-lg font-semibold mb-4">Actividad por Día</h2>
              <p className="text-gray-500">Gráfico de actividad</p>
            </div>
          </div>
        </main>
      </div>
    </div>
  )
}
