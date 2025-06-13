import axios from 'axios'
import router from './router'

const api = axios.create({
  baseURL: 'http://localhost:8080', // >> we need to change this to the actual API URL
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
})

api.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response && error.response.status === 401) {
      console.warn('Unauthorized, redirecting to login...')
      router.push('/signin')
    }
    return Promise.reject(error)
  }
)

export default api
