<script lang="ts">
    import {useNavigate} from "svelte-navigator";
    import {AxiosError, type AxiosResponse} from "axios";
    import type {IAuth, IProfileResponse, IAlert} from "@/types";
    import instance from "@/libs/instance";
    import {auth} from "@/stores/authStore";
    import {alertStore} from "@/stores/alertStore";
    import {onMount} from "svelte";

    const navigate = useNavigate();

    let isLoading: boolean = true;

    async function checkCredentials() {
        try {
            const response: AxiosResponse<IProfileResponse> = await instance.get("/auth/profile");

            const authState: IAuth = {
                id_location: response.data.data.id_location,
                role: response.data.data.role,
                nik: response.data.data.nik,
                isAuthenticated: true,
                photo: "default.png",
            };

            auth.set(authState);
        } catch (error: AxiosError) {
            if (!$alertStore.show) {
                const alertState: IAlert = {
                    type: "error",
                    text: error.response.statusText,
                    show: true,
                };

                alertStore.set(alertState);
            }
        } finally {
            isLoading = false;
        }
    }

    onMount(() => {
        async function load() {
            await checkCredentials();
        }

        load();
    });

    $: {
        if ($auth.isAuthenticated && !isLoading) {
            navigate("/dashboard", {replace: true});
        } else if (!$auth.isAuthenticated && !isLoading) {
            navigate("/login", {replace: true});
        }
    }
</script>

<div>
    redirecting to...
</div>