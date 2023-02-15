<script lang="ts">
    import {fly} from "svelte/transition";
    import {AxiosError} from "axios";
    import {useNavigate, Link} from "svelte-navigator";
    import type {IAlert, IAuth} from "@/types";
    import {alertStore} from "@/stores/alertStore";
    import {auth} from "@/stores/authStore";
    import instance from "@/libs/instance";
    import {IconChevronDown, IconChevronUp, IconLogout, IconCog} from "@/components";

    let openBottomMenu: boolean = false;

    function toggleBottomMenu() {
        openBottomMenu = !openBottomMenu;
    }

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

<div class="group fixed left-0 top-0 z-50 h-full transition w-20 bg-gray-50 shadow-xl border-r-3xl flex flex-col transition-width ease hover:w-81">
    <div class="flex justify-start items-center p-1.5">
        <img alt="logo" src="/images/sima.png" class="h-10"/>
    </div>
    <div class="pb-8 h-full scroll-auto"></div>
    <div class="flex flex-col justify-center pb-6">
        <button type="button"
                class="flex flex-row items-center cursor-pointer pt-8 pb-4 px-5 text-gray-500 dark:text-white"
                on:click={toggleBottomMenu}>
            <img alt="avatar" src="{import.meta.env.VITE_API_URL_AVATAR}/{$auth.photo}"
                 class="border-2 border-solid border-transparent rounded-full h-11 w-11"/>
            <span class="transition-all whitespace-nowrap flex flex-col font-semibold ml-1.5 invisible text-gray-500 group-hover:visible dark:text-white">{$auth.nik}</span>
            {#if openBottomMenu}
                <IconChevronUp
                        className="h-6 w-6 ml-auto text-gray-500 whitespace-nowrap invisible group-hover:visible dark:text-white"/>
            {:else }
                <IconChevronDown
                        className="h-6 w-6 ml-auto text-gray-500 whitespace-nowrap invisible group-hover:visible dark:text-white"/>
            {/if}
        </button>
        {#if openBottomMenu}
            <ul class="m-0 outline-0 overflow-hidden list-none py-0 px-6" in:fly={{ y: 500, duration: 120 }}
                out:fly={{ y: 90, duration: 150, delay: 50 }}>
                <li class="rounded-lg transition-all cursor-pointer mt-1 select-none hover:bg-gray-200">
                    <button type="button"
                            class="flex flex-row items-center py-0.5 pr-4 rounded-lg border-8 border-transparent"
                            on:click={onSignOut}>
                        <IconLogout className="h-5 w-5 mr-3 text-gray-500 dark:text-white"/>
                        <span class="whitespace-nowrap text-gray-500 invisible dark:text-white group-hover:visible">Sign Out</span>
                    </button>
                </li>
                <li class="rounded-lg transition-all cursor-pointer mt-1 select-none hover:bg-gray-200">
                    <button type="button"
                            class="flex flex-row items-center py-0.5 pr-4 rounded-lg border-8 border-transparent">
                        <IconCog className="h-5 w-5 mr-3 text-gray-500 dark:text-white"/>
                        <span class="whitespace-nowrap text-gray-500 invisible dark:text-white group-hover:visible">Settings</span>
                    </button>
                </li>
                <li class="rounded-lg transition-all cursor-pointer mt-1 select-none hover:bg-gray-200">
                    <Link to="/profile"
                          class="flex h-full w-full flex-row items-center py-0.5 pr-4 rounded-lg border-8 border-transparent">
                        <i class="icofont-user text-lg mr-3 text-gray-500 dark:text-white"></i>
                        <span class="whitespace-nowrap text-gray-500 invisible dark:text-white group-hover:visible">Profile</span>
                    </Link>
                </li>
            </ul>
        {/if}
    </div>
</div>