<template>
    <div class="w-full h-[calc(100vh-75px)] flex justify-center items-center">
        <div class="w-1/4 h-2/4 border border-white rounded-md flex flex-col justify-center items-center text-white">
            <h2 class="text-xl">
                {{ user?.name }}
            </h2>
            <h3 class="mt-5">
                {{ user?.email }}
            </h3>
            <h3 class="mt-5">
                Rol: {{ user?.role?.name }}
            </h3>
        </div>
    </div>
</template>

<script lang="ts" setup>
    definePageMeta({
        middleware: 'auth'
    })

    const authStore = useAuthStore()

    const { user } = storeToRefs(authStore)
</script>