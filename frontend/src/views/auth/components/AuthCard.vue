<template>
  <div class="s-card">
    <form class="s-card__form" @submit.prevent="emit('submit')">
      <div class="s-card__title">{{ title }}</div>
      <slot></slot>
    </form>

    <div :class="illustrationClass"></div>

    <div class="s-card__links">
      <q-btn v-for="link in links" :key="link.path" :to="link.path" no-caps flat>{{ link.title }}</q-btn>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
  title: string;
  links: Array<{ path: string; title: string }>;
  name: "sign-in" | "sign-up" | "verify-email" | "forgot-password" | "blocked" | "totp";
  onSubmit?: () => void
}>();

const emit = defineEmits<{ (e: 'submit'): void }>();

const illustrationClass = computed(() => {
  return ["s-card__illustration", `s-card__illustration--${props.name}`];
});
</script>

<style scoped lang="scss">
.s-card {
  display: flex;
  position: relative;
  width: 800px;
  height: 500px;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 0 12px 1px rgba(210, 164, 241, 0.7);
  overflow: hidden;
}

.s-card__form {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  gap: 24px;
  width: 50%;
  padding: 24px;
}

.s-card__title {
  text-align: center;
  font-size: 36px;
  font-weight: 300;
}

.s-card__illustration {
  width: 50%;
  background-size: contain;
  background-repeat: no-repeat;
  background-position: center center;
}

.s-card__illustration--sign-in {
  background-image: url(/sign-in.png);
}

.s-card__illustration--sign-up {
  background-image: url(/sign-up.png);
}

.s-card__illustration--verify-email {
  background-image: url(/verify-email.png);
}

.s-card__illustration--forgot-password {
  background-image: url(/forgot-password.png);
}

.s-card__illustration--blocked {
  background-image: url(/blocked.png);
}

.s-card__illustration--totp {
  background-image: url(/totp.png);
}

.s-card__links {
  display: flex;
  gap: 8px;
  position: absolute;
  right: 12px;
  bottom: 12px;
}
</style>
