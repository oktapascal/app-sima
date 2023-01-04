import { defineStore } from "pinia"

export const useThemeStore = defineStore('themes', {
    state: function () {
        return {
            mode: window.localStorage.getItem("theme") || "light"
        }
    },
    getters: {
        isDark: function (state): boolean {
            return state.mode !== "light"
        },
        getMode: function (state): string {
            return state.mode
        }
    },
    actions: {
        toggleMode(value: string) {
            this.mode = value
            window.localStorage.setItem("theme", value)
        }
    }
})