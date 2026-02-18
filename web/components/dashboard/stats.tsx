'use client'

import { useQuery } from '@tanstack/react-query'
import { api } from '@/lib/api'
import { Server, Plug, CircleDot, Activity } from 'lucide-react'

interface DashboardStats {
  summary: {
    total_racks: number
    total_panels: number
    total_nodes: number
    active_nodes: number
    inactive_nodes: number
    maintenance_nodes: number
  }
}

export function DashboardStats() {
  const { data, isLoading } = useQuery<DashboardStats>({
    queryKey: ['dashboard-stats'],
    queryFn: async () => {
      // TODO: Implementar llamada real
      return {
        summary: {
          total_racks: 8,
          total_panels: 32,
          total_nodes: 768,
          active_nodes: 654,
          inactive_nodes: 45,
          maintenance_nodes: 69,
        },
      }
    },
  })

  if (isLoading) {
    return <div>Cargando estadísticas...</div>
  }

  const stats = [
    {
      name: 'Racks',
      value: data?.summary.total_racks || 0,
      icon: Server,
      color: 'bg-blue-500',
    },
    {
      name: 'Patch Panels',
      value: data?.summary.total_panels || 0,
      icon: Plug,
      color: 'bg-green-500',
    },
    {
      name: 'Nodos Totales',
      value: data?.summary.total_nodes || 0,
      icon: CircleDot,
      color: 'bg-purple-500',
    },
    {
      name: 'Nodos Activos',
      value: data?.summary.active_nodes || 0,
      icon: Activity,
      color: 'bg-emerald-500',
    },
  ]

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      {stats.map((stat) => (
        <div key={stat.name} className="bg-white rounded-lg shadow p-6">
          <div className="flex items-center">
            <div className={`p-3 rounded-lg ${stat.color}`}>
              <stat.icon className="w-6 h-6 text-white" />
            </div>
            <div className="ml-4">
              <p className="text-sm font-medium text-gray-600">{stat.name}</p>
              <p className="text-2xl font-bold text-gray-900">{stat.value}</p>
            </div>
          </div>
        </div>
      ))}
    </div>
  )
}
