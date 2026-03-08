import axios, { AxiosInstance, AxiosError } from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1';

interface ApiResponse<T> {
  data: T;
  status: number;
}

interface AuthResponse {
  access_token: string;
  refresh_token: string;
  user: {
    id: string;
    email: string;
    full_name: string;
    status: string;
  };
  tenant: {
    id: string;
    name: string;
    plan: string;
  };
}

interface RegisterRequest {
  company_name: string;
  full_name: string;
  email: string;
  password: string;
}

interface LoginRequest {
  email: string;
  password: string;
}

class ApiClient {
  private client: AxiosInstance;

  constructor() {
    this.client = axios.create({
      baseURL: API_BASE_URL,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Interceptor para agregar token JWT
    this.client.interceptors.request.use((config) => {
      const token = localStorage.getItem('accessToken');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    });

    // Interceptor para renovar token si expira
    this.client.interceptors.response.use(
      (response) => response,
      async (error: AxiosError) => {
        const originalRequest = error.config as any;

        if (error.response?.status === 401 && !originalRequest._retry) {
          originalRequest._retry = true;

          try {
            const refreshToken = localStorage.getItem('refreshToken');
            if (!refreshToken) {
              throw new Error('No refresh token');
            }

            const response = await axios.post(`${API_BASE_URL}/auth/refresh`, {
              refresh_token: refreshToken,
            });

            const { access_token } = response.data;
            localStorage.setItem('accessToken', access_token);

            // Reintentar request original
            originalRequest.headers.Authorization = `Bearer ${access_token}`;
            return this.client(originalRequest);
          } catch (err) {
            // Limpiar tokens y redirigir a login
            localStorage.removeItem('accessToken');
            localStorage.removeItem('refreshToken');
            localStorage.removeItem('userData');
            localStorage.removeItem('tenantData');
            window.location.href = '/login';
            throw err;
          }
        }

        return Promise.reject(error);
      }
    );
  }

  // ============================================================================
  // AUTENTICACIÓN
  // ============================================================================

  async register(data: RegisterRequest): Promise<AuthResponse> {
    const response = await this.client.post<AuthResponse>('/auth/register', data);
    return response.data;
  }

  async login(data: LoginRequest): Promise<AuthResponse> {
    const response = await this.client.post<AuthResponse>('/auth/login', data);
    return response.data;
  }

  async refresh(refreshToken: string): Promise<{ access_token: string }> {
    const response = await this.client.post<{ access_token: string }>('/auth/refresh', {
      refresh_token: refreshToken,
    });
    return response.data;
  }

  async logout(): Promise<void> {
    await this.client.post('/auth/logout');
  }

  // ============================================================================
  // USUARIOS
  // ============================================================================

  async getUser(userId: string) {
    const response = await this.client.get(`/users/${userId}`);
    return response.data;
  }

  async listUsers(page = 1, limit = 20) {
    const response = await this.client.get('/users', {
      params: { page, limit },
    });
    return response.data;
  }

  async updateUser(userId: string, data: any) {
    const response = await this.client.put(`/users/${userId}`, data);
    return response.data;
  }

  async deleteUser(userId: string) {
    await this.client.delete(`/users/${userId}`);
  }

  // ============================================================================
  // ROLES
  // ============================================================================

  async listRoles(page = 1, limit = 20) {
    const response = await this.client.get('/roles', {
      params: { page, limit },
    });
    return response.data;
  }

  async getRole(roleId: string) {
    const response = await this.client.get(`/roles/${roleId}`);
    return response.data;
  }

  async createRole(data: any) {
    const response = await this.client.post('/roles', data);
    return response.data;
  }

  async updateRole(roleId: string, data: any) {
    const response = await this.client.put(`/roles/${roleId}`, data);
    return response.data;
  }

  async deleteRole(roleId: string) {
    await this.client.delete(`/roles/${roleId}`);
  }

  // ============================================================================
  // FACTURACIÓN
  // ============================================================================

  async getSubscription() {
    const response = await this.client.get('/billing/subscription');
    return response.data;
  }

  async createCheckoutSession(planId: string) {
    const response = await this.client.post('/billing/checkout', { plan_id: planId });
    return response.data;
  }

  async listInvoices(page = 1, limit = 20) {
    const response = await this.client.get('/billing/invoices', {
      params: { page, limit },
    });
    return response.data;
  }

  // ============================================================================
  // INVENTARIO (Preparado para Fase 1)
  // ============================================================================

  async listNodes(page = 1, limit = 20) {
    const response = await this.client.get('/inventory/nodes', {
      params: { page, limit },
    });
    return response.data;
  }

  async createNode(data: any) {
    const response = await this.client.post('/inventory/nodes', data);
    return response.data;
  }

  async updateNode(nodeId: string, data: any) {
    const response = await this.client.put(`/inventory/nodes/${nodeId}`, data);
    return response.data;
  }

  async deleteNode(nodeId: string) {
    await this.client.delete(`/inventory/nodes/${nodeId}`);
  }

  // ============================================================================
  // PLANOS (Preparado para Fase 1)
  // ============================================================================

  async listPlans(page = 1, limit = 20) {
    const response = await this.client.get('/plans', {
      params: { page, limit },
    });
    return response.data;
  }

  async createPlan(data: any) {
    const response = await this.client.post('/plans', data);
    return response.data;
  }

  async updatePlan(planId: string, data: any) {
    const response = await this.client.put(`/plans/${planId}`, data);
    return response.data;
  }

  async deletePlan(planId: string) {
    await this.client.delete(`/plans/${planId}`);
  }
}

export const apiClient = new ApiClient();

// Funciones de utilidad para autenticación
export const setAuthTokens = (accessToken: string, refreshToken: string) => {
  localStorage.setItem('accessToken', accessToken);
  localStorage.setItem('refreshToken', refreshToken);
};

export const clearAuthTokens = () => {
  localStorage.removeItem('accessToken');
  localStorage.removeItem('refreshToken');
  localStorage.removeItem('userData');
  localStorage.removeItem('tenantData');
};

export const getAccessToken = () => localStorage.getItem('accessToken');
export const getRefreshToken = () => localStorage.getItem('refreshToken');
export const isAuthenticated = () => !!getAccessToken();
