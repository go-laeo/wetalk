import { createPinia } from 'pinia'
import { createApp } from 'vue'
import Oruga from '@oruga-ui/oruga-next'
import '@oruga-ui/oruga-next/dist/oruga-full-vars.css'
import App from './App.vue'
import router from './router'

const pinia = createPinia()
pinia.use(({ store }) => {
    store.router = router
})

const app = createApp(App)
app.use(Oruga)
app.use(pinia)
app.use(router)
app.mount('#app')
