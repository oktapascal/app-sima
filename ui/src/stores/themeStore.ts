import {writable} from "svelte/store";

export const theme = writable<string>(window.localStorage.getItem("theme") || "light");