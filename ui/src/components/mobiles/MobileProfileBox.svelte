<script lang="ts">
    import {useNavigate} from "svelte-navigator";
    import {createEventDispatcher} from "svelte";
    import {fly} from "svelte/transition";
    import {theme} from "@/stores/themeStore";
    import {alertStore} from "@/stores/alertStore";
    import {auth} from "@/stores/authStore";
    import instance from "@/libs/instance";
    import type {IAlert, IAuth} from "@/types";
    import {IconWindowClose, IconMoon, IconSun, IconLogout} from "@/components";
    import {AxiosError} from "axios";

    const dispatch = createEventDispatcher();
    const navigate = useNavigate();

    function onHideProfileBox() {
        dispatch("hide", {
            value: false,
        });
    }

    function toggleTheme(value) {
        if (value === "light") {
            theme.set("dark");
            window.localStorage.setItem("theme", "dark");
            document.documentElement.classList.add("dark");
        } else {
            theme.set("light");
            window.localStorage.setItem("theme", "light");
            document.documentElement.classList.remove("dark");
        }
    }

    async function onSignOut() {
        try {
            await instance.post("/auth/logout");

            const authState: IAuth = {
                id_location: null,
                isAuthenticated: false,
                role: null,
                nik: null,
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

    $: textTheme = $theme === "dark" ? "Dark Mode" : "Light Mode";
</script>

<div transition:fly={{ duration: 500, delay: 100, y: 500 }}
     class="absolute top-0 h-screen w-full bg-gray-100 dark:bg-black z-50">
    <div class="relative flex flex-col">
        <div class="flex-none">
            <button class="float-right font-medium rounded-full p-2 inline-flex items-center focus:outline-none"
                    on:click={onHideProfileBox}>
                <IconWindowClose className="h-7 w-7 text-gray-500 dark:text-white"/>
            </button>
        </div>
        <hr class="border-t border-gray-500"/>
        <div class="flex-1">
            <div class="flex flex-col p-2.5">
                <div class="flex-none">
                    <h6 class="text-lg text-gray-500 font-bold dark:text-white">Profile</h6>
                </div>
                <div class="flex-1 mt-2">
                    <button type="button"
                            class="rounded-lg px-2 py-2.5 inline-flex items-center w-full hover:bg-gray-200 focus:ring-2 focus:outline-none focus:ring-gray-300 dark:hover:bg-gray-600 dark:focus:ring-gray-500">
                        <img alt="avatar" src="{import.meta.env.VITE_API_URL_AVATAR}/{$auth.photo}"
                             class="w-12 h-12 rounded-full"/>
                        <span class="text-gray-500 font-medium text-md text-center pl-4 dark:text-white">{$auth.nik}</span>
                    </button>
                </div>
            </div>
        </div>
        <hr class="border-t border-gray-500"/>
        <div class="flex-1">
            <div class="flex flex-col p-2.5">
                <div class="flex-none">
                    <h6 class="text-lg text-gray-500 font-bold dark:text-white">Settings</h6>
                </div>
                <div class="flex-1 mt-2">
                    <button type="button"
                            class="rounded-lg px-2 py-2.5 inline-flex items-center w-full hover:bg-gray-200 focus:ring-2 focus:outline-none focus:ring-gray-300 dark:hover:bg-gray-600 dark:focus:ring-gray-500"
                            on:click={() => toggleTheme($theme)}>
                        {#if $theme === "dark"}
                            <IconMoon className="h-7 w-7 text-gray-500 dark:text-white"/>
                        {:else }
                            <IconSun className="h-7 w-7 text-gray-500 dark:text-white"/>
                        {/if}
                        <span class="text-gray-500 font-medium text-md text-center pl-4 dark:text-white">{textTheme}</span>
                    </button>
                </div>
                <div class="flex-1 mt-2">
                    <button type="button"
                            class="rounded-lg px-2 py-2.5 inline-flex items-center w-full hover:bg-gray-200 focus:ring-2 focus:outline-none focus:ring-gray-300 dark:hover:bg-gray-600 dark:focus:ring-gray-500"
                            on:click={onSignOut}>
                        <IconLogout className="h-7 w-7 text-gray-500 dark:text-white"/>
                        <span class="text-gray-500 font-medium text-md text-center pl-4 dark:text-white">Sign Out</span>
                    </button>
                </div>
            </div>
        </div>
    </div>
</div>