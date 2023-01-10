<script lang="ts" setup>
// Declare emits object with "onShowProfileBox" key
const emits = defineEmits(["onShowProfileBox"])

// Declare onShowProfileBox function that does not take any arguments and does not return a value
function onShowProfileBox() {
  // Emit "onShowProfileBox" event with true as the value
  emits("onShowProfileBox", true)
}
</script>

<template>
  <button type="button" class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200" @click="onShowProfileBox">
    <img alt="avatar" src="/images/avatars/avatar.png" class="w-7 h-7 rounded-full" />
  </button>
</template>
