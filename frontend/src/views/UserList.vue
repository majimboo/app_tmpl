<template>
  <div>
    <div class="mb-6">
      <h1 class="text-2xl font-semibold text-gray-900">Users</h1>
      <p class="mt-2 text-sm text-gray-600">Manage system users</p>
    </div>

    <div class="bg-white shadow overflow-hidden sm:rounded-md">
      <div v-if="loading" class="p-6 text-center">
        <div class="text-gray-500">Loading users...</div>
      </div>
      
      <div v-else-if="error" class="p-6 text-center text-red-600">
        {{ error }}
      </div>
      
      <ul v-else-if="users.length > 0" role="list" class="divide-y divide-gray-200">
        <li v-for="user in users" :key="user.id" class="px-6 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="h-10 w-10 rounded-full bg-gray-300 flex items-center justify-center">
                  <svg class="h-6 w-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                  </svg>
                </div>
              </div>
              <div class="ml-4">
                <div class="text-sm font-medium text-gray-900">{{ user.username }}</div>
                <div class="text-sm text-gray-500">ID: {{ user.id }}</div>
              </div>
            </div>
            <div class="flex items-center space-x-2">
              <router-link 
                :to="`/users/${user.id}/edit`"
                class="inline-flex items-center px-3 py-1.5 border border-gray-300 text-xs font-medium rounded text-gray-700 bg-white hover:bg-gray-50"
              >
                Edit
              </router-link>
              <button 
                @click="deleteUser(user.id)"
                class="inline-flex items-center px-3 py-1.5 border border-red-300 text-xs font-medium rounded text-red-700 bg-red-50 hover:bg-red-100"
              >
                Delete
              </button>
            </div>
          </div>
        </li>
      </ul>
      
      <div v-else class="p-6 text-center text-gray-500">
        No users found
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'

export default {
  name: 'UserList',
  setup() {
    const users = ref([])
    const loading = ref(true)
    const error = ref('')
    
    const loadUsers = async () => {
      try {
        loading.value = true
        const response = await axios.get('/api/users')
        users.value = response.data
      } catch (err) {
        error.value = 'Failed to load users'
        console.error('Error loading users:', err)
      } finally {
        loading.value = false
      }
    }
    
    const deleteUser = async (userId) => {
      if (!confirm('Are you sure you want to delete this user?')) {
        return
      }
      
      try {
        await axios.delete(`/api/users/${userId}`)
        users.value = users.value.filter(user => user.id !== userId)
      } catch (err) {
        alert('Failed to delete user')
        console.error('Error deleting user:', err)
      }
    }
    
    onMounted(() => {
      loadUsers()
    })
    
    return {
      users,
      loading,
      error,
      deleteUser
    }
  }
}
</script>