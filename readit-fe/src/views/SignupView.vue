<template>
  <div class="auth-container">
    <div class="auth-card">
      <div class="auth-header">
        <h2>Create Account</h2>
        <p class="auth-subtitle">Join our community today</p>
      </div>
      
      <form @submit.prevent="signUp" class="auth-form">
        <div class="form-group">
          <label for="email">Email</label>
          <div class="input-wrapper">
            <span class="input-icon">✉️</span>
            <input 
              id="email"
              v-model="email" 
              type="email" 
              placeholder="your@email.com" 
              required 
              autocomplete="email"
            />
          </div>
        </div>
        
        <div class="form-group">
          <label for="password">Password</label>
          <div class="input-wrapper">
            <span class="input-icon">🔒</span>
            <input 
              id="password"
              v-model="password" 
              type="password" 
              placeholder="••••••••" 
              required 
              autocomplete="new-password"
            />
            <button 
              type="button" 
              class="toggle-password" 
              @click="togglePasswordVisibility"
            >
              {{ showPassword ? '👁️' : '👁️‍🗨️' }}
            </button>
          </div>
          <div class="password-strength" v-if="password">
            <div class="strength-meter">
              <div 
                class="strength-value" 
                :style="{ width: passwordStrength + '%' }"
                :class="strengthClass"
              ></div>
            </div>
            <span class="strength-text">{{ strengthText }}</span>
          </div>
        </div>
        
        <div class="form-group">
          <label for="confirmPassword">Confirm Password</label>
          <div class="input-wrapper">
            <span class="input-icon">🔒</span>
            <input 
              id="confirmPassword"
              v-model="confirmPassword" 
              type="password" 
              placeholder="••••••••" 
              required 
              autocomplete="new-password"
            />
          </div>
          <p v-if="passwordMismatch" class="validation-error">
            Passwords do not match
          </p>
        </div>
        
        <div class="terms-agreement">
          <label class="checkbox-container">
            <input type="checkbox" v-model="agreeToTerms" required>
            <span class="checkmark"></span>
            I agree to the <a href="#" class="terms-link">Terms of Service</a> and <a href="#" class="terms-link">Privacy Policy</a>
          </label>
        </div>
        
        <button 
          type="submit" 
          class="submit-button"
          :disabled="isSubmitting || passwordMismatch || !agreeToTerms"
        >
          <span v-if="isSubmitting" class="spinner"></span>
          <span v-else>Create Account</span>
        </button>
      </form>
      
      <div 
        v-if="message" 
        class="message" 
        :class="{ 'success': isSuccess, 'error': !isSuccess }"
      >
        {{ message }}
      </div>
      
      <div class="auth-footer">
        <p>Already have an account? <a href="#" class="signin-link" @click.prevent="goToSignIn">Sign in</a></p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';

const email = ref('');
const password = ref('');
const confirmPassword = ref('');
const message = ref('');
const isSubmitting = ref(false);
const isSuccess = ref(false);
const showPassword = ref(false);
const agreeToTerms = ref(false);
const router = useRouter();

const passwordMismatch = computed(() => {
  return confirmPassword.value && password.value !== confirmPassword.value;
});

const passwordStrength = computed(() => {
  if (!password.value) return 0;
  
  let strength = 0;
  // Length check
  if (password.value.length >= 8) strength += 25;
  // Contains uppercase
  if (/[A-Z]/.test(password.value)) strength += 25;
  // Contains number
  if (/[0-9]/.test(password.value)) strength += 25;
  // Contains special char
  if (/[^A-Za-z0-9]/.test(password.value)) strength += 25;
  
  return strength;
});

const strengthClass = computed(() => {
  if (passwordStrength.value <= 25) return 'weak';
  if (passwordStrength.value <= 50) return 'medium';
  if (passwordStrength.value <= 75) return 'good';
  return 'strong';
});

const strengthText = computed(() => {
  if (passwordStrength.value <= 25) return 'Weak';
  if (passwordStrength.value <= 50) return 'Medium';
  if (passwordStrength.value <= 75) return 'Good';
  return 'Strong';
});

const togglePasswordVisibility = () => {
  showPassword.value = !showPassword.value;
  const passwordInput = document.getElementById('password') as HTMLInputElement;
  if (passwordInput) {
    passwordInput.type = showPassword.value ? 'text' : 'password';
  }
};

const goToSignIn = () => {
  router.push('/signin');
};

const signUp = async () => {
  if (passwordMismatch.value) {
    message.value = 'Passwords do not match.';
    isSuccess.value = false;
    return;
  }
  
  try {
    isSubmitting.value = true;
    message.value = '';
    
    const res = await fetch('/signup', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ 
        email: email.value, 
        password: password.value 
      })
    });
    
    const data = await res.json();
    
    if (res.ok) {
      isSuccess.value = true;
      message.value = 'Account created successfully! Redirecting to sign in...';
      setTimeout(() => router.push('/signin'), 1500);
    } else {
      isSuccess.value = false;
      message.value = data.error || 'Sign up failed. Please try again.';
    }
  } catch (error) {
    isSuccess.value = false;
    message.value = 'Connection error. Please try again later.';
  } finally {
    isSubmitting.value = false;
  }
};
</script>

<style scoped>
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 2rem;
  background: linear-gradient(135deg, #f8f9fc 0%, #eef1f8 100%);
}

.auth-card {
  width: 100%;
  max-width: 420px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  transition: transform 0.3s, box-shadow 0.3s;
}

.auth-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.12);
}

.auth-header {
  padding: 2rem 2rem 1rem;
  text-align: center;
}

.auth-header h2 {
  font-size: 1.75rem;
  font-weight: 600;
  color: #333;
  margin-bottom: 0.5rem;
}

.auth-subtitle {
  color: #666;
  font-size: 0.95rem;
}

.auth-form {
  padding: 1rem 2rem 2rem;
}

.form-group {
  margin-bottom: 1.25rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: #444;
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.input-icon {
  position: absolute;
  left: 1rem;
  color: #888;
  font-size: 0.875rem;
}

input[type="email"],
input[type="password"],
input[type="text"] {
  width: 100%;
  padding: 0.875rem 1rem 0.875rem 2.5rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-family: 'JetBrains Mono', monospace;
  font-size: 0.95rem;
  color: #333;
  background-color: #f9fafb;
  transition: border-color 0.2s, box-shadow 0.2s, background-color 0.2s;
}

input:focus {
  outline: none;
  border-color: #5b6dfd;
  box-shadow: 0 0 0 3px rgba(91, 109, 253, 0.15);
  background-color: #fff;
}

.toggle-password {
  position: absolute;
  right: 1rem;
  background: none;
  border: none;
  cursor: pointer;
  color: #888;
  font-size: 1rem;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.toggle-password:hover {
  color: #555;
}

.password-strength {
  margin-top: 0.75rem;
}

.strength-meter {
  height: 4px;
  background-color: #eee;
  border-radius: 2px;
  margin-bottom: 0.25rem;
}

.strength-value {
  height: 100%;
  border-radius: 2px;
  transition: width 0.3s ease;
}

.strength-value.weak {
  background-color: #ff4d4f;
}

.strength-value.medium {
  background-color: #faad14;
}

.strength-value.good {
  background-color: #52c41a;
}

.strength-value.strong {
  background-color: #1890ff;
}

.strength-text {
  font-size: 0.75rem;
  color: #888;
}

.validation-error {
  color: #ff4d4f;
  font-size: 0.75rem;
  margin-top: 0.5rem;
  animation: shake 0.5s ease-in-out;
}

@keyframes shake {
  0%, 100% { transform: translateX(0); }
  20%, 60% { transform: translateX(-5px); }
  40%, 80% { transform: translateX(5px); }
}

.terms-agreement {
  margin-bottom: 1.5rem;
}

.checkbox-container {
  display: flex;
  align-items: center;
  position: relative;
  padding-left: 28px;
  cursor: pointer;
  font-size: 0.875rem;
  color: #555;
  user-select: none;
  line-height: 1.4;
}

.checkbox-container input {
  position: absolute;
  opacity: 0;
  cursor: pointer;
  height: 0;
  width: 0;
}

.checkmark {
  position: absolute;
  top: 0;
  left: 0;
  height: 18px;
  width: 18px;
  background-color: #f1f3f7;
  border: 1px solid #ddd;
  border-radius: 4px;
  transition: all 0.2s;
}

.checkbox-container:hover input ~ .checkmark {
  background-color: #e9ecf5;
}

.checkbox-container input:checked ~ .checkmark {
  background-color: #5b6dfd;
  border-color: #5b6dfd;
}

.checkmark:after {
  content: "";
  position: absolute;
  display: none;
}

.checkbox-container input:checked ~ .checkmark:after {
  display: block;
}

.checkbox-container .checkmark:after {
  left: 6px;
  top: 2px;
  width: 4px;
  height: 9px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

.terms-link {
  color: #5b6dfd;
  text-decoration: none;
  transition: color 0.2s;
}

.terms-link:hover {
  color: #4a5ae0;
  text-decoration: underline;
}

.submit-button {
  width: 100%;
  padding: 0.875rem;
  background: linear-gradient(135deg, #5b6dfd 0%, #4a5ae0 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-family: 'JetBrains Mono', monospace;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s, opacity 0.2s;
  position: relative;
  overflow: hidden;
}

.submit-button:hover {
  box-shadow: 0 4px 12px rgba(91, 109, 253, 0.3);
  transform: translateY(-2px);
}

.submit-button:active {
  transform: translateY(0);
}

.submit-button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.spinner {
  display: inline-block;
  width: 20px;
  height: 20px;
  border: 3px solid rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  border-top-color: white;
  animation: spin 1s ease-in-out infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.message {
  margin: 0 2rem 1.5rem;
  padding: 0.75rem;
  border-radius: 6px;
  font-size: 0.875rem;
  text-align: center;
  animation: fadeIn 0.3s ease-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(-10px); }
  to { opacity: 1; transform: translateY(0); }
}

.success {
  background-color: #e7f8ef;
  color: #0c6b58;
  border: 1px solid #a7e9d4;
}

.error {
  background-color: #feeef0;
  color: #b71c1c;
  border: 1px solid #fccdd2;
}

.auth-footer {
  padding: 1.5rem 2rem;
  text-align: center;
  background-color: #f9fafb;
  border-top: 1px solid #eee;
}

.auth-footer p {
  font-size: 0.875rem;
  color: #666;
}

.signin-link {
  color: #5b6dfd;
  text-decoration: none;
  font-size: 0.875rem;
  font-weight: 500;
  transition: color 0.2s;
}

.signin-link:hover {
  color: #4a5ae0;
  text-decoration: underline;
}

/* Responsive adjustments */
@media (max-width: 480px) {
  .auth-card {
    box-shadow: none;
    border-radius: 0;
  }
  
  .auth-container {
    padding: 1rem;
    background: white;
  }
  
  .auth-header h2 {
    font-size: 1.5rem;
  }
}
</style>