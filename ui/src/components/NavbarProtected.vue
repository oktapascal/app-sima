<script lang="ts" setup>
import { ref } from "vue"; // Import ref function from vue
import { useBreakPointStore } from "@/stores/breakpoint"; // Import useBreakPointStore hook from breakpoint store
import IconMenu from "@/components/icon/IconMenu.vue"; // Import IconMenu component
import IconMenuOpen from "@/components/icon/IconMenuOpen.vue"; // Import IconMenuOpen component
import NavbarMobile from "@/components/mobile/MobileNavbar.vue"; // Import NavbarMobile component
import NavbarWeb from "@/components/web/WebNavbar.vue"; // Import NavbarWeb component

// Initialize breakPointStore hook with useBreakPointStore hook
const breakPointStore = useBreakPointStore()

// Initialize menuOpen as a ref with an initial value of false
const menuOpen = ref<boolean>(false)

// Declare toggleMenu function that does not take any arguments and does not return a value
function toggleMenu() {
  // Toggle the value of menuOpen between true and false
  menuOpen.value = !menuOpen.value
}
</script>

<template>
  <div class="flex-1">
    <div class="flex flex-wrap">
      <div>
        <button type="button" class="border border-gray-400 rounded-lg p-2.5 mr-2 lg:mr-0 focus:outline-none focus:ring-2 focus:ring-gray-400" @click="toggleMenu">
          <IconMenuOpen v-if="menuOpen" className="w-7 h-7 fill-current dark:text-white"/>
          <IconMenu v-else className="w-7 h-7 fill-current dark:text-white"/>
        </button>
      </div>
    </div>
  </div>
  <div class="flex-none">
    <NavbarMobile />
<!--    <NavbarMobile v-if="breakPointStore.checkIsMobile" />-->
<!--    <NavbarWeb v-else />-->
  </div>
</template>