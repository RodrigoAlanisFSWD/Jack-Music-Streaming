<template>
    <ul class="text-white flex items-center justify-end h-full">
        <template v-if="status === AuthStatus.AUTHENTICATED">
            <app-nav-link>Home</app-nav-link>
            <app-nav-link to="/search">Search</app-nav-link>
            <app-nav-link to="/playlists">Playlists</app-nav-link>
        </template>

        <template v-else-if="status !== AuthStatus.AUTH_INITIAL">
            <app-nav-link to="/auth/signIn">Sign In</app-nav-link>
        </template>
    </ul>
</template>

<script lang="ts" setup>
import { AuthStatus } from '~~/store/authStore';

const authStore = useAuthStore()

const { status } = storeToRefs(authStore)
</script>