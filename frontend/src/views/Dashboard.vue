<template>
  <div class="space-y-8">
    <!-- Welcome Header -->
    <div class="bg-gradient-to-r from-blue-600 via-blue-700 to-indigo-700 rounded-2xl p-4 sm:p-6 lg:p-8 text-white relative overflow-hidden">
      <div class="absolute inset-0 bg-black/10"></div>
      <div class="relative">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
          <div>
            <h1 class="text-2xl sm:text-3xl font-bold mb-2">Welcome back!</h1>
            <p class="text-blue-100 text-base sm:text-lg">Here's what's happening with your application today.</p>
          </div>
          <div class="hidden lg:block">
            <div class="flex items-center space-x-4">
              <div class="text-right">
                <div class="text-2xl font-bold">{{ getCurrentTime() }}</div>
                <div class="text-blue-200 text-sm">{{ getCurrentDate() }}</div>
              </div>
              <div class="w-16 h-16 bg-white/20 rounded-full flex items-center justify-center">
                <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"/>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- Decorative Elements -->
      <div class="absolute -right-16 -top-16 w-32 h-32 bg-white/10 rounded-full"></div>
      <div class="absolute -right-8 -bottom-8 w-24 h-24 bg-white/5 rounded-full"></div>
    </div>

    <!-- Quick Actions -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      <BaseTooltip text="Quickly create a new user with default permissions. This will open the registration form." position="top">
        <button 
          @click="quickAddUser"
          class="group bg-white rounded-xl p-6 shadow-lg hover:shadow-xl transition-all duration-300 border border-gray-100 hover:border-blue-200 w-full"
        >
          <div class="flex items-center justify-between mb-4">
            <div class="w-12 h-12 bg-blue-100 group-hover:bg-blue-200 rounded-lg flex items-center justify-center transition-colors">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
              </svg>
            </div>
            <svg class="w-5 h-5 text-gray-400 group-hover:text-blue-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
            </svg>
          </div>
          <h3 class="font-semibold text-gray-900 mb-2">Add New User</h3>
          <p class="text-sm text-gray-600">Create a new user account</p>
        </button>
      </BaseTooltip>

      <BaseTooltip text="Access the user management page to view, edit, and manage all user accounts in the system." position="top">
        <router-link 
          to="/users"
          class="group bg-white rounded-xl p-6 shadow-lg hover:shadow-xl transition-all duration-300 border border-gray-100 hover:border-emerald-200 block"
        >
          <div class="flex items-center justify-between mb-4">
            <div class="w-12 h-12 bg-emerald-100 group-hover:bg-emerald-200 rounded-lg flex items-center justify-center transition-colors">
              <svg class="w-6 h-6 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.25 2.25 0 11-4.5 0 2.25 2.25 0 014.5 0z"/>
              </svg>
            </div>
            <svg class="w-5 h-5 text-gray-400 group-hover:text-emerald-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
            </svg>
          </div>
          <h3 class="font-semibold text-gray-900 mb-2">Manage Users</h3>
          <p class="text-sm text-gray-600">View and edit user accounts</p>
        </router-link>
      </BaseTooltip>

      <button 
        @click="exportData"
        class="group bg-white rounded-xl p-6 shadow-lg hover:shadow-xl transition-all duration-300 border border-gray-100 hover:border-purple-200"
      >
        <div class="flex items-center justify-between mb-4">
          <div class="w-12 h-12 bg-purple-100 group-hover:bg-purple-200 rounded-lg flex items-center justify-center transition-colors">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
            </svg>
          </div>
          <svg class="w-5 h-5 text-gray-400 group-hover:text-purple-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
          </svg>
        </div>
        <h3 class="font-semibold text-gray-900 mb-2">Export Data</h3>
        <p class="text-sm text-gray-600">Download user reports</p>
      </button>

      <button 
        @click="viewLogs"
        class="group bg-white rounded-xl p-6 shadow-lg hover:shadow-xl transition-all duration-300 border border-gray-100 hover:border-orange-200"
      >
        <div class="flex items-center justify-between mb-4">
          <div class="w-12 h-12 bg-orange-100 group-hover:bg-orange-200 rounded-lg flex items-center justify-center transition-colors">
            <svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
            </svg>
          </div>
          <svg class="w-5 h-5 text-gray-400 group-hover:text-orange-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
          </svg>
        </div>
        <h3 class="font-semibold text-gray-900 mb-2">System Logs</h3>
        <p class="text-sm text-gray-600">View activity logs</p>
      </button>
    </div>

    <!-- Statistics Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 md:gap-6">
      <!-- Total Users Card -->
      <BaseTransition name="scale" appear>
        <div class="bg-white rounded-2xl p-6 shadow-lg border border-gray-100 hover:shadow-xl transition-shadow duration-300">
          <div class="flex items-center justify-between mb-4">
            <div class="w-12 h-12 bg-blue-100 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.25 2.25 0 11-4.5 0 2.25 2.25 0 014.5 0z"/>
              </svg>
            </div>
            <div class="flex items-center space-x-1 text-green-600 text-sm font-medium">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18"/>
              </svg>
              <span>+12%</span>
            </div>
          </div>
          <div class="space-y-2">
            <BaseTransition name="fade">
              <h3 v-if="!loading" class="text-2xl font-bold text-gray-900">{{ userCount }}</h3>
              <BaseLoading v-else type="skeleton" size="small" />
            </BaseTransition>
            <p class="text-sm font-medium text-gray-600">Total Users</p>
            <p class="text-xs text-gray-500">{{ userCount > 0 ? '+2 new this week' : 'No users yet' }}</p>
          </div>
        </div>
      </BaseTransition>

      <!-- System Status Card -->
      <div class="bg-white rounded-2xl p-6 shadow-lg border border-gray-100">
        <div class="flex items-center justify-between mb-4">
          <div class="w-12 h-12 bg-green-100 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
          </div>
          <div class="flex items-center space-x-1">
            <div class="w-2 h-2 bg-green-400 rounded-full animate-pulse"></div>
            <span class="text-xs text-gray-500">Live</span>
          </div>
        </div>
        <div class="space-y-2">
          <h3 class="text-2xl font-bold text-green-600">Online</h3>
          <p class="text-sm font-medium text-gray-600">System Status</p>
          <p class="text-xs text-gray-500">All services operational</p>
        </div>
      </div>

      <!-- Performance Card -->
      <div class="bg-white rounded-2xl p-6 shadow-lg border border-gray-100">
        <div class="flex items-center justify-between mb-4">
          <div class="w-12 h-12 bg-purple-100 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
            </svg>
          </div>
          <div class="flex items-center space-x-1 text-purple-600 text-sm font-medium">
            <span>99.9%</span>
          </div>
        </div>
        <div class="space-y-2">
          <h3 class="text-2xl font-bold text-gray-900">Fast</h3>
          <p class="text-sm font-medium text-gray-600">Performance</p>
          <p class="text-xs text-gray-500">Average response: 120ms</p>
        </div>
      </div>

      <!-- Activity Card -->
      <div class="bg-white rounded-2xl p-6 shadow-lg border border-gray-100">
        <div class="flex items-center justify-between mb-4">
          <div class="w-12 h-12 bg-orange-100 rounded-xl flex items-center justify-center">
            <svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"/>
            </svg>
          </div>
          <div class="flex items-center space-x-1 text-orange-600 text-sm font-medium">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18"/>
            </svg>
            <span>+5%</span>
          </div>
        </div>
        <div class="space-y-2">
          <h3 class="text-2xl font-bold text-gray-900">{{ activityCount }}</h3>
          <p class="text-sm font-medium text-gray-600">Recent Activity</p>
          <p class="text-xs text-gray-500">Last hour</p>
        </div>
      </div>
    </div>

    <!-- Help Section -->
    <div class="bg-gradient-to-r from-gray-50 to-gray-100 rounded-2xl p-8">
      <div class="max-w-3xl">
        <div class="flex items-start space-x-4">
          <div class="flex-shrink-0">
            <div class="w-12 h-12 bg-blue-100 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
          </div>
          <div class="flex-1">
            <h3 class="text-lg font-semibold text-gray-900 mb-2">Need help getting started?</h3>
            <p class="text-gray-600 mb-4">
              This dashboard provides an overview of your application. Use the navigation on the left to manage users, 
              view reports, and configure settings. All data is updated in real-time.
            </p>
            <div class="flex flex-wrap gap-3">
              <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors text-sm font-medium">
                View Documentation
              </button>
              <button class="px-4 py-2 bg-white text-gray-700 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors text-sm font-medium">
                Contact Support
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import apiService from '../services/api'
import BaseTooltip from '../components/BaseTooltip.vue'
import BaseLoading from '../components/BaseLoading.vue'
import BaseTransition from '../components/BaseTransition.vue'

export default {
  name: 'Dashboard',
  components: {
    BaseTooltip,
    BaseLoading,
    BaseTransition
  },
  setup() {
    const router = useRouter()
    const userCount = ref(0)
    const activityCount = ref(0)
    const loading = ref(true)
    
    const loadUserCount = async () => {
      try {
        loading.value = true
        const result = await apiService.getUsers()
        
        if (result.success) {
          userCount.value = result.data.length
        } else {
          console.error('Failed to load user count:', result.error)
        }
      } catch (error) {
        console.error('Failed to load user count:', error)
      } finally {
        loading.value = false
      }
    }
    
    const getCurrentTime = () => {
      return new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
    }
    
    const getCurrentDate = () => {
      return new Date().toLocaleDateString([], { 
        weekday: 'long', 
        year: 'numeric', 
        month: 'long', 
        day: 'numeric' 
      })
    }
    
    const quickAddUser = () => {
      // TODO: Implement quick add user modal
      alert('Quick Add User feature coming soon!')
    }
    
    const exportData = () => {
      // TODO: Implement data export
      alert('Export functionality coming soon!')
    }
    
    const viewLogs = () => {
      // TODO: Implement logs view
      alert('System logs feature coming soon!')
    }
    
    // Simulate some activity data
    const loadActivityData = () => {
      activityCount.value = Math.floor(Math.random() * 50) + 10
    }
    
    onMounted(() => {
      loadUserCount()
      loadActivityData()
      
      // Update time every minute
      setInterval(() => {
        // Force reactivity update
      }, 60000)
    })
    
    return {
      userCount,
      activityCount,
      loading,
      getCurrentTime,
      getCurrentDate,
      quickAddUser,
      exportData,
      viewLogs
    }
  }
}
</script>