<script setup lang="ts">
import IconSun from "@/components/icon/IconSun.vue" // Import IconSun component
import IconMoon from "@/components/icon/IconMoon.vue"; // Import IconMoon component
import { useThemeStore } from "@/stores/themes"; // Import useThemeStore hook from themes store

// Initialize themeStore hook with useThemeStore hook
const themeStore = useThemeStore()

// Declare toggleThemes function that takes a value of type string
function toggleThemes(value: string) {
  if(value === "light") { // If the value is "light"
    themeStore.toggleMode("dark") // Set the theme mode to dark
    document.documentElement.classList.add("dark"); // Add the "dark" class to the document element
  } else { // If the value is not "light"
    themeStore.toggleMode("light") // Set the theme mode to light
    document.documentElement.classList.remove("dark"); // Remove the "dark" class from the document element
  }
}
</script>

<template>
  <div class="flex-1"></div>
  <div class="flex-none">
      <button type="button" @click="toggleThemes(themeStore.getMode)" class="border border-gray-400 rounded-lg p-2.5 mr-2 lg:mr-0 focus:outline-none focus:ring-2 focus:ring-gray-400">
        <IconSun v-if="themeStore.isDark" className="w-7 h-7 fill-current dark:text-white"/>
        <IconMoon v-else className="w-7 h-7 fill-current dark:text-white"/>
      </button>
  </div>
</template>
