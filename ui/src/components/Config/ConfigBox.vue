<script lang="ts" setup>
import {ref} from "vue";
import {onClickOutside} from "@vueuse/core";

const box = ref();

const emits = defineEmits(["onToggle"]);

onClickOutside(box, (event) => {
  emits("onToggle", false);
});
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
