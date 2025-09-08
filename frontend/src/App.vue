<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
import { useAuthStore } from './stores/auth'
import './api/axios'

export default {
  name: 'App',
  setup() {
    const authStore = useAuthStore()
    
    authStore.initializeAuth()
    
    return {}
  }
}
</script>