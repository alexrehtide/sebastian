<template>
  <q-page class="s-page">
    <auth-card
      title="Sign Up"
      name="sign-up"
      :links="[{ path: '/auth/sign_in', title: 'Sign In' }]"
      @submit="beginRegistration"
    >
      <div class="s-sign-up__block">
        <q-input v-model="email" label="Email" outlined dense />

        <q-input v-model="username" label="Username" outlined dense />

        <q-input v-model="password" label="Password" type="password" outlined dense />

        <q-input v-model="passwordRepeat" label="Password (repeat)" type="password" outlined dense />
      </div>

      <q-btn color="primary" type="submit" no-caps unelevated class="s-sign-up__button">Sign Up</q-btn>
    </auth-card>
  </q-page>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { useRouter } from "vue-router";
import api from "api";
import AuthCard from "../components/AuthCard.vue";

const props = defineProps<{ verificationCode?: string }>();

const router = useRouter();

const email = ref("");
const username = ref("");
const password = ref("");
const passwordRepeat = ref("");

const beginRegistration = () => {
  return api("/api/registration/begin", {
    email: email.value,
    username: username.value,
    password: password.value,
  }).then(console.log);
};

onBeforeMount(() => {
  if (props.verificationCode) {
    return api("/api/registration/end", { verificationCode: props.verificationCode })
      .then((data) => {
        localStorage.setItem("ACCESS_TOKEN", data.accessToken);
        localStorage.setItem("REFRESH_TOKEN", data.refreshToken);
      })
      .then(() => router.replace("/"));
  }
});
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
