<script lang="ts">
    export let type: "text" | "password" = "text";
    export let placeholder: string;
    export let name: string;
    export let label: string;
    export let readonly: boolean = false;
    export let isError: boolean = false;
    export let className: string = "";
</script>

<label for="{name}" class="label block mb-2 text-sm font-medium text-gray-500 dark:text-white"
       class:error="{isError}">{label}</label>
<div class="relative">
    <input class="input border-2 border-gray-300 text-black placeholder-gray-300 text-sm rounded-lg block w-full p-2.5 focus:ring-gray-200 dark:text-gray-50 dark:bg-gray-500 dark:placeholder-gray-300 dark:border-gray-600 {className}"
           class:error="{isError}"
           id="{name}"
           {type}
           {placeholder}
           {name}
           {readonly}
    >
    <span class="absolute right-3 top-2">
        <slot name="append"/>
    </span>
</div>