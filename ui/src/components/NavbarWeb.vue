<script lang="ts" setup>
import Notification from "@/components/Notification.vue";
import IconLogout from "@/components/icon/IconLogout.vue";
</script>

<template>
  <div class="flex flex-row flex-nowrap mr-2">
    <div class="flex-1">
      <Notification/>
    </div>
    <div class="flex-none">
      <button data-tooltip-target="tooltip-bottom-2" data-tooltip-placement="bottom" class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200">
        <IconLogout className="h-7 w-7 fill-current text-gray-500 dark:text-white"/>
      </button>
      <div id="tooltip-bottom-2" role="tooltip" class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium bg-gray-300 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700 dark:text-white">
        Sign Out
        <div class="tooltip-arrow" data-popper-arrow></div>
      </div>
    </div>
  </div>
</template>