<script lang="ts">
    export let className: string = "";
</script>

<p class="mt-1 pl-1 text-sm text-red-600 font-medium dark:text-red-500 {className}">
    <slot/>
</p>