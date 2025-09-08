<template>
  <header class="bg-white shadow-sm border-b">
    <div class="px-6 py-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <h1 class="text-xl font-semibold text-gray-800">App Template</h1>
        </div>
        <div class="flex items-center space-x-4">
          <span class="text-sm text-gray-600">Welcome, {{ user?.username }}</span>
          <button 
            @click="logout" 
            class="px-4 py-2 text-sm font-medium text-white bg-red-600 rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500"
          >
            Logout
          </button>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

export default {
  name: 'AppHeader',
  setup() {
    const router = useRouter()
    const authStore = useAuthStore()
    
    const user = computed(() => authStore.user)
    
    const logout = () => {
      authStore.logout()
      router.push('/login')
    }
    
    return {
      user,
      logout
    }
  }
}
</script>