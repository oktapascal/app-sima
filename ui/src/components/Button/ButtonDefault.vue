<script lang="ts" setup>
import {toRefs} from "vue";

interface Props {
  disabled?: boolean;
  type: "button" | "submit" | "reset";
  label: string;
  className: string;
  disabledClassName: string;
}

const props = withDefaults(defineProps<Props>(), {
  disabled: false,
});

const {className, disabled, disabledClassName, label, type} = toRefs<Props>(props);

const emits = defineEmits(["onClick"]);

function onClick() {
  emits("onClick");
}
</script>

<template>
  <button :disabled="disabled" :type="type" :class="[ disabled ? disabledClassName : '', className]" @click="onClick">
    {{ label }}
  </button>
</template>