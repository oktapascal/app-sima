<script lang="ts">
    import {auth} from "@/stores/authStore";
</script>

<div class="fixed left-0 top-0 z-50 h-full transition w-20 bg-gray-50 shadow-xl border-r-3xl flex flex-col">
    <div class="flex justify-start items-center p-1.5">
        <img alt="logo" src="/images/sima.png" class="h-10"/>
    </div>
    <div class="pb-8 h-full scroll-auto"></div>
    <div class="flex flex-col justify-center pb-6">
        <button type="button"
                class="flex flex-row items-center cursor-pointer pt-8 pb-4 px-5 text-gray-500 dark:text-white">
            <img alt="avatar" src="{import.meta.env.VITE_API_URL_AVATAR}/{$auth.photo}"
                 class="border-2 border-solid border-transparent rounded-full h-11 w-11"/>
        </button>
    </div>
</div>