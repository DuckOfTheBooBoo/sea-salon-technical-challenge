import { createWebHistory, createRouter, RouteRecordRaw } from 'vue-router';
import LandingPage from '@/views/LandingPage.vue';
import Login from '@/views/auth/LoginPage.vue';
import SignUpPage from '@/views/auth/SignUpPage.vue';

const routes: RouteRecordRaw[] = [
    {
        path: '/',
        name: 'landing',
        component: LandingPage
    },
    {
        path: '/login',
        name: 'login',
        component: Login
    },
    {
        path: '/signup',
        name: 'signup',
        component: SignUpPage
    }
]

export default createRouter({
    history: createWebHistory(),
    routes,
})