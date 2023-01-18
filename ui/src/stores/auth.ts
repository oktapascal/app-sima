import axios, {AxiosError, type AxiosResponse} from "axios";
import type {IAlert, IAuth, IAuthResponse} from "@/types";
import {useAlertStore} from "./alert";
import {defineStore} from "pinia";

export const useAuthStore = defineStore("auth", {
    state: function () {
        return {
            id_user: null,
            kode_lokasi: null,
            role: null,
            isAuthenticated: false,
        };
    },
    getters: {
        getUserSession: function (state: IAuth) {
            return state;
        },
        getStatusAuthenticated: function (state: IAuth) {
            return state.isAuthenticated;
        },
    },
    actions: {
        setUserSession: function (data: IAuth) {
            this.id_user = data.id_user;
            this.kode_lokasi = data.kode_lokasi;
            this.role = data.role;
            this.isAuthenticated = data.isAuthenticated;
        },
        setUserSessionFromCredentials: async function () {
            const alertStore = useAlertStore();
            const self = this;

            try {
                const response: AxiosResponse<IAuthResponse> = await axios.get("/api/auth/session/user-access", {
                    withCredentials: true,
                });

                const data: IAuth = {
                    id_user: response.data.data.id_user,
                    role: response.data.data.role,
                    kode_lokasi: response.data.data.kode_lokasi,
                    isAuthenticated: true,
                };

                self.setUserSession(data);
            } catch (error: unknown) {
                if (error instanceof AxiosError) {
                    if (error.response?.status === 500) {
                        const alert: IAlert = {
                            show: true,
                            type: "error",
                            text: error.response.data?.statusMessage,
                        };

                        alertStore.showAlert(alert);
                    }
                }
            }
        },
    },
});