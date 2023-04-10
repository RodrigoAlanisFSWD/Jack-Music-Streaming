<template>
    <div class="h-[70px] grid grid-cols-[1fr_400px_1fr] items-center w-screen">
        <header class="flex">
            <div>
                <NuxtLink to="/">
                    <div class="text-white flex items-end px-5">
                        <h1 class="text-2xl">
                            Jack
                        </h1>
                        <h4 class="text-primary text-sm ml-1">
                            ms
                        </h4>
                    </div>
                </NuxtLink>
            </div>

            <div class="text-xl text-white flex w-[150px] justify-around items-center ml-2">
                <NuxtLink :to="role?.name !== 'ARTIST' ? '/profile' : '/dashboard'">
                    <i class="uil uil-user"></i>
                </NuxtLink>
                <i class="uil uil-envelope"></i>
                <i class="uil uil-setting"></i>
            </div>
        </header>

        <app-search class="w-[400px]"></app-search>

        <app-navigation class="h-full"></app-navigation>
    </div>
</template>

<script lang="ts" setup>
const authStore = useAuthStore()

const { user, role } = storeToRefs(authStore)
</script>