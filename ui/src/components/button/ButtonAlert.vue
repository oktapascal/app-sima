<script lang="ts" setup>
import {computed} from "vue"
import {useAlertStore} from "@/stores/alert";

// Props interface for the ButtonAlert component
interface Props {
  label: string
}

// Use the alert store
const alertStore = useAlertStore()

// Define the props for the component
const props = defineProps<Props>()

// Define the events emitted by the component
const emits = defineEmits(["onClick"])

// Create a computed property for the button class
// based on the type of alert (success, error, warning, or info)
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

// Emit an onClick event when the button is clicked
function onClick() {
  emits("onClick")
}

</script>

<template>
  <button type="button" class="btn-alert ml-auto text-white border font-medium rounded-lg text-xs px-3 py-1.5 text-center focus:ring-2 focus:outline-none" aria-label="Close" :class="buttonTypeClass" @click="onClick">
    {{ props.label }}
  </button>
</template>