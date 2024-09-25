import { createApp } from 'vue'
import App from './App.vue'
import SigninPage from './components/SigninPage.vue'
import { createMemoryHistory, createRouter } from 'vue-router'

const routes = [
    { path: '/', component: App },
    { path: '/Signin', component: SigninPage },
  ]

const router = createRouter({
    history: createMemoryHistory(),
    routes,
  })

createApp(App)
    .use(router)
    .mount('#app')