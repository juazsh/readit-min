<template>
  <div style="display: flex; flex-direction: row; gap: 2rem;  margin: auto;">
    <div>
      <h2 class="section-title">Create a Post</h2>
      <form @submit.prevent="submitPost" class="post-form">
        <input v-model="title" placeholder="Title" required />
        <textarea v-model="content" placeholder="What's on your mind?" required></textarea>
        <button type="submit">Post</button>
      </form>
    </div>

    <div>
      <h2 class="section-title">Latest Posts</h2>
      <div v-if="posts.length === 0">No posts yet.</div>
      <div v-for="post in posts" :key="post.ID" class="post-card">
        <h3>{{ post.title }}</h3>
        <p>{{ post.content }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const posts = ref([])
const title = ref('')
const content = ref('')

const fetchPosts = async () => {
  const res = await axios.get('/posts')
  posts.value = res.data
}

const submitPost = async () => {
  await axios.post('/posts', { title: title.value, content: content.value })
  title.value = ''
  content.value = ''
  fetchPosts()
}

onMounted(fetchPosts)
</script>

<style scoped>
.section-title {
  margin-bottom: 1rem;
  font-size: 1.5rem;
  color: #333;
}

.post-form {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  margin-bottom: 2rem;
}

.post-form input,
.post-form textarea {
  padding: 0.75rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  font-size: 1rem;
}

.post-form button {
  align-self: flex-start;
  padding: 0.5rem 1.5rem;
  background-color: #5b6dfd;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: bold;
}

.post-form button:hover {
  background-color: #4a5ae0;
}

.post-card {
  padding: 1rem;
  color: #332312;
  border-radius: 10px;
  margin-bottom: 1rem;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.03);
}

.post-card h3 {
  margin: 0 0 0.5rem 0;
}

.post-card p {
  margin: 0;
  color: #555;
}
</style>
