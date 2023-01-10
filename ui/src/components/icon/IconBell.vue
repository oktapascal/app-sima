<script lang="ts" setup>
// Declare Props interface with a single property "className" of type string
interface Props {
  className: string
}

// Initialize props variable with defineProps function and specify the type as the Props interface
const props = defineProps<Props>()
</script>

<template>
  <svg  :class="props.className" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24">
    <path fill="currentColor" d="M21 19v1H3v-1l2-2v-6c0-3.1 2.03-5.83 5-6.71V4a2 2 0 0 1 2-2a2 2 0 0 1 2 2v.29c2.97.88 5 3.61 5 6.71v6l2 2m-7 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2"/>
  </svg>
</template>