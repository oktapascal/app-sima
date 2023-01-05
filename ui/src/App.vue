<script lang="ts" setup>
import {onBeforeUnmount, onMounted} from "vue";
import {useBreakPointStore} from "@/stores/breakpoint";
import {useAlertStore} from "@/stores/alert";
import NavbarBase from "@/components/NavbarBase.vue";
import NavbarGuest from "@/components/NavbarGuest.vue";
import AlertDefault from "@/components/alert/AlertDefault.vue";

const breakPointStore = useBreakPointStore();
const alertStore = useAlertStore()

function onResize() {
  const width = window.screen.width;

  if (320 <= width && 768 >= width) {
    breakPointStore.updateCheckMobile(true);
  } else {
    breakPointStore.updateCheckMobile(false);
  }
}

onMounted(async () => {
  if (localStorage.getItem("theme") == null) {
    document.documentElement.classList.remove("dark");
  } else {
    if (localStorage.getItem("theme") === "light") {
      document.documentElement.classList.remove("dark");
    } else {
      document.documentElement.classList.add("dark");
    }
  }

  const width = window.screen.width;

  if (320 <= width && 768 >= width) {
    breakPointStore.updateCheckMobile(true);
  } else {
    breakPointStore.updateCheckMobile(false);
  }

  window.addEventListener("resize", onResize);
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", onResize);
});
</script>

<template>
  <NavbarBase>
    <NavbarGuest/>
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
