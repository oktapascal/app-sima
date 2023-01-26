<script lang="ts">
    import {onMount} from "svelte";
    import {useFocus} from "svelte-navigator";

    export let component;
    export let delayMs = null;

    const registerFocus = useFocus();

    let loadedComponent = null;
    let timeout;
    let showFallback = !delayMs;

    let props;
    $: {
        const {component, delayMs, ...restProps} = $$props;
        props = restProps;
    }

    onMount(() => {
        if (delayMs) {
            timeout = setTimeout(() => {
                showFallback = true;
            }, delayMs);
        }

        component().then((module) => {
            loadedComponent = module.default;
        });

        return () => clearTimeout(timeout);
    });
</script>

{#if loadedComponent}
    <svelte:component this={loadedComponent} {registerFocus} {...props}/>
{:else if showFallback}
    <slot/>
{/if}