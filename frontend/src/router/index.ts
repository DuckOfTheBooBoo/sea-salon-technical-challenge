import { createWebHistory, createRouter, type RouteRecordRaw, type RouteLocationNormalized, type NavigationGuardNext } from 'vue-router';
import LandingPage from '@/views/LandingPage.vue';
import Login from '@/views/auth/LoginPage.vue';
import SignUpPage from '@/views/auth/SignUpPage.vue';
import AdminDashboard from '@/views/admin/Dashboard.vue';
import CustomerDashboard from '@/views/customer/Dashboard.vue';
import ForbiddenPage from '@/views/ForbiddenPage.vue';
import { getUser } from '@/service/userApi';
import type { User, JWTPayload } from '@/types';
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

router.beforeEach(async (to, from, next) => {
  try {
    let isValid: boolean = localStorage.getItem('token') ? true : false;
    // decode jwt
    const token: string | null = localStorage.getItem("token")

    console.log(token)

    if (!token) {
        next({ name: 'login' });
        return
    }

    const decoded = jwtDecode<JWTPayload>(token);
    
    console.log(decoded)

    if (decoded.exp !== undefined) {
        if (decoded.exp * 1000 > Date.now()) {
            isValid = true;
        } else {
            isValid = false;
        }
    }

    let redirectTo: string | undefined;

    if (!isValid && to.name !== 'login' && to.name !== 'signup') {
        redirectTo = 'login';
      } else if (isValid && (to.name === 'login' || to.name === 'signup')) {
        console.log('HAI')
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
      } else if (isValid && to.name === 'admin-dashboard' && decoded.Role !== 'Admin') {
        redirectTo = 'forbidden';
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