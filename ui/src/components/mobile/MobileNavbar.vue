<script lang="ts" setup>
import { ref } from "vue"; // Import ref function from vue
import Notification from "@/components/Notification.vue"; // Import Notification component
import ProfileMobile from "@/components/mobile/MobileProfile.vue"; // Import MobileProfile component
import MobileProfileBox from "@/components/mobile/MobileProfileBox.vue"; // Import MobileProfileBox component

// Initialize openProfileBox variable as a ref with a default value of false
const openProfileBox = ref<boolean>(false)

// Declare onToggleProfileBox function that takes a value of type boolean
function onToggleProfileBox(value: boolean) {
  // Set the value of openProfileBox to the value passed to the function
  openProfileBox.value = value
}
</script>

<template>
  <div class="flex flex-row flex-nowrap mr-2">
    <div class="flex-1">
      <Notification/>
    </div>
    <div class="flex-none">
      <ProfileMobile @onShowProfileBox="onToggleProfileBox"/>
    </div>
  </div>
  <Teleport to="body">
    <Transition enter-active-class="animate__animated animate__fadeInUp" leave-active-class="animate__animated animate__fadeOutDown">
      <MobileProfileBox v-if="openProfileBox" @onCloseProfileBox="onToggleProfileBox"/>
    </Transition>
  </Teleport>
</template>