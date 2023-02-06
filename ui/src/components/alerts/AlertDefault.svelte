<script lang="ts">
    import {onMount} from "svelte";
    import {fade} from "svelte/transition";
    import {IconAlert, IconAlertBox, IconAlertCircle, IconCheckCircle} from "@/components";
    import {alertStore} from "@/stores/alertStore";
    import {breakpoint} from "@/stores/breakpointStore";

    const icon = {
        success: IconCheckCircle,
        error: IconAlert,
        warning: IconAlertCircle,
        info: IconAlertBox,
    };

    const text = {
        success: "Success Message",
        error: "Error Message",
        warning: "Warning Message",
        info: "Information",
    };

    const btnClass = {
        success: "btn-success",
        error: "btn-error",
        warning: "btn-warning",
        info: "btn-default",
    };

    function hideAlert() {
        alertStore.update((state) => {
            return {...state, show: false};
        });
    }

    let timeout: number;
    onMount(() => {
        timeout = setTimeout(hideAlert, 5000);

        return () => clearTimeout(timeout); // destroy
    });
</script>

<div transition:fade
     role="alert"
     class="alert p-4 mb-4 text-white border rounded-lg fixed w-auto {$alertStore.type} {$breakpoint ? 'top-center':'bottom-right'}">
    <div class="flex items-center">
        <svelte:component this={icon[$alertStore.type]} className="icon mr-2 w-6 h-6 text-white"/>
        <h3 class="text-lg font-medium text-white">{text[$alertStore.type]}</h3>
    </div>
    <div class="mt-2 mb-4 text-sm text-white">
        {$alertStore.text}
    </div>
    <div class="flex">
        <button type="button"
                class="btn-alert text-white font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 focus:outline-none focus:ring-2 {btnClass[$alertStore.type]}"
                on:click={hideAlert}>
            Tutup
        </button>
    </div>
</div>