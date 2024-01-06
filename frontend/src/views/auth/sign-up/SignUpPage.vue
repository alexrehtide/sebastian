<template>
  <q-page class="s-page">
    <auth-card
      title="Sign Up"
      name="sign-up"
      :links="[{ path: '/auth/sign_in', title: 'Sign In' }]"
    >
      <div class="s-sign-up__block">
        <q-input model-value="" label="Email" outlined dense />
      
        <q-input model-value="" label="Username" outlined dense />

        <q-input model-value="" label="Password" type="password" outlined dense />

        <q-input model-value="" label="Password (repeat)" type="password" outlined dense />
      </div>

      <q-btn color="primary" no-caps unelevated class="s-sign-up__button">Sign Up</q-btn>
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

.s-sign-up__block {
  display: flex;
  flex-direction: column;
  gap: 12px;
}
</style>
