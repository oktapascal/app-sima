<script lang="ts">
    import {Route} from "svelte-navigator";
    import PublicRouteGuard from "./PublicRouteGuard.svelte";

    export let path: string;
</script>

<Route {path}>
    <PublicRouteGuard>
        <slot/>
    </PublicRouteGuard>
</Route>