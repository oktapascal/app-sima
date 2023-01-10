<script lang="ts" setup>

// Props interface defines the types and default values for the component's props
interface Props {
  disabled?: boolean // optional prop that defaults to false
  fullWidth?: boolean // optional prop that defaults to false
  type: "button" | "submit" | "reset" // required prop with a limited set of values
  label: string // required prop with a string value
  className: string // required prop with a string value
  disabledClassName: string // required prop with a string value
}

// withDefaults function adds default values to the Props interface
const props = withDefaults(defineProps<Props>(), {
  disabled: false,
  fullWidth: false
})

// defineEmits function creates an object that maps prop names to their corresponding event names
const emits = defineEmits(["onClick"])

// onClick function is called when the component's "onClick" event is emitted
function onClick() {
  emits("onClick")
}

</script>

<template>
  <div class="flex flex-col w-full py-4">
    <button :disabled="props.disabled" :type="props.type" :class="[ props.disabled ? props.disabledClassName : '', props.className]" @click="onClick">{{ props.label }}</button>
  </div>
</template>