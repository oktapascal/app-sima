<script lang="ts" setup>
// Declare Props interface with a single property "className" of type string
interface Props {
  className: string
}

// Initialize props variable with defineProps function and specify the type as the Props interface
const props = defineProps<Props>()
</script>

<template>
  <svg :class="props.className" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid meet"
       viewBox="0 0 24 24">
    <path fill="currentColor"
          d="M21 15.61L19.59 17l-5.01-5l5.01-5L21 8.39L17.44 12L21 15.61M3 6h13v2H3V6m0 7v-2h10v2H3m0 5v-2h13v2H3Z" />
  </svg>
</template>