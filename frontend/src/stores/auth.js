import { defineStore } from 'pinia'
import axios from 'axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: localStorage.getItem('token') || null,
    loading: false
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    isLoading: (state) => state.loading
  },

  actions: {
    async login(credentials) {
      this.loading = true
      try {
        const response = await axios.post('/api/auth/login', credentials)
        this.token = response.data.token
        this.user = response.data.user
        localStorage.setItem('token', this.token)
        axios.defaults.headers.common['Authorization'] = `Bearer ${this.token}`
        return { success: true }
      } catch (error) {
        return { 
          success: false, 
          error: error.response?.data?.error || 'Login failed' 
        }
      } finally {
        this.loading = false
      }
    },

    async register(userData) {
      this.loading = true
      try {
        const response = await axios.post('/api/auth/register', userData)
        this.token = response.data.token
        this.user = response.data.user
        localStorage.setItem('token', this.token)
        axios.defaults.headers.common['Authorization'] = `Bearer ${this.token}`
        return { success: true }
      } catch (error) {
        return { 
          success: false, 
          error: error.response?.data?.error || 'Registration failed' 
        }
      } finally {
        this.loading = false
      }
    },

    logout() {
      this.user = null
      this.token = null
      localStorage.removeItem('token')
      delete axios.defaults.headers.common['Authorization']
    },

    initializeAuth() {
      if (this.token) {
        axios.defaults.headers.common['Authorization'] = `Bearer ${this.token}`
      }
    }
  }
})