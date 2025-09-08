<template>
  <div>
    <div class="mb-6">
      <h1 class="text-2xl font-semibold text-gray-900">Edit User</h1>
      <p class="mt-2 text-sm text-gray-600">Update user information</p>
    </div>

    <div v-if="loading" class="text-center py-4">
      Loading user...
    </div>

    <div v-else-if="error" class="text-red-600 text-center py-4">
      {{ error }}
    </div>

    <div v-else class="bg-white shadow sm:rounded-lg">
      <form @submit.prevent="updateUser" class="space-y-6 py-6 px-6">
        <div>
          <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
          <input
            v-model="form.username"
            type="text"
            id="username"
            required
            class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-3 py-2 border"
          />
        </div>

        <div>
          <label for="password" class="block text-sm font-medium text-gray-700">
            New Password (leave blank to keep current)
          </label>
          <input
            v-model="form.password"
            type="password"
            id="password"
            class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm px-3 py-2 border"
          />
        </div>

        <div v-if="updateError" class="text-red-600 text-sm">
          {{ updateError }}
        </div>

        <div class="flex justify-between">
          <router-link 
            to="/users"
            class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-medium py-2 px-4 rounded"
          >
            Cancel
          </router-link>
          <button
            type="submit"
            :disabled="updating"
            class="bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2 px-4 rounded disabled:opacity-50"
          >
            {{ updating ? 'Updating...' : 'Update User' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'

export default {
  name: 'UserEdit',
  setup() {
    const route = useRoute()
    const router = useRouter()
    
    const user = ref(null)
    const form = ref({
      username: '',
      password: ''
    })
    
    const loading = ref(true)
    const error = ref('')
    const updating = ref(false)
    const updateError = ref('')
    
    const loadUser = async () => {
      try {
        const userId = route.params.id
        const response = await axios.get(`/api/users/${userId}`)
        user.value = response.data
        form.value.username = response.data.username
      } catch (err) {
        error.value = 'Failed to load user'
        console.error('Error loading user:', err)
      } finally {
        loading.value = false
      }
    }
    
    const updateUser = async () => {
      updating.value = true
      updateError.value = ''
      
      try {
        const userId = route.params.id
        const updateData = { username: form.value.username }
        
        if (form.value.password) {
          updateData.password = form.value.password
        }
        
        await axios.put(`/api/users/${userId}`, updateData)
        router.push('/users')
      } catch (err) {
        updateError.value = err.response?.data?.error || 'Failed to update user'
        console.error('Error updating user:', err)
      } finally {
        updating.value = false
      }
    }
    
    onMounted(() => {
      loadUser()
    })
    
    return {
      user,
      form,
      loading,
      error,
      updating,
      updateError,
      updateUser
    }
  }
}
</script>