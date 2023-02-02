<script lang="ts">
    import {onMount} from "svelte";
    import {useNavigate} from "svelte-navigator";
    import {auth} from "@/stores/authStore";
    import {alert} from "@/stores/alertStore";
    import {AxiosError, type AxiosResponse} from "axios";
    import type {IAuth, IProfileResponse, IAlert} from "@/types";
    import instance from "@/libs/instance";

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
                photo: response.data.data.photo || "default.png",
            };

            auth.set(authState);
        } catch (error: AxiosError) {
            if (!$alert.show) {
                const alertState: IAlert = {
                    type: "error",
                    text: error.response.statusText,
                    show: true,
                };

                alert.set(alertState);
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

    $: if ($auth.isAuthenticated && !isLoading) {
        navigate("/dashboard", {replace: true});
    }
</script>

{#if !$auth.isAuthenticated}
    <slot/>
{/if}