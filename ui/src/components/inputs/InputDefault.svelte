<script lang="ts">
    export let type: "text" | "number" | "email" = "text";
    export let placeholder: string;
    export let name: string;
    export let label: string;
    export let readonly: boolean = false;
    export let isError: boolean = false;
    export let className: string = "";

    let props;

    $: {
        const {type, placeholder, name, label, readonly, isError, className, ...restprops} = $$props;
        props = restprops;
    }
</script>

<label for="{name}" class="label block mb-2 text-sm font-medium text-gray-500 dark:text-white"
       class:error="{isError}">{label}</label>
<input class="input border-2 border-gray-300 text-black placeholder-gray-300 text-sm rounded-lg block w-full p-2.5 focus:ring-gray-200 dark:text-gray-50 dark:bg-gray-500 dark:placeholder-gray-300 dark:border-gray-600 {className}"
       class:error="{isError}"
       id="{name}"
       {type}
       {placeholder}
       {name}
       {readonly}
       {...props}
>