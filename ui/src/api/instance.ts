import axios from "axios"
import routes from "@/router"

const instance = axios.create()

instance.interceptors.request.use(function (config) {
    config.withCredentials = true

    return config
}, function (error) {
    return Promise.reject(error)
})

instance.interceptors.response.use(function (response) {
    return response
}, async function (error) {
    if(error.response.status === 401) {
        await routes.push({
            path: "/login",
            query: {
                redirect: routes.currentRoute.value.path
            }
        })
    }

    return Promise.reject(error)
})
export default instance