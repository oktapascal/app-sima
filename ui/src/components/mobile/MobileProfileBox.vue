<script lang="ts" setup>
import { computed } from "vue"; // Import computed function from vue
import { useRouter } from "vue-router"; // Import useRouter hook from vue-router
import { useThemeStore } from "@/stores/themes"; // Import useThemeStore hook from themes store
import { useAuthStore } from "@/stores/auth"; // Import useAuthStore hook from auth store
import IconWindowClose from "@/components/icon/IconWindowClose.vue"; // Import IconWindowClose component
import IconMoon from "@/components/icon/IconMoon.vue"; // Import IconMoon component
import IconLogout from "@/components/icon/IconLogout.vue"; // Import IconLogout component
import instance from "@/api/instance"; // Import instance object
import { type IAlert, useAlertStore } from "@/stores/alert"; // Import IAlert type and useAlertStore hook from alert store

// Initialize authStore hook with useAuthStore hook
const authStore = useAuthStore()

// Initialize themeStore hook with useThemeStore hook
const themeStore = useThemeStore()

// Initialize alertStore hook with useAlertStore hook
const alertStore = useAlertStore()

// Initialize router hook with useRouter hook
const router = useRouter()

// Declare emits object with "onCloseProfileBox" key
const emits = defineEmits(["onCloseProfileBox"])

// Initialize themeText as a computed value that returns a string depending on the value of themeStore.isDark
const themeText = computed(() => {
  return themeStore.isDark ? "Dark Mode" : "Light Mode"
})

// Declare onCloseProfileBox function that does not take any arguments and does not return a value
function onCloseProfileBox() {
  // Emit "onCloseProfileBox" event with false as the value
  emits("onCloseProfileBox", false)
}

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

// Declare onSignOut function that returns a Promise
async function onSignOut() {
  try {
    // Make a POST request to the "/auth/logout" endpoint using the instance object
    await instance.post("/auth/logout")

    // Reset the auth store
    authStore.$reset()

    // Navigate to the "/login" route
    await router.push({ path: "/login" })
  } catch (error: unknown) {
    // If an error occurs, create an alert object with the appropriate values
    const alert: IAlert = {
      show: true,
      type: "error",
      text: "Terjadi kesalahan tidak diketahui"
    }

    // Show the alert by calling showAlert function on the alertStore
    alertStore.showAlert(alert)
  }
}
</script>

<template>
  <div class="absolute top-0 h-screen w-full bg-gray-100 dark:bg-black z-50">
    <div class="relative flex flex-col">
      <div class="flex-none">
          <button class="focus:outline-none font-medium rounded-full p-2 inline-flex items-center" @click="onCloseProfileBox">
            <IconWindowClose className="h-7 w-7 fill-current text-gray-500 dark:text-white"/>
          </button>
      </div>
      <hr class="border-t border-gray-500" />
      <div class="flex-1">
        <div class="flex flex-col p-2.5">
          <div class="flex-none">
            <h6 class="text-lg text-gray-500 font-bold dark:text-white">Profile</h6>
          </div>
          <div class="flex-1 mt-2">
            <button type="button" class="rounded-lg px-2 py-2.5 inline-flex items-center w-full hover:bg-gray-200 focus:ring-2 focus:outline-none focus:ring-gray-300 dark:hover:bg-gray-600 dark:focus:ring-gray-500">
              <img alt="avatar" src="/images/avatars/avatar.png" class="w-12 h-12 rounded-full"/>
              <span class="text-gray-500 font-medium text-md text-center pl-4 dark:text-white">Nama User</span>
            </button>
          </div>
        </div>
      </div>
      <hr class="border-t border-gray-500" />
      <div class="flex-1">
        <div class="flex flex-col p-2.5">
          <div class="flex-none">
            <h6 class="text-lg text-gray-500 font-bold dark:text-white">Settings</h6>
          </div>
          <div class="flex-1 mt-2">
            <button type="button" class="rounded-lg px-2 py-2.5 inline-flex items-center w-full hover:bg-gray-200 focus:ring-2 focus:outline-none focus:ring-gray-300 dark:hover:bg-gray-600 dark:focus:ring-gray-500" @click="toggleThemes(themeStore.getMode)">
              <IconMoon className="h-7 w-7 fill-current text-gray-500 dark:text-white"/>
              <span class="text-gray-500 font-medium text-md text-center pl-4 dark:text-white">{{ themeText }}</span>
            </button>
          </div>
          <div class="flex-1 mt-2">
            <button type="button" class="rounded-lg px-2 py-2.5 inline-flex items-center w-full hover:bg-gray-200 focus:ring-2 focus:outline-none focus:ring-gray-300 dark:hover:bg-gray-600 dark:focus:ring-gray-500" @click="onSignOut">
              <IconLogout className="h-7 w-7 fill-current text-gray-500 dark:text-white"/>
              <span class="text-gray-500 font-medium text-md text-center pl-4 dark:text-white">Sign Out</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>