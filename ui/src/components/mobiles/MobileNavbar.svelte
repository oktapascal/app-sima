<script lang="ts">
    import Portal from "svelte-portal";
    import {Notification, MobileProfile, MobileProfileBox} from "@/components";

    let isOpenBox: boolean = false;

    function onToggleProfileBox(event) {
        isOpenBox = event.detail.value;
    }
</script>

<div class="flex flex-row flex-nowrap mr-2">
    <div class="flex-1">
        <Notification/>
    </div>
    <div class="flex-none">
        <MobileProfile on:show={onToggleProfileBox}/>
    </div>
</div>

<Portal to="body">
    {#if isOpenBox}
        <MobileProfileBox on:hide={onToggleProfileBox}/>
    {/if}
</Portal>