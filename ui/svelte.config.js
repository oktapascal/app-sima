import {vitePreprocess} from "@sveltejs/vite-plugin-svelte";

export default {
    // Consult https://svelte.dev/docs#compile-time-svelte-preprocess
    // for more information about preprocessors
    preprocess: vitePreprocess(),
    onwarn: (warning, handler) => {
        const {code, frame} = warning;
        if (code === "css-unused-selector")
            return;

        handler(warning);
    },
};
