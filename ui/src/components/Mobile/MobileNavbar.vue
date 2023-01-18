<script lang="ts" setup>
import {ref} from "vue";
import {Notification, MobileProfile, MobileProfileBox} from "@/components";

const openProfileBox = ref<boolean>(false);

function onToggleProfileBox(value: boolean) {
  openProfileBox.value = value;
}
</script>

<template>
  <div class="flex flex-row flex-nowrap mr-2">
    <div class="flex-1">
      <Notification/>
    </div>
    <div class="flex-none">
      <MobileProfile @onShowProfileBox="onToggleProfileBox"/>
    </div>
  </div>
  <Teleport to="body">
    <Transition enter-active-class="animate__animated animate__fadeInUp"
                leave-active-class="animate__animated animate__fadeOutDown">
      <MobileProfileBox v-if="openProfileBox" @onCloseProfileBox="onToggleProfileBox"/>
    </Transition>
  </Teleport>
</template>