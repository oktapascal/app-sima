<script lang="ts" setup>
import { useRouter } from "vue-router"; // Import useRouter hook from vue-router
import Notification from "@/components/Notification.vue"; // Import Notification component
import ConfigApp from "@/components/config/ConfigApp.vue"; // Import ConfigApp component
import IconLogout from "@/components/icon/IconLogout.vue"; // Import IconLogout component
import instance from "@/api/instance"; // Import instance object
import { type IAlert, useAlertStore } from "@/stores/alert"; // Import IAlert type and useAlertStore hook from alert store
import { useAuthStore } from "@/stores/auth"; // Import useAuthStore hook from auth store

// Initialize authStore hook with useAuthStore hook
const authStore = useAuthStore()

// Initialize alertStore hook with useAlertStore hook
const alertStore = useAlertStore()

// Initialize router hook with useRouter hook
const router = useRouter()

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
  <div class="flex flex-row flex-nowrap mr-2">
    <div class="flex-1">
      <Notification />
    </div>
    <div class="flex-1">
      <ConfigApp />
    </div>
    <div class="flex-none">
      <button data-tooltip-target="tooltip-bottom-2" data-tooltip-placement="bottom" class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200" @click="onSignOut">
        <IconLogout className="h-7 w-7 fill-current text-gray-500 dark:text-white"/>
      </button>
      <div id="tooltip-bottom-2" role="tooltip" class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium bg-gray-300 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700 dark:text-white">
        Sign Out
        <div class="tooltip-arrow" data-popper-arrow></div>
      </div>
    </div>
  </div>
</template>