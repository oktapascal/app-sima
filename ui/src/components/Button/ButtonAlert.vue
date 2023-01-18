<script lang="ts" setup>
import {computed, toRefs} from "vue";
import {useAlertStore} from "@/stores/alert";

interface Props {
  label: string;
}

const alertStore = useAlertStore();

const props = defineProps<Props>();

const {label} = toRefs<Props>(props);

const emits = defineEmits(["onClick"]);

const classBtn = {
  success: "success-btn",
  error: "error-btn",
  warning: "warning-btn",
  info: "default-btn",
};

function onClick() {
  emits("onClick");
}
</script>

<template>
  <button type="button"
          class="btn-alert ml-auto text-white border font-medium rounded-lg text-xs px-3 py-1.5 text-center focus:ring-2 focus:outline-none"
          aria-label="Close" :class="classBtn[alertStore.getAlert.type]" @click="onClick">
    {{ label }}
  </button>
</template>