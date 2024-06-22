import { createWebHistory, createRouter, RouteRecordRaw } from 'vue-router';
import LandingPage from '@/views/LandingPage.vue';
import Login from '@/views/LoginPage.vue';

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
    }
]

export default createRouter({
    history: createWebHistory(),
    routes,
})