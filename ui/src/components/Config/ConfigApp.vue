<script lang="ts" setup>
import {ref, watch} from "vue";
import {IconCog, ConfigBox} from "@/components";
import {useThemeStore} from "@/stores/themes";

const themeStore = useThemeStore();

const openConfig = ref<boolean>(false);
const checked = ref(themeStore.isDark);

function onToggle(value: boolean) {
  openConfig.value = value;
}

watch(checked, () => {
  if (checked.value) {
    themeStore.toggleMode("dark");
    document.documentElement.classList.add("dark");
  } else {
    themeStore.toggleMode("light");
    document.documentElement.classList.remove("dark");
  }
});
</script>

<template>
  <div>
    <button data-tooltip-target="tooltip-bottom-3" data-tooltip-placement="bottom"
            class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200 dark:hover:bg-gray-600"
            @click="onToggle(true)">
      <IconCog className="h-7 w-7 fill-current text-gray-500 dark:text-white"/>
    </button>
    <div id="tooltip-bottom-3" role="tooltip"
         class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium bg-gray-300/60 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700/60 dark:text-white">
      Config App
      <div class="tooltip-arrow" data-popper-arrow></div>
    </div>
    <Transition enter-active-class="animate__animated animate__slideInRight"
                leave-active-class="animate__animated animate__slideOutRight">
      <ConfigBox v-if="openConfig" @onToggle="onToggle">
        <template #default>
          <div class="flex flex-col">
            <div class="px-4 py-3">
              <h5 class="text-lg font-semibold dark:text-white my-2.5">Theme App</h5>
              <div>
                <label class="label cursor-pointer justify-start radio-control my-2.5">
                  <input type="radio" name="radio-theme" class="radio-custom" v-model="checked" :value="false"/>
                  <span class="label-text font-semibold text-sm ml-2 dark:text-white">Light</span>
                </label>
                <label class="label cursor-pointer justify-start radio-control my-2.5">
                  <input type="radio" name="radio-theme" class="radio-custom" v-model="checked" :value="true"/>
                  <span class="label-text font-semibold text-sm ml-2 dark:text-white">Dark</span>
                </label>
              </div>
            </div>
            <hr class="border-t border-gray-400"/>
          </div>
        </template>
      </ConfigBox>
    </Transition>
  </div>
</template>