import { createApp } from 'vue'
import App from './App.vue'
import SigninPage from './Pages/SignIn.vue'
import { createMemoryHistory, createRouter } from 'vue-router'
import SignUp from './Pages/SignUp.vue'
import SignIn from './Pages/SignIn.vue'
import About from './Pages/About.vue'

const routes = [
    { path: '/', component: App },
    { path: '/Signin', component: SignIn },
    { path: '/SignUp', component: SignUp },
    { path: '/about', component: About }
  ]

const router = createRouter({
    history: createMemoryHistory(),
    routes,
  })

createApp(App)
    .use(router)
    .mount('#app')