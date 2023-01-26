<script lang="ts">
    import Lazy from "./Lazy.svelte";

    export let component;
    export let delayMs = null;

    let props;
    $: {
        const {component, ...restprops} = $$props;
        props = restprops;
    }
</script>

<Lazy {component} {delayMs}>
    <slot/>
</Lazy>