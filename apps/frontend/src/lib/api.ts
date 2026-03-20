import axios from 'axios';

const BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export const apiClient = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
  timeout: 10000,
});

// Request interceptor – attach auth token if present
apiClient.interceptors.request.use((config) => {
  if (typeof window !== 'undefined') {
    const token = localStorage.getItem('auth_token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
  }
  return config;
});

// Response interceptor – unwrap data
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      if (typeof window !== 'undefined') {
        localStorage.removeItem('auth_token');
      }
    }
    return Promise.reject(error);
  }
);

// Typed API helpers
export const api = {
  checkHealth: async () => {
    const { data } = await apiClient.get('/api/health');
    return data;
  },

  users: {
    list: async () => {
      const { data } = await apiClient.get('/api/users');
      return data;
    },
    getById: async (id: string) => {
      const { data } = await apiClient.get(`/api/users/${id}`);
      return data;
    },
    create: async (payload: { email: string; name: string }) => {
      const { data } = await apiClient.post('/api/users', payload);
      return data;
    },
    update: async (id: string, payload: { email?: string; name?: string }) => {
      const { data } = await apiClient.put(`/api/users/${id}`, payload);
      return data;
    },
    delete: async (id: string) => {
      const { data } = await apiClient.delete(`/api/users/${id}`);
      return data;
    },
  },
};
