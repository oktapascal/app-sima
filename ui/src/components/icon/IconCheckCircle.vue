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
            d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m-2 15l-5-5l1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9Z" />
    </svg>
</template>