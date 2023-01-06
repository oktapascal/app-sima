<script lang="ts" setup>
import {ref, watch} from "vue"
import IconCog from "@/components/icon/IconCog.vue";
import ConfigBox from "@/components/config/ConfigBox.vue";
import {useThemeStore} from "@/stores/themes";

const themeStore = useThemeStore();

const openConfig = ref<boolean>(false)
const checked = ref(themeStore.isDark)

function onToggle(value: boolean) {
  openConfig.value = value
}

watch(checked, () => {
  if (checked.value) {
    themeStore.toggleMode("dark");
    document.documentElement.classList.add("dark");
  } else {
    themeStore.toggleMode("light");
    document.documentElement.classList.remove("dark");
  }
})
</script>

<template>
  <div>
    <button class="p-2.5 border border-blue-700 bg-blue-700 fixed top-2/4 right-0 truncate -mt-6 text-center z-50 rounded-l-lg rounded-r-none border-0 hover:bg-blue-700 dark:bg-gray-700 dark:hover:bg-gray-700" @click="onToggle(true)">
      <IconCog className="h-7 w-7 text-white dark:text-white" />
    </button>
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