<script lang="ts" setup>
import IconCamera from "@/components/icon/IconCamera.vue";  // Import the camera icon component
import InputDefault from "@/components/input/InputDefault.vue"; // Import the default input field component
import ButtonDefault from "@/components/button/ButtonDefault.vue"; // Import the default button component

</script>

<template>
  <div class="mt-2.5 flex flex-col h-full w-full bg-white shadow-md dark:bg-gray-700">
    <div class="flex-none p-2.5">
      <h6 class="text-lg font-bold text-gray-500 font-bold dark:text-white">Profile</h6>
    </div>
    <hr class="border-t border-gray-500" />
    <div class="flex-1 flex justify-center">
      <div class="relative p-2.5 my-2 rounded-full border border-gray-500">
        <img alt="profile-image" src="/images/avatars/avatar.png" class="rounded-full h-28 w-28" />
        <input type="file" class="hidden" ref="uploadInput" />
        <button type="button" class="absolute top-24 right-0 p-2 rounded-full bg-gray-300/60 focus:outline-none focus:ring-2 focus:ring-gray-500 dark:bg-gray-600/60">
          <IconCamera className="h-5 w-5 fill-current dark:text-white" />
        </button>
      </div>
    </div>
    <div class="flex-1 mt-2.5 p-2.5">
      <form>
        <div class="grid grid-cols-1 gap-4">
          <div>
            <InputDefault name="nik" placeholder="NIK" label="NIK" type="text" />
          </div>
          <div>
            <InputDefault name="nama" placeholder="Nama" label="Nama" type="text" />
          </div>
          <div>
            <InputDefault name="email" placeholder="Email" label="Email" type="email" />
          </div>
          <div>
            <InputDefault name="no_telp" placeholder="Nomor Telepon" label="Nomor Telepon" type="text" />
          </div>
          <div>
            <InputDefault name="alamat" placeholder="Alamat" label="Alamat" type="text" />
          </div>
          <div class="flex flex-col w-full py-4">
            <ButtonDefault label="Simpan" type="submit"
               className="bg-blue-700 text-white hover:bg-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 focus:outline-none focus:ring-2 focus:ring-blue-300 dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus-gray-500"
               disabledClassName="cursor-not-allowed bg-blue-500 dark:bg-gray-500" />
          </div>
        </div>
      </form>
    </div>
  </div>
</template>