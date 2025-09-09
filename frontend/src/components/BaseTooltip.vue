<template>
  <div class="relative inline-block" @mouseenter="showTooltip" @mouseleave="hideTooltip">
    <slot></slot>
    <div 
      v-if="isVisible" 
      :class="tooltipClasses"
      class="absolute z-50 px-3 py-2 text-xs font-medium text-white bg-neutral-900 rounded-lg shadow-lg opacity-0 transition-opacity duration-200"
      :style="tooltipStyle"
      ref="tooltip"
    >
      {{ text }}
      <div class="tooltip-arrow absolute w-2 h-2 bg-neutral-900 transform rotate-45" :class="arrowClasses"></div>
    </div>
  </div>
</template>

<script>
import { ref, computed, nextTick } from 'vue'

export default {
  name: 'BaseTooltip',
  props: {
    text: {
      type: String,
      required: true
    },
    position: {
      type: String,
      default: 'top',
      validator: (value) => ['top', 'bottom', 'left', 'right'].includes(value)
    },
    delay: {
      type: Number,
      default: 500
    }
  },
  setup(props) {
    const isVisible = ref(false)
    let showTimeout = null
    let hideTimeout = null

    const tooltipClasses = computed(() => {
      const positions = {
        top: 'bottom-full left-1/2 transform -translate-x-1/2 mb-2',
        bottom: 'top-full left-1/2 transform -translate-x-1/2 mt-2',
        left: 'right-full top-1/2 transform -translate-y-1/2 mr-2',
        right: 'left-full top-1/2 transform -translate-y-1/2 ml-2'
      }
      return positions[props.position]
    })

    const arrowClasses = computed(() => {
      const arrows = {
        top: 'top-full left-1/2 transform -translate-x-1/2 -mt-1',
        bottom: 'bottom-full left-1/2 transform -translate-x-1/2 -mb-1',
        left: 'left-full top-1/2 transform -translate-y-1/2 -ml-1',
        right: 'right-full top-1/2 transform -translate-y-1/2 -mr-1'
      }
      return arrows[props.position]
    })

    const tooltipStyle = computed(() => {
      return isVisible.value ? { opacity: '1' } : { opacity: '0' }
    })

    const showTooltip = () => {
      if (hideTimeout) {
        clearTimeout(hideTimeout)
        hideTimeout = null
      }
      
      showTimeout = setTimeout(() => {
        isVisible.value = true
        nextTick(() => {
          const tooltip = document.querySelector('.tooltip-arrow').parentElement
          if (tooltip) {
            tooltip.style.opacity = '1'
          }
        })
      }, props.delay)
    }

    const hideTooltip = () => {
      if (showTimeout) {
        clearTimeout(showTimeout)
        showTimeout = null
      }
      
      hideTimeout = setTimeout(() => {
        isVisible.value = false
      }, 100)
    }

    return {
      isVisible,
      tooltipClasses,
      arrowClasses,
      tooltipStyle,
      showTooltip,
      hideTooltip
    }
  }
}
</script>

<style scoped>
.tooltip-arrow {
  filter: drop-shadow(0 1px 2px rgba(0, 0, 0, 0.1));
}
</style>