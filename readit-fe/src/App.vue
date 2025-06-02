<template>
  <div class="app-container">
    <header class="app-header">
      <h1>readit-min-v1</h1>
    </header>
    <router-view />
  </div>
</template>

<script>
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;600;700&display=swap');

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'JetBrains Mono', monospace;
  background-color: #f8f9fc;
  height: 100vh;
  overflow-x: hidden;
}

.app-container {
  /* min-height: 100vh; */
  display: flex;
  flex-direction: column;
}

.app-header {
  background: linear-gradient(135deg, #5b6dfd 0%, #4a5ae0 100%);
  color: white;
  padding: 1.5rem 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.app-header h1 {
  font-size: 2rem;
  font-weight: 600;
  letter-spacing: -0.5px;
}

.main-content {
  flex: 1;
  padding: 2rem;
  width: 100%;
  max-width: none;
}
</style>