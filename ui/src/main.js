import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:80/api/'
axios.defaults.withCredentials = true

createApp(App).use(router).mount('#app')
