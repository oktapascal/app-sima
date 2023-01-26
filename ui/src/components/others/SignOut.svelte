<script lang="ts">
    import {createEventDispatcher} from "svelte";
    import tooltip from "@/libs/tooltip";
    import {IconLogout} from "@/components";

    const dispatch = createEventDispatcher();

    function onClick() {
        dispatch("action");
    }
</script>

<div class="mx-0.5">
    <button use:tooltip type="button" data-tooltip-template="tooltip-signout"
            class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200 dark:hover:bg-gray-600"
            on:click={onClick}>
        <IconLogout className="h-7 w-7 text-gray-500 dark:text-white"/>
    </button>
    <div id="tooltip-signout" role="tooltip" class="hidden">
        Sign Out
    </div>
</div>