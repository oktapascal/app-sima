import axios from "axios"
import routes from "@/router"
import {useAlertStore, type IAlert} from "@/stores/alert";

const BASE_URL = "http://localhost:8080/api"

const instance = axios.create({
    withCredentials: true,
    baseURL: BASE_URL
})

instance.interceptors.request.use(function (config) {
    config.withCredentials = true

    return config
}, function (error) {
    return Promise.reject(error)
})

instance.interceptors.response.use(function (response) {
    return response
}, async function (error) {
    const alertStore = useAlertStore()
    if(error.response.status === 401) {
        await routes.push({
            path: "/login",
            query: {
                redirect: routes.currentRoute.value.path
            }
        })
    }

    if(error.response.status === 404) {
        const alert: IAlert = {
            show: true,
            type: "error",
            text: error.response.data?.statusMessage
        }

        alertStore.showAlert(alert)
    }

    if (error.response?.status === 500) {
        const alert: IAlert = {
            show: true,
            type: "error",
            text: "Terjadi kesalahan pada server"
        }

        alertStore.showAlert(alert)
    }

    return Promise.reject(error)
})
export default instance