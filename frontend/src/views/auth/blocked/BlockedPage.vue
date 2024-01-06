<template>
  <q-page class="s-page">
    <auth-card
      title="Your account is blocked"
      name="blocked"
      :links="[]"
    >
      <q-btn color="primary" no-caps unelevated class="s-blocked__button">Log Out</q-btn>
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
</style>
