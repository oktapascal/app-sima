<script lang="ts">
    import {AxiosError} from "axios";
    import {useNavigate} from "svelte-navigator";
    import type {IAlert, IAuth} from "@/types";
    import {alertStore} from "@/stores/alertStore";
    import {auth} from "@/stores/authStore";
    import instance from "@/libs/instance";
    import {Notification, Config, SignOut} from "@/components";

    const navigate = useNavigate();

    async function onSignOut() {
        try {
            await instance.post("/auth/logout");

            const authState: IAuth = {
                isAuthenticated: false,
                nik: null,
                role: null,
                id_location: null,
                photo: "default.png",
            };

            auth.set(authState);

            const alertState: IAlert = {
                show: true,
                text: "Kamu berhasil logout",
                type: "success",
            };

            alertStore.set(alertState);

            navigate("/login", {replace: true});
        } catch (error: AxiosError) {
            const alertState: IAlert = {
                show: true,
                text: "Terjadi kesalahan pada server",
                type: "error",
            };

            alertStore.set(alertState);
        }
    }
</script>

<div class="flex flex-row flex-nowrap mr-2">
    <div class="flex-1">
        <Notification/>
    </div>
    <div class="flex-1">
        <Config/>
    </div>
    <div class="flex-none">
        <SignOut on:action={onSignOut}/>
    </div>
</div>