import {writable} from "svelte/store";
import type {IAlert} from "@/types";

export const alertStore = writable<IAlert>({
    text: "",
    show: false,
    type: "info",
});