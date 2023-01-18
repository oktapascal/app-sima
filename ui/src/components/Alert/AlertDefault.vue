<script lang="ts" setup>
import {computed, onBeforeUnmount, onMounted, ref} from "vue";
import {useBreakPointStore} from "@/stores/breakpoint";
import {useAlertStore} from "@/stores/alert";
import {IconAlert, IconAlertBox, IconAlertCircle, IconCheckCircle, ButtonAlert} from "@/components";

const breakPointStore = useBreakPointStore();
const alertStore = useAlertStore();

let time = ref<null | number>(null);

const position = computed(() => {
  return breakPointStore.checkIsMobile ? "top-center" : "bottom-right";
});

const icon = {
  success: IconCheckCircle,
  error: IconAlert,
  warning: IconAlertCircle,
  info: IconAlertBox,
};

const text = {
  success: "Success Message",
  error: "Error Message",
  warning: "Warning Message",
  info: "Information",
};

function onClose() {
  alertStore.$reset();
}

onMounted(() => {
  time.value = setTimeout(() => {
    alertStore.$reset();
  }, 5000);
});

onBeforeUnmount(() => {
  clearTimeout(time.value!);
});
</script>

<template>
  <div role="alert" class="alert p-4 mb-4 text-white border rounded-lg fixed w-auto"
       :class="[breakPointStore.checkIsMobile ? 'top-center' : 'bottom-right', alertStore.getAlert.type]">
    <div class="flex items-center">
      <component :is="icon[alertStore.getAlert.type]" class="mr-2 w-6 h-6 text-white"/>
      <h3 class="text-lg font-medium text-white">{{ text[alertStore.getAlert.type] }}</h3>
    </div>
    <div class="mt-2 mb-4 text-sm text-white">
      {{ alertStore.getAlert.text }}
    </div>
    <div class="flex">
      <ButtonAlert label="Tutup" @onClick="onClose"/>
    </div>
  </div>
</template>
