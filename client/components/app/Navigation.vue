<template>
    <ul class="text-white flex items-center justify-end h-full">
        <template v-if="user">
            <app-nav-link>Home</app-nav-link>
            <app-nav-link>Playlists</app-nav-link>
            <app-nav-link>Artists</app-nav-link>
        </template>

        <template v-else>
            <app-nav-link to="/signIn">Sign In</app-nav-link>
        </template>
    </ul>
</template>

<script lang="ts" setup>
import { useUser } from '~~/hooks/user';

const user = useUser()
</script>