<script lang="ts" setup>
import {computed} from "vue";
import {useRouter} from "vue-router";
import {useThemeStore} from "@/stores/themes";
import {useAuthStore} from "@/stores/auth";
import IconWindowClose from "@/components/icon/IconWindowClose.vue";
import IconMoon from "@/components/icon/IconMoon.vue";
import IconLogout from "@/components/icon/IconLogout.vue";
import instance from "@/api/instance";
import {type IAlert, useAlertStore} from "@/stores/alert";

const authStore = useAuthStore()
const themeStore = useThemeStore()
const alertStore = useAlertStore()

const router = useRouter()

const emits = defineEmits(["onCloseProfileBox"])

const themeText = computed(() => {
  return themeStore.isDark ? "Dark Mode" : "Light Mode"
})

function onCloseProfileBox() {
  emits("onCloseProfileBox", false)
}

function toggleThemes(value: string) {
  if(value === "light") {
    themeStore.toggleMode("dark")
    document.documentElement.classList.add("dark");
  } else {
    themeStore.toggleMode("light")
    document.documentElement.classList.remove("dark");
  }
}

async function onSignOut() {
  try {
    await instance.post("/auth/logout")

    authStore.$reset()
    await router.push({ path: "/login" })
  } catch (error: unknown) {
    const alert: IAlert = {
      show: true,
      type: "error",
      text: "Terjadi kesalahan tidak diketahui"
    }

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