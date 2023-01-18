<script lang="ts" setup>
import {withDefaults, toRefs} from "vue";

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

const {label, name, placeholder, readonly, type} = toRefs<Props>(props);

</script>

<template>
  <VField v-slot="{field, meta}" :name="name">
    <label :class="{
                'error': !meta.valid && meta.touched
            }" :for="name" class="label block mb-2 text-sm font-medium dark:text-white">
      {{ label }}
    </label>
    <input autocomplete="off"
           class="input border-2 border-gray-300 text-black placeholder-gray-300 text-sm rounded-lg block w-full p-2.5 focus:ring-gray-200 dark:text-gray-50 dark:bg-gray-500 dark:placeholder-gray-300 dark:border-gray-600"
           :type="type" :placeholder="placeholder" :readonly="readonly" :class="{
                    'error': !meta.valid && meta.touched,
                    'bg-gray-100 cursor-not-allowed': readonly
                }" v-bind="field"/>
  </VField>
  <VErrorMessage as="p" class="mt-1 pl-1 text-sm text-red-600 font-medium dark:text-red-500" :name="name"/>
</template>