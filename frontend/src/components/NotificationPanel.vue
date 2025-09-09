<template>
  <div>
    <!-- Mobile full-screen overlay -->
    <div 
      v-if="isOpen && isMobile" 
      class="fixed inset-0 z-[60] bg-white flex flex-col h-screen w-screen"
    >
      <!-- Mobile header - matches AppHeader mobile height -->
      <div class="flex items-center justify-between px-3 py-3 border-b border-neutral-200 bg-white">
        <div class="w-10"></div> <!-- Spacer for centering -->
        <h2 class="text-lg font-semibold text-neutral-900">Notifications</h2>
        <button 
          @click="closeNotifications" 
          class="p-1.5 rounded-lg hover:bg-neutral-100 transition-colors"
        >
          <svg class="w-6 h-6 text-neutral-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
          </svg>
        </button>
      </div>
      
      <!-- Mobile notification content -->
      <div class="flex-1 overflow-y-auto p-4 bg-white min-h-0">
        <NotificationList />
      </div>
    </div>

    <!-- Desktop right sidebar -->
    <div 
      v-if="isOpen && !isMobile" 
      class="fixed right-0 top-0 h-screen w-80 bg-white shadow-xl border-l border-neutral-200 z-[40] transform transition-transform duration-300 flex flex-col"
      :class="isOpen ? 'translate-x-0' : 'translate-x-full'"
    >
      <!-- Desktop header - matches AppHeader desktop height -->
      <div class="flex items-center justify-between px-6 py-4 border-b border-neutral-200 bg-white">
        <div class="w-8"></div> <!-- Spacer for centering -->
        <h2 class="text-lg font-semibold text-neutral-900">Notifications</h2>
        <button 
          @click="closeNotifications" 
          class="p-2 rounded-lg hover:bg-neutral-100 transition-colors"
        >
          <svg class="w-5 h-5 text-neutral-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
          </svg>
        </button>
      </div>
      
      <!-- Desktop notification content -->
      <div class="flex-1 overflow-y-auto p-4 bg-white min-h-0">
        <NotificationList />
      </div>
    </div>

    <!-- Desktop overlay backdrop -->
    <div 
      v-if="isOpen && !isMobile" 
      class="fixed inset-0"
      @click="closeNotifications"
    ></div>
  </div>
</template>

<script>
import { computed, onMounted, onUnmounted } from 'vue'
import NotificationList from './NotificationList.vue'

export default {
  name: 'NotificationPanel',
  components: {
    NotificationList
  },
  props: {
    isOpen: {
      type: Boolean,
      default: false
    }
  },
  emits: ['close'],
  setup(props, { emit }) {
    const isMobile = computed(() => {
      return window.innerWidth < 768
    })

    const closeNotifications = () => {
      emit('close')
    }

    // Handle escape key
    const handleEscape = (event) => {
      if (event.key === 'Escape' && props.isOpen) {
        closeNotifications()
      }
    }

    onMounted(() => {
      document.addEventListener('keydown', handleEscape)
    })

    onUnmounted(() => {
      document.removeEventListener('keydown', handleEscape)
    })

    return {
      isMobile,
      closeNotifications
    }
  }
}
</script>