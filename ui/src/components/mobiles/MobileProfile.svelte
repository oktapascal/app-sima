<script lang="ts">
    import {createEventDispatcher} from "svelte";
    import {auth} from "@/stores/authStore";

    const dispatch = createEventDispatcher();

    function onShowProfileBox() {
        dispatch("show", {
            value: true,
        });
    }

    console.log($auth.photo);
</script>

<div class="mx-0.5">
    <button type="button"
            class="p-2.5 text-center inline-flex items-center rounded-full hover:bg-gray-100 focus:ring-2 focus:outline-none focus:ring-gray-200"
            on:click={onShowProfileBox}>
        <img alt="avatar" src="http://192.168.1.10:8080/api/storage/avatars/{$auth.photo}"
             class="w-7 h-7 rounded-full"/>
    </button>
</div>