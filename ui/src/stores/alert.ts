import {defineStore} from "pinia";
import type {IAlert} from "@/types";

export const useAlertStore = defineStore("alert", {
    state: function (): IAlert {
        return {
            text: "",
            type: "info",
            show: false,
        };
    },
    getters: {
        getAlert: function (state): IAlert {
            return state;
        },
    },
    actions: {
        showAlert: function (data: IAlert) {
            this.text = data.text;
            this.type = data.type;
            this.show = data.show;
        },
    },
});