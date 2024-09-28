import { createApp } from 'vue'
import App from './App.vue'
import SigninPage from './components/Pages/SignIn.vue'
import { createMemoryHistory, createRouter } from 'vue-router'
import SignUp from './components/Pages/SignUp.vue'
import SignIn from './components/Pages/SignIn.vue'
import About from './components/Pages/About.vue'

const routes = [
    { path: '/',
      name: 'home',
      component: App },
    { path: '/SingIn',
      name: 'SingIn',
      component: SignIn },
    { path: '/SignUp',
      name: 'SingUp',
      component: SignUp },
    { path: '/About',
      name: 'About',
      component: About },
  ]

const router = createRouter({
    history: createMemoryHistory(),
    routes,
  })

createApp(App)
    .use(router)
    .mount('#app')