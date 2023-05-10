<script lang="ts" setup>
import { Playlist } from '~~/models/playlist';

defineProps<{
    playlist: Playlist,
    info: boolean,
}>()

const router = useRouter()
</script>

<template>
    <div class="w-full p-3 text-white px-5" :class="{
        'grid grid-cols-3 justify-items-center items-center': info,
        'flex justify-between items-center': !info
    }">
        <div class="flex justify-self-start">
            <img :src="playlist ? 'http://localhost:8080/api/file/' + playlist?.logo_id : ''"
                class="w-[50px] h-[50px] bg-white mr-5 cursor-pointer" @click="router.push('/playlists/detail/' + playlist.id)" />
            <div>
                <h2 @click="router.push('/playlists/detail/' + playlist.id)" class="hover:text-primary cursor-pointer">
                    {{ playlist?.name }}
                </h2>
                <h3 class="text-sm text-[#dcdcdc] hover:text-primary cursor-pointer">
                    {{ playlist?.author.name }}
                </h3>
            </div>
        </div>
        <span v-if="info">
            October 5, 2022
        </span>
        <span v-if="info" class="justify-self-end">
            10 Songs
        </span>
        <slot name="action">

        </slot>
    </div>
</template>