import { createApp } from 'vue'
import App from './App.vue'
import SigninPage from './Pages/SignIn.vue'
import { createMemoryHistory, createRouter } from 'vue-router'
import SignUp from './Pages/SignUp.vue'
import SignIn from './Pages/SignIn.vue'
import About from './Pages/About.vue'

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