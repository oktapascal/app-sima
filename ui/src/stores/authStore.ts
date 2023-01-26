import {writable} from "svelte/store";
import type {IAuth} from "@/types";

export const auth = writable<IAuth>({
    nik: null,
    kode_lokasi: null,
    role: null,
    isAuthenticated: false,
});