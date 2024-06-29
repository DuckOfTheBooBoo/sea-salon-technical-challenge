import { createWebHistory, createRouter, type RouteRecordRaw, type RouteLocationNormalized, type NavigationGuardNext } from 'vue-router';
import LandingPage from '@/views/LandingPage.vue';
import Login from '@/views/auth/LoginPage.vue';
import SignUpPage from '@/views/auth/SignUpPage.vue';
import AdminDashboard from '@/views/admin/Dashboard.vue';
import CustomerDashboard from '@/views/customer/Dashboard.vue';
import { getUser } from '@/service/userApi';
import { User } from '@/types';

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
    },
    {
        path: '/admin',
        name: 'admin-dashboard',
        component: AdminDashboard,
        meta: {protected: true}
    },
    {
        path: '/customer',
        name: 'customer-dashboard',
        component: CustomerDashboard,
        meta: {protected: true}
    }
]

const router = createRouter({ history: createWebHistory(), routes });

// We could use GET /api/users to check if token is valid, the endpoint itself will return a 401 if token is invalid, else returns the user information
// const checkTokenValidation = async (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext): Promise<void> => {
//     try {
//         const user: User = await getUser()
        
//         if (!isValid && to.name !== 'login' && to.name !== 'signup') {
//           next({ name: 'login' });
//         } else if (isValid && (to.name === 'login' || to.name === 'signup')) {
//           next({ name: 'explorer-files' });
//         } else {
//           next();
//         }
//       } catch (error: unknown) {
//         console.error('Router Token Check Error:', error);
//         next(false);
//     }
//   }  

// router.beforeEach(async (to, from, next) => {
//   try {
//     const isValid = await checkTokenValidation();
    
//     if (!isValid && to.name !== 'login' && to.name !== 'signup') {
//       next({ name: 'login' });
//     } else if (isValid && (to.name === 'login' || to.name === 'signup')) {
//       next({ name: 'explorer-files' });
//     } else {
//       next();
//     }
//   } catch (error) {
//     console.error('Router Token Check Error:', error);
//     next(false);
//   }
// });


export default router;