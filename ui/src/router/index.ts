import {createRouter, createWebHistory} from "vue-router";
import Login from "@/views/Login.vue";
import MainProtected from "@/views/protected/MainProtected.vue";
import {useAuthStore} from "@/stores/auth";

const Dashboard = () => import("../views/protected/Dashboard.vue")

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
                }
            ]
        }
    ],
});

router.beforeEach(async (to,from, next) => {
    const auth = useAuthStore()

    if(!auth.getStatusAuthenticated) {
        await auth.setUserSessionFromCredentials()
    }

    if(to.meta.requiresAuth) {
        if(auth.getStatusAuthenticated) {
            return next()
        } else {
            return next({ path: "/login" })
        }
    } else {
        if(auth.getStatusAuthenticated) {
            return next({ path: "/protected/dashboard"})
        } else {
            return next()
        }
    }
})

export default router;
