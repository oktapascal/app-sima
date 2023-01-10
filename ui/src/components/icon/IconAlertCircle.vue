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
            d="M13 13h-2V7h2m0 10h-2v-2h2M12 2A10 10 0 0 0 2 12a10 10 0 0 0 10 10a10 10 0 0 0 10-10A10 10 0 0 0 12 2Z" />
    </svg>
</template>