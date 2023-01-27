<script lang="ts">
    import {Loading} from "@/components";

    export let fullWidth: boolean = false;
    export let label: string;
    export let disabled: boolean = false;
    export let type: "button" | "submit" | "reset" = "button";
    export let className: string = "";

    let disabledClass: string = "";
    let fullWidthClass: string = "";

    $: disabledClass = disabled ? "cursor-not-allowed bg-red-500 dark:bg-gray-500" : "";
    $: fullWidthClass = fullWidth ? "w-full" : "";
</script>

<button class="bg-red-700 text-white hover:bg-red-800 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 focus:outline-none focus:ring-2 focus:ring-blue-300 dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus-gray-500 {className} {fullWidthClass} {disabledClass}"
        {type} {disabled}>
    {#if disabled}
        <Loading/>
    {/if}
    {label}
</button>