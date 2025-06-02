<template>
  <div class="post-container">
    <div class="post-form-section">
      <h2 class="section-title">Create a Post</h2>
      <form @submit.prevent="submitPost" class="post-form">
        <input 
          v-model="title" 
          placeholder="Title" 
          required 
          class="form-input"
        />
        <textarea 
          v-model="content" 
          placeholder="What's on your mind?" 
          required 
          class="form-textarea"
        ></textarea>
        <button type="submit" class="submit-button">Post</button>
      </form>
    </div>
    <div class="posts-section">
      <h2 class="section-title">Latest Posts</h2>
      <div v-if="posts.length === 0" class="no-posts">
        No posts yet. Be the first to share your thoughts!
      </div>
      <div class="posts-list">
        <div v-for="post in posts" :key="post.ID" class="post-card">
          <h3 class="post-title">{{ post.title }}</h3>
          <p class="post-content">{{ post.content }}</p>
          <div class="post-footer">
            <span class="post-date">Just now</span>
          </div>
        </div>
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
  try {
    const res = await axios.get('/posts')
    posts.value = res.data.reverse()
  } catch (error) {
    console.error('Error fetching posts:', error)
  }
}

const submitPost = async () => {
  try {
    await axios.post('/posts', { 
      title: title.value, 
      content: content.value 
    })
    title.value = ''
    content.value = ''
    fetchPosts()
  } catch (error) {
    console.error('Error submitting post:', error)
  }
}

onMounted(fetchPosts)
</script>

<style scoped>
.post-container {
  display: flex;
  gap: 2rem;
  width: 100%;
}

.post-form-section, .posts-section {
  flex: 1;
  min-width: 0;
}

.section-title {
  margin-bottom: 1.5rem;
  font-size: 1.5rem;
  color: #333;
  font-weight: 600;
  position: relative;
  text-align: left;
}

.section-title::after {
  content: '';
  position: absolute;
  bottom: -8px;
  left: 0;
  width: 40px;
  height: 3px;
  background-color: #5b6dfd;
  border-radius: 2px;
}

.post-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.form-input, .form-textarea {
  padding: 0.875rem;
  border: 1px solid #e1e4e8;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.2s, box-shadow 0.2s;
  background-color: #f9fafb;
  color: #333;;
}

.form-input:focus, .form-textarea:focus {
  outline: none;
  border-color: #5b6dfd;
  box-shadow: 0 0 0 3px rgba(91, 109, 253, 0.2);
  background-color: #fff;
}

.form-textarea {
  min-height: 120px;
  resize: vertical;
}

.submit-button {
  align-self: flex-start;
  padding: 0.625rem 1.5rem;
  background-color: #5b6dfd;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: background-color 0.2s, transform 0.1s;
  box-shadow: 0 2px 4px rgba(91, 109, 253, 0.2);
}

.submit-button:hover {
  background-color: #4a5ae0;
}

.submit-button:active {
  transform: translateY(1px);
}

.no-posts {
  color: #666;
  font-style: italic;
  text-align: center;
  padding: 2rem 0;
  background-color: #f9fafb;
  border-radius: 8px;
  border: 1px dashed #e1e4e8;
}

.posts-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  max-height: 500px;
  overflow-y: auto;
  padding-right: 0.5rem;
}

.posts-list::-webkit-scrollbar {
  width: 6px;
}

.posts-list::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 10px;
}

.posts-list::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 10px;
}

.posts-list::-webkit-scrollbar-thumb:hover {
  background: #a1a1a1;
}

.post-card {
  padding: 1.25rem;
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s, box-shadow 0.2s;
  border: 1px solid #eaedf0;
}

.post-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.08);
}

.post-title {
  margin: 0 0 0.75rem 0;
  color: #333;
  font-size: 1.25rem;
}

.post-content {
  margin: 0 0 1rem 0;
  color: #555;
  line-height: 1.5;
}

.post-footer {
  display: flex;
  justify-content: flex-end;
  font-size: 0.875rem;
  color: #888;
}

@media (max-width: 768px) {
  .post-container {
    flex-direction: column;
  }  
  .post-form-section, .posts-section {
    width: 100%;
  }
}
</style>