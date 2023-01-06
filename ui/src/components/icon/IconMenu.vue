<script lang="ts" setup>
interface Props {
  className: string
}

const props = defineProps<Props>()
</script>

<template>
  <svg :class="props.className" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid meet"
       viewBox="0 0 24 24">
    <path fill="currentColor" d="M3 6h18v2H3V6m0 5h18v2H3v-2m0 5h18v2H3v-2Z" />
  </svg>
</template>