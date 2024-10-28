<template>
  <div class="signup-page">
    <SignupForm />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import SignupForm from '@/components/SignupForm.vue'

export default defineComponent({
  name: 'SignupView',
  components: {
    SignupForm,
  },
})
</script>

<style scoped>
.signup-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
</style>
