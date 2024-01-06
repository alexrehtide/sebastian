<template>
  <q-page class="s-page">
    <auth-card
      title="Sign In"
      name="sign-in"
      :links="[
        { path: '/auth/sign_up', title: 'Sign Up' },
        { path: '/auth/forgot_password', title: 'Forgot Password' },
      ]"
      @submit="signIn"
    >
      <div class="s-sign-in__block">
        <q-input v-model="email" label="Email / Username" outlined dense />

        <q-input v-model="password" label="Password" type="password" outlined dense />
      </div>

      <q-btn color="primary" type="submit" no-caps unelevated class="s-sign-in__button">Sign In</q-btn>

      <div class="s-sign-in__platforms">
        <q-btn icon="fa-brands fa-google" round flat @click="signInWith('google')"></q-btn>
        <q-btn icon="fa-brands fa-yandex-international" round flat></q-btn>
        <q-btn icon="fa-brands fa-twitch" round flat></q-btn>
      </div>
    </auth-card>
  </q-page>
</template>

<script setup lang="ts">
import { ref, onBeforeMount } from "vue";
import { useRouter } from "vue-router";
import api from "api";
import AuthCard from "../components/AuthCard.vue";

const props = defineProps<{ exchange?: { platform: string; code: string; state: string } }>();

const router = useRouter();

const email = ref("");
const password = ref("");

const signInWith = (platform: string) => {
  return api("/api/oauth2/auth_code_url", { platform }).then(({ url }) => {
    window.location.href = url;
  });
};

const signIn = () => {
  return api("/api/auth/authenticate", { email: email.value, password: password.value })
    .then(onAuthenticate)
    .then(() => router.push("/"));
};

const onAuthenticate = (data: { accessToken: string; refreshToken: string }) => {
  localStorage.setItem("ACCESS_TOKEN", data.accessToken);
  localStorage.setItem("REFRESH_TOKEN", data.refreshToken);
};

onBeforeMount(() => {
  if (props.exchange) {
    return api("/api/oauth2/authenticate", props.exchange)
      .then(onAuthenticate)
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

.s-sign-in__block {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.s-sign-in__platforms {
  display: flex;
  flex-direction: column;
  gap: 12px;
  position: absolute;
  left: 12px;
  bottom: 12px;
}
</style>
