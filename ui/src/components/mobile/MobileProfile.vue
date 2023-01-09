<script lang="ts" setup>
const emits = defineEmits(["onShowProfileBox"])

function onShowProfileBox() {
  emits("onShowProfileBox", true)
}
</script>

<template>
  <button type="button" class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200" @click="onShowProfileBox">
    <img alt="avatar" src="/images/avatars/avatar.png" class="w-7 h-7 rounded-full" />
  </button>
</template>
