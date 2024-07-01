import { createWebHistory, createRouter, type RouteRecordRaw, type RouteLocationNormalized, type NavigationGuardNext } from 'vue-router';
import LandingPage from '@/views/LandingPage.vue';
import Login from '@/views/auth/LoginPage.vue';
import SignUpPage from '@/views/auth/SignUpPage.vue';
import AdminDashboard from '@/views/admin/Dashboard.vue';
import CustomerDashboard from '@/views/customer/Dashboard.vue';
import ForbiddenPage from '@/views/ForbiddenPage.vue';
import type { JWTPayload } from '@/types';
import { jwtDecode } from 'jwt-decode';

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
    },
    {
        path: '/forbidden',
        name: 'forbidden',
        component: ForbiddenPage
    },
]

const router = createRouter({ history: createWebHistory(), routes });

router.beforeEach(async (to, from, next) => {

  // Allow navigation if hash is present (this only happens in landing page) or route is 'landing' 
  if (to.hash !== "" || from.hash !== "" || to.name === 'landing') {
    next();
    return;
  }

  try {
    const token: string | null = localStorage.getItem("token");
    let isValid = false;
    let decoded: any;

    // Check if the token exists and is valid
    if (token) {
      decoded = jwtDecode<JWTPayload>(token);
      if (decoded.exp && decoded.exp * 1000 > Date.now()) {
        isValid = true;
      }
    }

    let redirectTo: string | undefined;

    if (!isValid) {
      // Allow access to 'login' and 'signup' routes for unauthenticated users
      if (to.name === 'login' || to.name === 'signup') {
        next();
        return;
      } else {
        redirectTo = 'login';
      }
    } else {
      // Handle redirection for authenticated users
      if (to.name === 'login' || to.name === 'signup') {
        switch (decoded.role) {
          case 'Customer':
            redirectTo = 'customer-dashboard';
            break;
          case 'Admin':
            redirectTo = 'admin-dashboard';
            break;
          default:
            console.error('Unexpected role or missing decoded.Role');
            break;
        }
      } else if (to.name === 'admin-dashboard' && decoded.role !== 'Admin') {
        redirectTo = 'forbidden';
      }
    }

    if (redirectTo) {
      next({ name: redirectTo });
    } else {
      next();
    }
  } catch (error) {
    console.error('Router Token Check Error:', error);
    next(false);
  }
});


export default router;