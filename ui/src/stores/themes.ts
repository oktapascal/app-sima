import {defineStore} from "pinia";

export const useThemeStore = defineStore("themes", {
    state: () => ({
        mode: window.localStorage.getItem("theme") || "light",
    }),
    getters: {
        isDark: (state) => state.mode !== "light",
        getMode: (state) => state.mode,
    },
    actions: {
        toggleMode(value: string) {
            this.mode = value;
            window.localStorage.setItem("theme", value);
        },
    },
});
