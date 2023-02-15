<script lang="ts">
    import {auth} from "@/stores/authStore";
    import {IconChevronDown, IconChevronUp} from "@/components";

    let openBottomMenu: boolean = false;

    function toggleBottomMenu() {
        openBottomMenu = !openBottomMenu;
    }
</script>

<div class="group fixed left-0 top-0 z-50 h-full transition w-20 bg-gray-50 shadow-xl border-r-3xl flex flex-col transition-width ease hover:w-81">
    <div class="flex justify-start items-center p-1.5">
        <img alt="logo" src="/images/sima.png" class="h-10"/>
    </div>
    <div class="pb-8 h-full scroll-auto"></div>
    <div class="flex flex-col justify-center pb-6">
        <button type="button"
                class="flex flex-row items-center cursor-pointer pt-8 pb-4 px-5 text-gray-500 dark:text-white"
                on:click={toggleBottomMenu}>
            <img alt="avatar" src="{import.meta.env.VITE_API_URL_AVATAR}/{$auth.photo}"
                 class="border-2 border-solid border-transparent rounded-full h-11 w-11"/>
            <span class="transition-all whitespace-nowrap flex flex-col font-semibold ml-1.5 invisible text-gray-500 group-hover:visible dark:text-white">{$auth.nik}</span>
            {#if openBottomMenu}
                <IconChevronUp
                        className="h-6 w-6 ml-auto text-gray-500 whitespace-nowrap invisible group-hover:visible dark:text-white"/>
            {:else }
                <IconChevronDown
                        className="h-6 w-6 ml-auto text-gray-500 whitespace-nowrap invisible group-hover:visible dark:text-white"/>
            {/if}
        </button>
    </div>
</div>