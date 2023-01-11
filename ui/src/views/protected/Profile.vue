<script lang="ts" setup>
import IconCamera from "@/components/icon/IconCamera.vue";  // Import the camera icon component
import InputDefault from "@/components/input/InputDefault.vue"; // Import the default input field component
import ButtonDefault from "@/components/button/ButtonDefault.vue"; // Import the default button component
</script>

<template>
  <div class="mt-2.5 h-full w-full bg-white shadow-md lg:rounded-lg lg:shadow-lg lg:border lg:border-gray-200 dark:bg-gray-700">
    <div class="p-2.5">
      <h6 class="text-lg font-bold text-gray-500 font-bold dark:text-white">Profile</h6>
    </div>
    <hr class="border-t border-gray-500" />
    <div class="flex flex-col flex-wrap p-2.5  lg:flex-row">
      <div class="flex justify-center lg:basis-1/2">
        <div class="relative">
          <div class="rounded-full border-2 border-gray-500">
            <img alt="profile-image" src="/images/avatars/avatar.png" class="rounded-full h-28 w-28" />
          </div>
          <input type="file" class="hidden" ref="uploadInput" />
          <button type="button" class="absolute top-20 right-0 p-2 rounded-full bg-gray-300/60 focus:outline-none focus:ring-2 focus:ring-gray-500 dark:bg-gray-600/60">
            <IconCamera className="h-5 w-5 fill-current dark:text-white" />
          </button>
        </div>
      </div>
      <div class="flex flex-col flex-1 my-4 lg:my-0">
        <form>
          <div class="my-2.5">
            <InputDefault name="nik" placeholder="NIK" label="NIK" type="text" />
          </div>
          <div class="my-2.5">
            <InputDefault name="nama" placeholder="Nama" label="Nama" type="text" />
          </div>
          <div class="my-2.5">
            <InputDefault name="email" placeholder="Email" label="Email" type="email" />
          </div>
          <div class="my-2.5">
            <InputDefault name="no_telp" placeholder="Nomor Telepon" label="Nomor Telepon" type="text" />
          </div>
          <div class="my-2.5">
            <InputDefault name="alamat" placeholder="Alamat" label="Alamat" type="text" />
          </div>
          <div class="mt-6">
            <ButtonDefault label="Simpan" type="submit"
             className="w-full bg-blue-700 text-white hover:bg-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 focus:outline-none focus:ring-2 focus:ring-blue-300 dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus-gray-500"
             disabledClassName="cursor-not-allowed bg-blue-500 dark:bg-gray-500"
            />
          </div>
        </form>
      </div>
    </div>
  </div>
</template>