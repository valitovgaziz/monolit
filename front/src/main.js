import { createApp } from 'vue'
import App from './App.vue'
import SigninPage from './Pages/SignIn.vue'
import { createMemoryHistory, createRouter } from 'vue-router'
import SignUp from './Pages/SignUp.vue'

const routes = [
    { path: '/', component: App },
    { path: '/Signin', component: SigninPage },
    { path: '/SignUp', component: SignUp }
  ]

const router = createRouter({
    history: createMemoryHistory(),
    routes,
  })

createApp(App)
    .use(router)
    .mount('#app')