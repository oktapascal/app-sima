<script lang="ts" setup>
import { ref } from "vue" // Import ref function from vue
import { onClickOutside } from "@vueuse/core" // Import onClickOutside function from vueuse/core

// Initialize box as a ref with no default value
const box = ref();

// Declare emits object with "onToggle" key
const emits = defineEmits(["onToggle"])

// Call onClickOutside function with arguments box and a callback function
onClickOutside(box, (event) => {
  // Emit "onToggle" event with false as the value
  emits("onToggle", false)
})
</script>

<template>
  <div class="config-app bg-gray-50 dark:bg-gray-700" ref="box">
    <slot/>
  </div>
</template>

<style scoped>
.config-app {
  @apply fixed h-full w-52 right-0 top-0;
  box-shadow: 0 2px 4px -1px rgb(0 0 0 / 20%), 0 4px 5px 0 rgb(0 0 0 / 14%), 0 1px 10px 0 rgb(0 0 0 / 12%);
  z-index: 99;
}
</style>
