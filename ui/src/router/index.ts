import {createRouter, createWebHistory} from "vue-router";
import Login from "@/views/Login.vue";
import MainProtected from "@/views/protected/MainProtected.vue";
import {useAuthStore} from "@/stores/auth";

const Dashboard = () => import("../views/protected/Dashboard.vue")
const Profile = () => import("../views/protected/Profile.vue")

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            redirect: "/login"
        },
        {
            path: "/login",
            component: Login,
            meta: {
                title: "Login Page",
                requiresAuth: false
            },
        },
        {
            path: "/protected",
            component: MainProtected,
            children: [
                {
                    path: "dashboard",
                    component: Dashboard,
                    meta: {
                        requiresAuth: true
                    }
                },
                {
                    path: "profile",
                    component: Profile,
                    meta: {
                        requiresAuth: true
                    }
                }
            ]
        }
    ],
});

// router.beforeEach(async (to, from, next) => {
//     // Initialize auth hook with useAuthStore hook
//     const auth = useAuthStore()
//
//     // If the user is not authenticated
//     if (!auth.getStatusAuthenticated) {
//         // Set the user session using their stored credentials
//         await auth.setUserSessionFromCredentials()
//     }
//
//     // If the route that the user is trying to access requires authentication
//     if (to.meta.requiresAuth) {
//         // If the user is authenticated
//         if (auth.getStatusAuthenticated) {
//             // Allow the user to access the route
//             return next()
//         } else {
//             // Redirect the user to the login page
//             return next({ path: "/login" })
//         }
//     } else {
//         // If the user is authenticated
//         if (auth.getStatusAuthenticated) {
//             // Redirect the user to the dashboard
//             return next({ path: "/protected/dashboard"})
//         } else {
//             // Allow the user to access the route
//             return next()
//         }
//     }
// });

export default router;