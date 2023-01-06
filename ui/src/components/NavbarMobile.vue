<script lang="ts" setup>
import Notification from "@/components/Notification.vue";
import ProfileMobile from "@/components/ProfileMobile.vue";

</script>

<template>
  <div class="flex flex-row flex-nowrap mr-2">
    <div class="flex-1">
      <Notification />
    </div>
    <div class="flex-none">
      <ProfileMobile />
    </div>
  </div>
</template>