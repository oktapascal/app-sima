<script lang="ts">
    import {createEventDispatcher} from "svelte";
    import {auth} from "@/stores/authStore";

    const dispatch = createEventDispatcher();

    function onShowProfileBox() {
        dispatch("show", {
            value: true,
        });
    }
</script>

<div class="mx-0.5">
    <button type="button"
            class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200"
            on:click={onShowProfileBox}>
        <img alt="avatar" src="{import.meta.env.VITE_API_URL_AVATAR}/{$auth.photo}"
             class="w-7 h-7 rounded-full"/>
    </button>
</div>