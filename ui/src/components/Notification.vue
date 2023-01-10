<script lang="ts" setup>
import { withDefaults } from "vue"; // Import withDefaults function from vue
import IconBell from "@/components/icon/IconBell.vue"; // Import IconBell component

// Define the Props interface with a single property, total, of type number
interface Props {
  total?: number
}

// Initialize props with withDefaults function, which takes the output of defineProps function and an object of default values
const props = withDefaults(defineProps<Props>(), {
  total: 0 // Set the default value of total to 0
})
</script>

<template>
  <div class="mx-1 relative">
    <button data-tooltip-target="tooltip-bottom-1" data-tooltip-placement="bottom" class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200">
      <IconBell className="h-7 w-7 fill-current text-gray-500 dark:text-white" />
    </button>
    <span class="badge" v-if="props.total > 0">{{ props.total }}</span>
    <div id="tooltip-bottom-1" role="tooltip" class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium bg-gray-300 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700 dark:text-white">
      Notification
      <div class="tooltip-arrow" data-popper-arrow></div>
    </div>
  </div>
</template>

<style scoped>
.badge {
  @apply absolute text-white -right-0.5 -top-1 rounded-lg bg-blue-700 py-0.5 px-1 dark:bg-gray-800;
  font-size: 0.6875rem;
  line-height: 0.625rem;
}
</style>