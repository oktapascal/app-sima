<script lang="ts">
    import {breakpoint} from "@/stores/breakpointStore";
    import {IconMenu, IconMenuOpen, WebNavbar, MobileNavbar} from "@/components";

    let isOpen: boolean = false;

    function onToggleMenu() {
        isOpen = !isOpen;
    }
</script>

<div class="flex-1">
    {#if $breakpoint}
        <div class="flex flex-wrap">
            <div>
                <button type="button"
                        class="border border-gray-300 rounded-lg p-2.5 mr-2 lg:mr-0 focus:outline-none focus:ring-2 focus:ring-gray-400"
                        on:click={onToggleMenu}>
                    {#if isOpen}
                        <IconMenuOpen className="w-7 h-7 text-gray-500 dark:text-white"/>
                    {:else }
                        <IconMenu className="w-7 h-7 text-gray-500 dark:text-white"/>
                    {/if}
                </button>
            </div>
        </div>
    {/if}
</div>
<div class="flex-none">
    {#if $breakpoint}
        <MobileNavbar/>
    {:else }
        <WebNavbar/>
    {/if}
</div>