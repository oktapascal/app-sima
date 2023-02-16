<script lang="ts">
    import tooltip from "@/libs/tooltip";
    import {theme} from "@/stores/themeStore";

    $: isDarkMode = $theme === "dark";
    $: textTooltipTheme = $theme === "dark" ? "Toggle Light Mode" : "Toggle Dark Mode";

    function toggleTheme(value) {
        if (value === "light") {
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

<div class="flex-1"></div>
<div class="flex-none">
    <button use:tooltip data-tooltip-template="tooltip-theme" type="button"
            class="border border-gray-300 rounded-lg p-2.5 mr-2 lg:mr-0 focus:outline-none focus:ring-2 focus:ring-gray-400"
            on:click={() => toggleTheme($theme)}>
        {#if isDarkMode}
            <i class="icofont-sun-alt icofont-2x text-gray-500 dark:text-white"></i>
        {:else}
            <i class="icofont-moon icofont-2x text-gray-500 dark:text-white"></i>
        {/if}
    </button>
    <div id="tooltip-theme" role="tooltip" class="hidden">
        {textTooltipTheme}
    </div>
</div>