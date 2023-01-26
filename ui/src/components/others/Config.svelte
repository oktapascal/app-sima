<script lang="ts">
    import tooltip from "@/libs/tooltip";
    import {IconCog} from "@/components";
</script>

<div class="mx-0.5">
    <button use:tooltip data-tooltip-template="tooltip-config" type="button"
            class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200 dark:hover:bg-gray-600">
        <IconCog className="h-7 w-7 text-gray-500 dark:text-white"/>
    </button>
    <div id="tooltip-config" role="tooltip" class="hidden">
        Settings
    </div>
</div>