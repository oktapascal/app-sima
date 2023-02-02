<script lang="ts">
    import {onMount} from "svelte";
    import {AxiosError, type AxiosResponse} from "axios";
    import {useNavigate, useLocation, useFocus} from "svelte-navigator";
    import {auth} from "@/stores/authStore";
    import instance from "@/libs/instance";
    import type {IProfileResponse, IAuth, IAlert} from "@/types";
    import {alertStore} from "@/stores/alertStore";

    const navigate = useNavigate();
    const location = useLocation();
    const registerFocus = useFocus();

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
            const alertState: IAlert = {
                type: "error",
                text: error.response.statusText,
                show: true,
            };

            alertStore.set(alertState);
        } finally {
            isLoading = false;
        }
    }

    onMount(() => {
        if (!$auth.isAuthenticated) {
            async function load() {
                await checkCredentials();
            }

            load();
        }
    });

    $: if (!$auth.isAuthenticated && !isLoading) {
        navigate("/login", {
            state: {redirectTo: $location.pathname},
            replace: true,
        });
    }
</script>

{#if $auth.isAuthenticated}
    <slot {registerFocus}/>
{/if}