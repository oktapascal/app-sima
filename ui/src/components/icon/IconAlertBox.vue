<script lang="ts" setup>
interface Props {
  className: string
}

const props = defineProps<Props>()
</script>

<template>
    <svg :class="props.className" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid meet"
        viewBox="0 0 24 24">
        <path fill="currentColor"
            d="M5 3h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2m8 10V7h-2v6h2m0 4v-2h-2v2h2Z" />
    </svg>
</template>