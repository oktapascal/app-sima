<script lang="ts" setup>
import {defineProps, withDefaults, defineEmits} from "vue";

interface Props {
  type: string;
  name: string;
  placeholder: string;
  label: string;
  readonly?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  readonly: false,
});

const emits = defineEmits(['onClickIcon'])

function onClickIcon(): void {
  emits('onClickIcon')
}
</script>

<template>
  <div class="flex flex-col w-full mb-6">
    <VField v-slot="{field, meta}" :name="props.name">
      <label :class="{
                'error': !meta.valid && meta.touched
            }" :for="props.name" class="input block mb-2 text-sm font-medium dark:text-white">
        {{ props.label }}
      </label>
      <div class="relative">
        <input autocomplete="off" class="bg-gray-50 border border-gray-500 text-black placeholder-gray-300 text-sm rounded-lg block w-full p-2.5 focus:ring-gray-200 dark:text-gray-50 dark:bg-gray-500 dark:placeholder-gray-300 dark:border-gray-600" :type="props.type" :placeholder="props.placeholder" :readonly="props.readonly" :class="{
                    'error': !meta.valid && meta.touched
                }" v-bind="field" />
        <span class="absolute right-3 top-2" @click="onClickIcon">
          <slot :isError="!meta.valid && meta.touched"/>
        </span>
      </div>
    </VField>
    <VErrorMessage as="p" class="mt-2 text-sm text-red-600 dark:text-red-500" :name="props.name"/>
  </div>
</template>