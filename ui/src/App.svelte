<script lang="ts">
    import "./assets/css/app.css";
    import {QueryClient, QueryClientProvider} from "@tanstack/svelte-query";
    import {onMount} from "svelte";
    import {Router, Route} from "svelte-navigator";
    import Portal from "svelte-portal";
    import {theme} from "@/stores/themeStore";
    import {alertStore} from "@/stores/alertStore";
    import {auth} from "@/stores/authStore";
    import {breakpoint} from "@/stores/breakpointStore";
    import {NavbarBase, NavbarUser, AlertDefault, NavbarGuest, WebSidebar, WebSettingbar} from "@/components";
    import LazyRoute from "@/pages/LazyRoute.svelte";
    import PublicRoute from "@/pages/PublicRoute.svelte";
    import PrivateRoute from "@/pages/PrivateRoute.svelte";
    import LoadingPage from "@/components/loadings/LoadingPage.svelte";
    import Redirect from "@/pages/Redirect.svelte";
    import NotFound from "@/pages/NotFound.svelte";

    const queryClient = new QueryClient();

    const delayModuleLoad = module =>
        new Promise(res =>
            setTimeout(() => res(module), Math.random() * 2000),
        );

    const Login = () => import("./pages/Login.svelte").then(delayModuleLoad);
    const Dashboard = () => import("./pages/protected/Dashboard.svelte").then(delayModuleLoad);
    const Profile = () => import("./pages/protected/Profile.svelte").then(delayModuleLoad);

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
    <QueryClientProvider client={queryClient}>
        <NavbarBase>
            {#if $auth.isAuthenticated}
                <NavbarUser/>
            {:else }
                <NavbarGuest/>
            {/if}
        </NavbarBase>
        {#if $auth.isAuthenticated}
            <WebSidebar/>
            <WebSettingbar/>
        {/if}
        <main class="lg:pl-23 lg:pr-2.5 lg:pb-2.5">
            <Route path="/" primary={false} let:location>
                <Redirect/>
            </Route>

            <PublicRoute path="login" let:location>
                <LazyRoute component={Login} delayMs={500}>
                    <LoadingPage/>
                </LazyRoute>
            </PublicRoute>

            <PrivateRoute path="dashboard" let:location>
                <LazyRoute component={Dashboard} delayMs={500}>
                    <LoadingPage/>
                </LazyRoute>
            </PrivateRoute>
            <PrivateRoute path="profile" let:location>
                <LazyRoute component={Profile} delayMs={500}>
                    <LoadingPage/>
                </LazyRoute>
            </PrivateRoute>

            <Route>
                <NotFound/>
            </Route>
        </main>
    </QueryClientProvider>
</Router>

<Portal target="body">
    {#if $alertStore.show}
        <AlertDefault/>
    {/if}
</Portal>