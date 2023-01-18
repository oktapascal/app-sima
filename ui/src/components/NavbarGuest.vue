<script setup lang="ts">
import {IconSun, IconMoon} from "@/components";
import {useThemeStore} from "@/stores/themes";

const themeStore = useThemeStore();

function toggleThemes(value: string) {
  if (value === "light") {
    themeStore.toggleMode("dark");
    document.documentElement.classList.add("dark");
  } else {
    themeStore.toggleMode("light");
    document.documentElement.classList.remove("dark");
  }
}
</script>

<template>
  <div class="flex-1"></div>
  <div class="flex-none">
    <button type="button" @click="toggleThemes(themeStore.getMode)"
            class="border border-gray-300 rounded-lg p-2.5 mr-2 lg:mr-0 focus:outline-none focus:ring-2 focus:ring-gray-400">
      <IconSun v-if="themeStore.isDark" className="w-7 h-7 fill-current dark:text-white"/>
      <IconMoon v-else className="w-7 h-7 dark:text-white"/>
    </button>
  </div>
</template>
