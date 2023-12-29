<template>
  <q-layout>
    <q-page-container>
      <q-page>
        <q-input v-model="email" label="Email" />
        <q-input v-model="password" type="password" label="Password" />
        <q-btn @click="authenticate">Войти</q-btn>
        <q-btn @click="authorize">Пользователь</q-btn>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { onMounted } from "vue";
import api from "api";
import { ref } from "vue";
import { reactive } from "vue";

const email = ref("");
const password = ref("");

const user = reactive({
  email: "",
  id: 0,
});

const authenticate = () => {
  return api("/api/auth/authenticate", { email: email.value, password: password.value }).then((data) => {
    localStorage.setItem("ACCESS_TOKEN", data.accessToken);
    localStorage.setItem("REFRESH_TOKEN", data.refreshToken);
  });
};

const authorize = () => {
  return api("/api/auth/authorize").then(console.log);
};
</script>

<style scoped lang="scss"></style>
