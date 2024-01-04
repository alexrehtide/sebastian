<template>
  <q-layout>
    <q-page-container>
      <q-page>
        <q-input v-model="email" label="Email" />
        <q-input v-model="password" type="password" label="Password" />
        <q-btn @click="authenticate">Войти</q-btn>
        <q-btn @click="authorize">Пользователь</q-btn>

        <q-img v-if="totpUrl" :src="totpUrl" width="200px" height="200px"></q-img>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import api from "api";
import { onMounted } from "vue";
import { ref, reactive } from "vue";
import qrcode from "qrcode";

const email = ref("");
const password = ref("");
const totpUrl = ref<string>("");

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
  return api("/api/auth/authorize");
};

onMounted(() => {
  api("/api/totp/generate").then((data) => {
    qrcode.toDataURL(
      data.url,
      {
        errorCorrectionLevel: "H",
        type: "image/png",
        margin: 1,
      },
      (err, url) => {
        if (err) {
          return;
        }

        totpUrl.value = url;
      }
    );
  });
});
</script>

<style scoped lang="scss"></style>
