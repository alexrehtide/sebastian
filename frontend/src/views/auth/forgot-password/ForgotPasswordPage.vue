<template>
  <q-page class="s-page">
    <auth-card
      v-if="stage === 'begin'"
      title="Forgot password"
      name="forgot-password"
      :links="[{ path: '/auth/sign_in', title: 'Sign In' }]"
      @submit="passwordResettingBegin"
    >
      <div class="s-forgot-password__block">
        <q-input v-model="email" label="Email" outlined dense />
      </div>

      <q-btn color="primary" type="submit" no-caps unelevated class="s-forgot-password__button">Send code</q-btn>
    </auth-card>

    <auth-card
      v-else-if="stage === 'end'"
      title="Enter new password"
      name="forgot-password"
      :links="[{ path: '/auth/sign_in', title: 'Sign In' }]"
      @submit="passwordResettingEnd"
    >
      <div class="s-forgot-password__block">
        <q-input v-model="newPassword" type="password" label="New password" outlined dense />

        <q-input v-model="newPasswordRepeat" type="password" label="New password (repeat)" outlined dense />
      </div>

      <q-btn color="primary" type="submit" no-caps unelevated class="s-forgot-password__button">Send code</q-btn>
    </auth-card>
  </q-page>
</template>

<script setup lang="ts">
import { ref } from "vue";
import AuthCard from "../components/AuthCard.vue";
import api from "api";
import { onBeforeMount } from "vue";

const props = defineProps<{ resettingCode?: string }>();

const email = ref("");
const newPassword = ref("");
const newPasswordRepeat = ref("");
const stage = ref<"begin" | "end">("begin");

const passwordResettingBegin = () => {
  return api("/api/password_resetting/begin", { email: email.value }).then(console.log);
};

const passwordResettingEnd = () => {
  return api("/api/password_resetting/end", {
    resettingCode: props.resettingCode ?? "",
    newPassword: newPassword.value,
  }).then(console.log);
};

onBeforeMount(() => {
  if (props.resettingCode) {
    stage.value = "end";
  }
});
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
