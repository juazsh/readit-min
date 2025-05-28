<template>
  <div class="app-container">
    <div class="content-box">
      <h1>Reddit Clone</h1>
      <PostList />
    </div>
  </div>
</template>

<script setup>
import PostList from './components/PostList.vue'
</script>

<style>
body {
  margin: 0;
  font-family: 'Segoe UI', sans-serif;
  background-color: #f8f9fc;
}

.app-container {
  display: flex;
  justify-content: center;
  padding: 2rem;
}

.content-box {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  max-width: 800px;
  width: 100%;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.05);
}
</style>
