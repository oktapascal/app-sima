<script lang="ts">
    import {Loading} from "@/components";

    export let label: string;
    export let disabled: boolean = false;
    export let type: "button" | "submit" | "reset" = "button";
    export let className: string = "";

    let disabledClass: string = "";
    $: disabledClass = disabled ? "cursor-not-allowed bg-blue-500 dark:bg-gray-500" : "";
</script>

<button class="bg-blue-700 text-white hover:bg-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 focus:outline-none focus:ring-2 focus:ring-blue-300 dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus-gray-500 {className} {disabledClass}"
        {type} {disabled}>
    {#if disabled}
        <Loading/>
    {/if}
    {label}
</button>