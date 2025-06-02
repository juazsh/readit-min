<template>
  <div class="auth-container">
    <h2> Sign Up</h2>
    <form @submit.prevent="signUp">
      <input v-model="email" type="email" placeholder="Enter your email" required />
      <input v-model="password" type="password" placeholder="Enter your password" required />
      <button type="submit">Create Account</button>
    </form>
    <p v-if="message" class="error">{{ message }}</p>
  </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const email = ref('');
const password = ref('');
const message = ref('');
const router = useRouter();
const signUp = async () => {
   const res = await fetch('/signup', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email: email.value, password: password.value })
  })
  const data = await res.json()
  if (res.ok) {
    message.value = 'Signup successful. You can sign in now.'
    setTimeout(() => router.push('/signin'), 1000)
  } else {
    message.value = data.error || 'Signup failed.'
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