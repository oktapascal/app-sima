<script lang="ts" setup>
import { withDefaults } from "vue"; // Import withDefaults function from vue

// Declare Props interface with five properties: type, name, placeholder, label, and readonly
interface Props {
  type: string;
  name: string;
  placeholder: string;
  label: string;
  readonly?: boolean;
}

// Initialize props variable with the result of calling withDefaults on defineProps function, and specify the type as the Props interface.
// Also pass an object with the key "readonly" set to false as the second argument to withDefaults
const props = withDefaults(defineProps<Props>(), {
  readonly: false,
});

// Declare emits object with "onClickIcon" key
const emits = defineEmits(['onClickIcon'])

// Declare onClickIcon function that does not take any arguments and does not return a value
function onClickIcon(): void {
  // Emit "onClickIcon" event
  emits('onClickIcon')
}
</script>

<template>
    <VField v-slot="{field, meta}" :name="props.name">
      <label :class="{
                'error': !meta.valid && meta.touched
            }" :for="props.name" class="label block mb-2 text-sm font-medium dark:text-white">
        {{ props.label }}
      </label>
      <div class="relative">
        <input autocomplete="off" class="input bg-gray-50 border border-gray-500 text-black placeholder-gray-300 text-sm rounded-lg block w-full p-2.5 focus:ring-gray-200 dark:text-gray-50 dark:bg-gray-500 dark:placeholder-gray-300 dark:border-gray-600" :type="props.type" :placeholder="props.placeholder" :readonly="props.readonly" :class="{
                    'error': !meta.valid && meta.touched
                }" v-bind="field" />
        <span class="absolute right-3 top-2" @click="onClickIcon">
          <slot :isError="!meta.valid && meta.touched"/>
        </span>
      </div>
    </VField>
    <VErrorMessage as="p" class="mt-1 pl-1 text-sm text-red-600 font-medium dark:text-red-500" :name="props.name"/>
</template>