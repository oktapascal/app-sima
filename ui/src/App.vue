<script lang="ts" setup>
// Import the onBeforeUnmount, onMounted, and useBreakPointStore, useAlertStore, and useAuthStore hooks from the vue and @/stores modules
import {onBeforeUnmount, onMounted} from "vue";
import {useBreakPointStore} from "@/stores/breakpoint";
import {useAlertStore} from "@/stores/alert";
import {useAuthStore} from "@/stores/auth";

// Import the NavbarBase, NavbarGuest, NavbarProtected, and AlertDefault components from the @/components module
import NavbarBase from "@/components/NavbarBase.vue";
import NavbarGuest from "@/components/NavbarGuest.vue";
import NavbarProtected from "@/components/NavbarProtected.vue";
import AlertDefault from "@/components/alert/AlertDefault.vue";

// Get the breakpoint store from the useBreakPointStore hook
const breakPointStore = useBreakPointStore();

// Get the alert store from the useAlertStore hook
const alertStore = useAlertStore()

// Get the auth store from the useAuthStore hook
const authStore = useAuthStore()

// Define a function to handle window resizes
function onResize() {
  // Get the width of the screen
  const width = window.screen.width;

  // Check if the width is between 320 and 768 (inclusive)
  if (320 <= width && 768 >= width) {
    // If it is, update the checkMobile state in the breakpoint store to true
    breakPointStore.updateCheckMobile(true);
  } else {
    // If it is not, update the checkMobile state in the breakpoint store to false
    breakPointStore.updateCheckMobile(false);
  }
}

// Use the onMounted hook to define a function that will run when the component is mounted
onMounted(async () => {
  // Check if the "theme" item is present in local storage
  if (localStorage.getItem("theme") == null) {
    // If it is not present, remove the "dark" class from the documentElement object
    document.documentElement.classList.remove("dark");
  } else {
    // If it is present, check the value of the "theme" item
    if (localStorage.getItem("theme") === "light") {
      // If the value is "light", remove the "dark" class from the documentElement object
      document.documentElement.classList.remove("dark");
    } else {
      // If the value is not "light", add the "dark" class to the documentElement object
      document.documentElement.classList.add("dark");
    }
  }

  // Get the width of the screen
  const width = window.screen.width;

  // Update the checkMobile state in the breakpoint store based on the width of the screen
  if (320 <= width && 768 >= width) {
    breakPointStore.updateCheckMobile(true);
  } else {
    breakPointStore.updateCheckMobile(false);
  }

  // Add an event listener for the "resize" event on the window object and call the onResize function as the event handler
  window.addEventListener("resize", onResize);
});

// Use the onBeforeUnmount hook to define a function that will run before the component is unmounted
onBeforeUnmount(() => {
  // Remove the event listener for the "resize" event on the window object
  window.removeEventListener("resize", onResize);
});
</script>

<template>
  <NavbarBase>
    <NavbarProtected />
<!--    <NavbarProtected v-if="authStore.getStatusAuthenticated" />-->
<!--    <NavbarGuest v-else />-->
  </NavbarBase>
  <main>
    <RouterView/>
  </main>
  <Teleport to="body">
    <Transition enter-active-class="animate__animated animate__fadeInRight" leave-active-class="animate__animated animate__fadeOutRight">
      <AlertDefault v-if="alertStore.getAlert.show"/>
    </Transition>
  </Teleport>
</template>
