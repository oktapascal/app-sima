<script lang="ts">
    import "./assets/css/app.css";
    import {onMount} from "svelte";
    import {Router} from "svelte-navigator";
    import Portal from "svelte-portal";
    import {theme} from "@/stores/themeStore";
    import {alert} from "@/stores/alertStore";
    import {auth} from "@/stores/authStore";
    import {breakpoint} from "@/stores/breakpointStore";
    import {NavbarBase, NavbarUser, AlertDefault, NavbarGuest} from "@/components";
    import LazyRoute from "@/pages/LazyRoute.svelte";
    import PublicRoute from "@/pages/PublicRoute.svelte";
    import PrivateRoute from "@/pages/PrivateRoute.svelte";
    import LoadingPage from "@/components/loadings/LoadingPage.svelte";
    
    const delayModuleLoad = module =>
        new Promise(res =>
            setTimeout(() => res(module), Math.random() * 2000),
        );

    const Login = () => import("./pages/Login.svelte").then(delayModuleLoad);
    const Dashboard = () => import("./pages/protected/Dashboard.svelte").then(delayModuleLoad);

    function onResize() {
        const width = window.screen.width;

        if (320 <= width && 768 >= width) {
            breakpoint.set(true);
        } else {
            breakpoint.set(false);
        }
    }

    onMount(() => {
        if ($theme === "dark") {
            document.documentElement.classList.add("dark");
            window.localStorage.setItem("theme", "dark");
        } else {
            document.documentElement.classList.remove("dark");
            window.localStorage.setItem("theme", "light");
        }

        onResize();

        window.addEventListener("resize", onResize);

        return () => window.removeEventListener("resize", onResize); // destroy
    });
</script>

<Router>
    <NavbarBase>
        {#if $auth.isAuthenticated}
            <NavbarUser/>
        {:else }
            <NavbarGuest/>
        {/if}
    </NavbarBase>
    <main class="lg:px-2.5 lg:pb-2.5">
        <PublicRoute path="login" let:location>
            <LazyRoute component={Login} delayMs={500}>
                <LoadingPage/>
            </LazyRoute>
        </PublicRoute>

        <PrivateRoute path="protected/dashboard" let:location>
            <LazyRoute component={Dashboard} delayMs={500}>
                <LoadingPage/>
            </LazyRoute>
        </PrivateRoute>
    </main>
</Router>

<Portal target="body">
    {#if $alert.show}
        <AlertDefault/>
    {/if}
</Portal>