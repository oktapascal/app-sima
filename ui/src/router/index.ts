import {createRouter, createWebHistory} from "vue-router";
import Login from "@/views/Login.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/login",
            component: Login,
            meta: {
                title: "Login Page",
            },
        },
    ],
});

export default router;
