<script lang="ts">
    import {Route} from "svelte-navigator";
    import PrivateRouteGuard from "./PrivateRouteGuard.svelte";

    export let path: string;
</script>

<Route {path}>
    <PrivateRouteGuard>
        <slot/>
    </PrivateRouteGuard>
</Route>