import { createRouter, createWebHistory } from "vue-router";
import { createAuthStore } from "../stores/auth";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '',
            redirect: { name: 'Home' },
            children: [
                {
                    name: 'Home',
                    path: '/',
                    component: () => import('../views/Home.vue')
                },
                {
                    name: 'Login',
                    path: '/login',
                    component: () => import('../views/Login.vue')
                },
                {
                    name: 'CreateAccount',
                    path: '/create-account',
                    component: () => import('../views/CreateAccount.vue')
                },
                {
                    name: 'ReadPost',
                    path: '/posts/:id',
                    component: () => import('../views/ReadPost.vue'),
                    props: route => ({ id: parseInt(route.params.id as string) })
                },
                {
                    name: 'Write',
                    path: '/write',
                    component: () => import('../views/Write.vue'),
                    beforeEnter(to, from, next) {
                        const auth = createAuthStore()
                        if (!auth.token) {
                            return next({ name: 'Login' })
                        }
                        return next()
                    }
                },
                {
                    name: 'AccountInfo',
                    path: '/account-info',
                    component: () => import('../views/AccountInfo.vue'),
                    beforeEnter(to, from, next) {
                        const auth = createAuthStore()
                        if (!auth.token) {
                            return next({ name: 'Login' })
                        }
                        return next()
                    }
                },
                {
                    name: 'ForceMarket',
                    path: '/force-market',
                    component: () => import('../views/ForceMarket.vue'),
                    beforeEnter(to, from, next) {
                        const auth = createAuthStore()
                        if (!auth.token) {
                            return next({ name: 'Login' })
                        }
                        return next()
                    }
                },
            ]
        },
    ]
})

export default router
