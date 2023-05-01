<script lang="ts" setup>
import { Ref } from 'vue'
import { Playlist } from '~~/models/playlist'

definePageMeta({
    layout: "main",
    middleware: "auth"
})

const playlist: Ref<Playlist | null> = ref(null)

const route = useRoute()

const { getPlaylist } = usePlaylistsService()
const { setPlaylist, setSong } = usePlayer()

onMounted(async () => {
    const res = await getPlaylist(route.params.playlist)

    playlist.value = res
})

const handlePlay = () => {
    if (playlist.value) {
        setPlaylist(playlist.value?.songs)
        setSong(playlist.value?.songs[0])
    }
}
</script>

<template>
    <div
        class="w-full h-full bg-[linear-gradient(90deg,#B18CFF_0%,#515ada_100%)] grid grid-rows-[300px_1fr]">
        <div class="p-12 flex text-white">
            <img :src="playlist ? 'http://localhost:8080/api/file/' + playlist?.logo_id : ''"
                class="w-[200px] h-[200px] bg-white mr-5 cursor-pointer shadow-2xl" />
            <div class="flex flex-col justify-between ml-8">
                <div>
                    <h3 class="mb-4">
                        Public Playlist
                    </h3>
                    <div class="flex items-center">
                        <h2 class="text-7xl font-bold">
                            {{ playlist?.name }}
                        </h2>
                        <div class="flex justify-center items-center w-[50px] h-[50px] bg-secondary rounded-full ml-5 mt-2">
                            <i class="uil uil-play text-2xl text-primary" @click="handlePlay"></i>
                        </div>
                    </div>

                </div>
                <div class="text-xl font-bold">
                    {{ playlist?.author.name }}
                </div>

            </div>

        </div>
        <div class="bg-black bg-opacity-75 h-full">
            <div
                class="grid grid-cols-4 justify-items-center items-center px-5 text-[#dcdcdc] border-b border-[#343434] py-2">
                <span class="justify-self-start">
                    Title
                </span>
                <span class="mr-11">
                    Album
                </span>
                <span class="mr-9">
                    Date added
                </span>
                <span class="justify-self-end mr-4">
                    <i class="uil uil-clock"></i>
                </span>
            </div>
            <app-song v-for="song in playlist?.songs" :key="song.id" :song="song"></app-song>
        </div>
    </div>
</template>