<script lang="ts">
    import {AxiosError} from "axios";
    import {useNavigate} from "svelte-navigator";
    import type {IAlert, IAuth} from "@/types";
    import {alert} from "@/stores/alertStore";
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
                kode_lokasi: null,
            };

            auth.set(authState);

            const alertState: IAlert = {
                show: true,
                text: "Kamu berhasil logout",
                type: "success",
            };

            alert.set(alertState);

            navigate("/login", {replace: true});
        } catch (error: AxiosError) {
            const alertState: IAlert = {
                show: true,
                text: "Terjadi kesalahan pada server",
                type: "error",
            };

            alert.set(alertState);
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