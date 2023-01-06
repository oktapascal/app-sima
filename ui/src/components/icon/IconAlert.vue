<script lang="ts" setup>
interface Props {
  className: string
}

const props = defineProps<Props>()
</script>

<template>
    <svg :class="props.className" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid meet"
        viewBox="0 0 24 24">
        <path fill="currentColor" d="M13 14h-2V9h2m0 9h-2v-2h2M1 21h22L12 2L1 21Z" />
    </svg>
</template>