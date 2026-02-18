'use client'

import { useQuery } from '@tanstack/react-query'
import { Clock, User } from 'lucide-react'

interface RecentScan {
  id: string
  node_rfid: string
  user_name: string
  scanned_at: string
}

export function RecentScans() {
  const { data, isLoading } = useQuery({
    queryKey: ['recent-scans'],
    queryFn: async () => {
      // TODO: Implementar llamada real
      return [
        {
          id: '1',
          node_rfid: 'E200341502001080',
          user_name: 'Juan Pérez',
          scanned_at: '2024-01-20T10:30:00Z',
        },
        {
          id: '2',
          node_rfid: 'E200341502001081',
          user_name: 'María García',
          scanned_at: '2024-01-20T09:15:00Z',
        },
      ] as RecentScan[]
    },
  })

  if (isLoading) {
    return <div>Cargando escaneos recientes...</div>
  }

  return (
    <div className="bg-white rounded-lg shadow p-6">
      <h2 className="text-lg font-semibold mb-4">Escaneos Recientes</h2>
      <div className="space-y-4">
        {data?.map((scan) => (
          <div
            key={scan.id}
            className="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
          >
            <div className="flex items-center space-x-3">
              <div className="p-2 bg-primary-100 rounded-full">
                <User className="w-4 h-4 text-primary-600" />
              </div>
              <div>
                <p className="font-medium text-gray-900">{scan.node_rfid}</p>
                <p className="text-sm text-gray-500">{scan.user_name}</p>
              </div>
            </div>
            <div className="flex items-center text-sm text-gray-500">
              <Clock className="w-4 h-4 mr-1" />
              {new Date(scan.scanned_at).toLocaleTimeString()}
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}
