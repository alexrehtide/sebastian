<template>
  <q-layout>
    <q-page-container>
      <q-page> </q-page>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import api from "api";
import { onMounted } from "vue";
import { ref, reactive } from "vue";
import qrcode from "qrcode";

onMounted(() => {
  return api("/api/auth/authorize").then(console.log)
})

// onMounted(() => {
//   api("/api/totp/generate").then((data) => {
//     qrcode.toDataURL(
//       data.url,
//       {
//         errorCorrectionLevel: "H",
//         type: "image/png",
//         margin: 1,
//       },
//       (err, url) => {
//         if (err) {
//           return;
//         }

//         totpUrl.value = url;
//       }
//     );
//   });
// });
</script>

<style scoped lang="scss"></style>
