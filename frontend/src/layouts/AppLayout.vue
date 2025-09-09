<template>
  <div class="min-h-screen bg-gradient-to-br from-neutral-50 via-primary-50 to-secondary-50 flex flex-col">
    <!-- Mobile overlay -->
    <div 
      v-if="isSidebarOpen" 
      class="fixed inset-0 bg-black/50 backdrop-blur-sm z-30 lg:hidden"
      @click="closeSidebar"
    ></div>
    
    <AppSidebar />
    <div class="flex flex-col min-h-screen lg:ml-64">
      <AppHeader />
      <main class="flex-1 p-4 sm:p-6 lg:p-8">
        <div class="max-w-7xl mx-auto">
          <router-view />
        </div>
      </main>
      <AppFooter />
    </div>
  </div>
</template>

<script>
import AppHeader from '../components/AppHeader.vue'
import AppSidebar from '../components/AppSidebar.vue'
import AppFooter from '../components/AppFooter.vue'
import { useSidebar } from '../composables/useSidebar'

export default {
  name: 'AppLayout',
  components: {
    AppHeader,
    AppSidebar,
    AppFooter
  },
  setup() {
    const { isSidebarOpen, closeSidebar } = useSidebar()
    
    return {
      isSidebarOpen,
      closeSidebar
    }
  }
}
</script>