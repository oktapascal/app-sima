<script lang="ts">
    import {fly} from "svelte/transition";
    import {theme} from "@/stores/themeStore";
    import {rightSideStore} from "@/stores/rightSideStore";
    import {clickOutside} from "@/libs/clickOutside";

    $: isDarkMode = $theme === "dark";

    function onClickOutside() {
        rightSideStore.set(false);
    }

    function onChangeRadio(event) {
        const value = event.currentTarget.value;

        if (value === "true") {
            theme.set("dark");
            window.localStorage.setItem("theme", "dark");
            document.documentElement.classList.add("dark");
        } else {
            theme.set("light");
            window.localStorage.setItem("theme", "light");
            document.documentElement.classList.remove("dark");
        }
    }
</script>

{#if $rightSideStore}
    <div use:clickOutside on:click_outside={onClickOutside}
         class="container-config-app fixed h-full w-52 right-0 top-0 bg-gray-50 dark:bg-gray-700"
         in:fly={{ x: 100, duration: 500 }}
         out:fly={{ x: 200, duration: 1000 }}>
        <div class="flex flex-col">
            <div class="px-4 py-3">
                <h5 class="text-lg font-semibold dark:text-white">Theme App</h5>
                <div>
                    <label class="my-2 cursor-pointer justify-start radio-control">
                        <input type="radio" name="radio-theme" class="radio-custom" value="{false}"
                               bind:group={isDarkMode}
                               on:change={onChangeRadio}/>
                        <span class="label-text font-semibold text-sm ml-2 dark:text-white">Light</span>
                    </label>
                    <label class="my-2 cursor-pointer justify-start radio-control">
                        <input type="radio" name="radio-theme" class="radio-custom" value="{true}"
                               bind:group={isDarkMode}
                               on:change={onChangeRadio}/>
                        <span class="label-text font-semibold text-sm ml-2 dark:text-white">Dark</span>
                    </label>
                </div>
            </div>
            <hr class="border-t border-gray-400"/>
        </div>
    </div>
{/if}

<style>
    .container-config-app {
        box-shadow: 0 2px 4px -1px rgb(0 0 0 / 20%), 0 4px 5px 0 rgb(0 0 0 / 14%), 0 1px 10px 0 rgb(0 0 0 / 12%);
        z-index: 99;
    }
</style>