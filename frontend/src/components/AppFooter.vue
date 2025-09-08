<template>
  <footer class="bg-white border-t py-4">
    <div class="px-6">
      <div class="text-center text-sm text-gray-600">
        © {{ currentYear }} App Template. All rights reserved.
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'AppFooter',
  data() {
    return {
      currentYear: new Date().getFullYear()
    }
  }
}
</script>