<script lang="ts" setup>
import {withDefaults, defineProps} from "vue";
interface Props {
  disabled?: boolean
  fullWidth?: boolean
  type: "button" | "submit" | "reset"
  label: string
  className: string
  disabledClassName: string
}

const props = withDefaults(defineProps<Props>(), {
  disabled: false,
  fullWidth: false
})
</script>

<template>
  <div class="flex flex-col w-full py-4">
    <button :disabled="props.disabled" :type="props.type" :class="[ props.disabled ? props.disabledClassName : '', props.className]">{{ props.label }}</button>
  </div>
</template>