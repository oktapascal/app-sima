<script lang="ts">
    import tooltip from "@/libs/tooltip";
    import {IconBell} from "@/components";

    export let total: number = 0;

    $: totalConvert = total > 9 ? "9+" : total;
    $: rightPosition = total >= 10 ? "right-0" : "right-2";
</script>

<div class="mx-1 relative">
    <button use:tooltip data-tooltip-template="tooltip-notification" type="button"
            class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200 dark:hover:bg-gray-600"
    >
        <IconBell className="h-7 w-7 fill-current text-gray-500 dark:text-white"/>
    </button>
    {#if total > 0}
        <span class="badge {rightPosition}">{totalConvert}</span>
    {/if}
    <div id="tooltip-notification" role="tooltip" class="hidden">
        Notification
    </div>
</div>

<style lang="postcss">
    .badge {
        @apply absolute text-white top-1 rounded-full bg-blue-700 p-1 dark:bg-gray-800;
        font-size: 0.6875rem;
        line-height: 0.625rem;
    }
</style>

