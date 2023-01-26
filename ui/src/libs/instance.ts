import axios from "axios";

const BASE_URL = import.meta.env.VITE_BASE_URL;

const instance = axios.create({
    withCredentials: true,
    baseURL: BASE_URL,
});

instance.interceptors.request.use(function (config) {
    config.withCredentials = true;

    return config;
}, function (error) {
    return Promise.reject(error);
});

instance.interceptors.response.use(
    function (response) {
        return response;
    },
    function (error) {
        // const navigate = useNavigate()
        //
        // if (error.response.status === 500) {
        //     const alertState: IAlert = {
        //         text: "Terjadi kesalahan pada server",
        //         type: "error",
        //         show: true,
        //     };
        //
        //     alert.set(alertState);
        // }
        //
        // if (error.response.status === 404) {
        //     const alertState: IAlert = {
        //         text: error.response.data.statusMessage,
        //         type: "error",
        //         show: true,
        //     };
        //
        //     alert.set(alertState);
        // }
        //
        // if (error.response.status === 401) {
        //     navigate("/login", {replace: true})
        // }

        return Promise.reject(error);
    },
);

export default instance;