import {writable} from "svelte/store";
import type {IAuth} from "@/types";

export const auth = writable<IAuth>({
    nik: null,
    id_location: null,
    role: null,
    isAuthenticated: false,
});