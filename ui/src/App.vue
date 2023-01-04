<script lang="ts" setup>
import {onBeforeUnmount, onMounted} from "vue";
import {useBreakPointStore} from "@/stores/breakpoint";
import NavbarBase from "@/components/NavbarBase.vue";
import NavbarGuest from "@/components/NavbarGuest.vue";

const breakPointStore = useBreakPointStore();

function onResize() {
  const width = window.screen.width;

  if (320 <= width && 768 >= width) {
    breakPointStore.updateCheckMobile(true);
  } else {
    breakPointStore.updateCheckMobile(false);
  }
}

onMounted(() => {
  if (localStorage.getItem("theme") == null) {
    document.documentElement.classList.remove("dark");
  } else {
    if (localStorage.getItem("theme") === "light") {
      document.documentElement.classList.remove("dark");
    } else {
      document.documentElement.classList.add("dark");
    }
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
</template>
