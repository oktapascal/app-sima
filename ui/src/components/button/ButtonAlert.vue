<script lang="ts" setup>
import {computed} from "vue"
import {useAlertStore} from "@/stores/alert";

interface Props {
  label: string
}

const alertStore = useAlertStore()

const props = defineProps<Props>()

const emits = defineEmits(["onClick"])

const buttonTypeClass = computed(() => {
  if(alertStore.getAlert.type === "success") {
    return "success-btn"
  }

  if(alertStore.getAlert.type === "error") {
    return "error-btn"
  }

  if(alertStore.getAlert.type === "warning") {
    return "warning-btn"
  }

  return "default-btn"
})

function onClick() {
  emits("onClick")
}
</script>

<template>
  <button type="button" class="btn-alert ml-auto text-white border font-medium rounded-lg text-xs px-3 py-1.5 text-center focus:ring-2 focus:outline-none" aria-label="Close" :class="buttonTypeClass" @click="onClick">
    {{ props.label }}
  </button>
</template>