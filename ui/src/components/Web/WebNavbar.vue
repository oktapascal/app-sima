<script lang="ts" setup>
import {useRouter} from "vue-router";
import {Notification, ConfigApp, IconLogout} from "@/components";
import instance from "@/api/instance";
import {useAlertStore} from "@/stores/alert";
import {useAuthStore} from "@/stores/auth";
import type {IAlert} from "@/types";

const authStore = useAuthStore();

const alertStore = useAlertStore();

const router = useRouter();

async function onSignOut() {
  try {
    await instance.post("/auth/logout");

    authStore.$reset();

    await router.push({path: "/login"});
  } catch (error: unknown) {
    const alert: IAlert = {
      show: true,
      type: "error",
      text: "Terjadi kesalahan tidak diketahui",
    };

    alertStore.showAlert(alert);
  }
}
</script>

<template>
  <div class="flex flex-row flex-nowrap mr-2">
    <div class="flex-1">
      <Notification/>
    </div>
    <div class="flex-1">
      <ConfigApp/>
    </div>
    <div class="flex-none">
      <button data-tooltip-target="tooltip-bottom-2" data-tooltip-placement="bottom"
              class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200 dark:hover:bg-gray-600"
              @click="onSignOut">
        <IconLogout className="h-7 w-7 fill-current text-gray-500 dark:text-white"/>
      </button>
      <div id="tooltip-bottom-2" role="tooltip"
           class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium bg-gray-300/60 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700/60 dark:text-white">
        Sign Out
        <div class="tooltip-arrow" data-popper-arrow></div>
      </div>
    </div>
  </div>
</template>