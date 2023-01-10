<script lang="ts" setup>
import {ref,computed,onMounted,onBeforeUnmount} from "vue"
import {useBreakPointStore} from "@/stores/breakpoint"
import {useAlertStore} from "@/stores/alert"
import IconAlert from "@/components/icon/IconAlert.vue";
import IconAlertBox from "@/components/icon/IconAlertBox.vue";
import IconAlertCircle from "@/components/icon/IconAlertCircle.vue";
import IconCheckCircle from "@/components/icon/IconCheckCircle.vue";
import ButtonAlert from "@/components/button/ButtonAlert.vue";

// Import and use the breakpoint and alert stores
const breakPointStore = useBreakPointStore()
const alertStore = useAlertStore()

// Create a ref to store a timeout ID
let time = ref<null|number>(null)

// Create a computed property for the alert position
// (top-center for mobile, bottom-right for desktop)
const position = computed(() => {
  return breakPointStore.checkIsMobile ? "top-center" : "bottom-right"
})

// Create a computed property for the alert text
// based on the type of alert (success, error, warning, or info)
const text = computed(() => {
  if(alertStore.getAlert.type === "success") {
    return "Success Message"
  }

  if(alertStore.getAlert.type === "error") {
    return "Error Message"
  }

  if(alertStore.getAlert.type === "warning") {
    return "Warning Message"
  }

  return "Information"
})

// Reset the alert store when the alert is closed
function onClose() {
  alertStore.$reset()
}

// When the component is mounted, set a timeout to reset the alert store
onMounted(() => {
  time.value = setTimeout(() => {
    alertStore.$reset()
  }, 5000)
})

// Clear the timeout when the component is unmounted
onBeforeUnmount(() => {
  clearTimeout(time.value!)
})
</script>

<template>
  <div role="alert" class="alert p-4 mb-4 text-white border rounded-lg fixed w-auto" :class="[position, alertStore.getAlert.type]">
    <div class="flex items-center">
      <template v-if="alertStore.getAlert.type === 'success'">
        <IconCheckCircle className="mr-2 w-6 h-6 text-white"/>
      </template>
      <template v-else-if="alertStore.getAlert.type === 'error'">
        <IconAlert className="mr-2 w-6 h-6 text-white"/>
      </template>
      <template v-else-if="alertStore.getAlert.type === 'warning'">
        <IconAlertCircle className="mr-2 w-6 h-6 text-white"/>
      </template>
      <template v-else>
        <IconAlertBox className="mr-2 w-6 h-6 text-white"/>
      </template>
      <h3 class="text-lg font-medium text-white">{{ text }}</h3>
    </div>
    <div class="mt-2 mb-4 text-sm text-white">
      {{ alertStore.getAlert.text }}
    </div>
    <div class="flex">
      <ButtonAlert label="Tutup" @onClick="onClose" />
    </div>
  </div>
</template>
