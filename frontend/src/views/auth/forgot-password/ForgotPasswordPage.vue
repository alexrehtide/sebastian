<template>
  <q-page class="s-page">
    <auth-card
      title="Forgot password"
      name="forgot-password"
      :links="[{ path: '/auth/sign_in', title: 'Sign In' }]"
    >
      <div class="s-forgot-password__block">
        <q-input model-value="" label="Email" outlined dense />
      </div>

      <q-btn color="primary" no-caps unelevated class="s-forgot-password__button">Send code</q-btn>
    </auth-card>
  </q-page>
</template>

<script setup lang="ts">
import AuthCard from "../components/AuthCard.vue";
</script>

<style scoped lang="scss">
.s-page {
  display: flex;
  align-items: center;
  justify-content: center;
}

.s-forgot-password__block {
  display: flex;
  flex-direction: column;
  gap: 12px;
}
</style>
