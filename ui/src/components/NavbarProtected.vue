<script lang="ts" setup>
import {ref} from "vue";
import {useBreakPointStore} from "@/stores/breakpoint";
import {WebNavbar, IconMenu, IconMenuOpen, MobileNavbar} from "@/components";

const breakPointStore = useBreakPointStore();

const menuOpen = ref<boolean>(false);

function toggleMenu() {
  menuOpen.value = !menuOpen.value;
}
</script>

<template>
  <div class="flex-1">
    <div class="flex flex-wrap">
      <div>
        <button type="button"
                class="border border-gray-300 rounded-lg p-2.5 mr-2 lg:mr-0 focus:outline-none focus:ring-2 focus:ring-gray-400"
                @click="toggleMenu">
          <IconMenuOpen v-if="menuOpen" className="w-7 h-7 fill-current dark:text-white"/>
          <IconMenu v-else className="w-7 h-7 fill-current dark:text-white"/>
        </button>
      </div>
    </div>
  </div>
  <div class="flex-none">
    <MobileNavbar v-if="breakPointStore.checkIsMobile"/>
    <WebNavbar v-else/>
  </div>
</template>