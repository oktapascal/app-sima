import { defineStore } from "pinia"

export const useBreakPointStore = defineStore("breakpoint", {
    state: function () {
        return {
            isMobile: false
        }
    },
    getters: {
        checkIsMobile: function (state) {
            return state.isMobile
        }
    },
    actions: {
        updateCheckMobile(value: boolean) {
            this.isMobile = value;
        },
    }
})