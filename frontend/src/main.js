//import './assets/main.css'
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:8080'; 
axios.defaults.withCredentials = true; 

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Toast) 

app.mount('#app')
