<script lang="ts" setup>
import { Ref } from 'vue'
import { Playlist } from '~~/models/playlist'

definePageMeta({
    layout: "player",
    middleware: "auth"
})

const { getPlaylists } = usePlaylistsService()

const playlists: Ref<Playlist[]> = ref([])

const fetchData = async (page: number) => {
    playlists.value = await getPlaylists(page)
}

onMounted(async () => {
    await fetchData(1)
})
</script>

<template>
<div class="w-full h-[calc(100vh-180px)]">
        <app-pagination @change-page="fetchData">
            <div class="w-full h-full flex flex-col">
                <div class="grid grid-cols-3 justify-items-center items-center px-5 text-[#dcdcdc] border-b border-[#343434] py-2">
                    <span class="justify-self-start">
                        Title
                    </span>
                    <span class="mr-9">
                        Date added
                    </span>
                    <span class="justify-self-end mr-6">
                        <i class="uil uil-clock"></i>
                    </span>
                </div>
                <app-playlist v-for="playlist in playlists" :key="playlist.id" :playlist="playlist"></app-playlist>
            </div>
        </app-pagination>
    </div>
</template>