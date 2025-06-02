<template>
  <div class="auth-container">
    <h2>Sign In</h2>
    <form @submit.prevent="signIn">
      <input v-model="email" type="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit">Sign In</button>
    </form>
    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
const email = ref('');
const password = ref('');
const message = ref('');
const router = useRouter();

const signIn = async () => {
  const res = await fetch('/signin', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email: email.value, password: password.value })
  });
  const data = await res.json();
  if (res.ok) {
    message.value = 'Sign in successful!';
    setTimeout(() => router.push('/'), 1000);
  } else {
    message.value = data.error || 'Sign in failed.';
  }
}
</script>

<style scoped>
.auth-container {
  max-width: 400px;
  margin: 40px auto;
  padding: 20px;
  background: #fff;
  border-radius: 8px;
}
input, button {
  display: block;
  margin: 10px 0;
  width: 100%;
}
</style>