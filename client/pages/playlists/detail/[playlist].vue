<script lang="ts" setup>
import { Ref } from 'vue'
import { Playlist } from '~~/models/playlist'
import { Song } from '~~/models/song'

definePageMeta({
    layout: "main",
    middleware: "auth"
})

const playlist: Ref<Playlist | null> = ref(null)

const route = useRoute()
const router = useRouter()

const { getPlaylist, removeSong } = usePlaylistsService()
const { setPlaylist, setSong } = usePlayer()
const { addPlaylist, state: { playlists }, removePlaylist } = useLibraryService()
 
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

const handleRemove = async (songID: number) => {
    if (playlist.value) {
        await removeSong(songID, playlist.value?.id)

        playlist.value.songs = playlist.value.songs.filter((song: Song) => song.id != songID)
    }
}

const handleAddPlaylist = async () => {
    if (!playlist.value) return

    await addPlaylist(playlist.value.id)
    
    router.push("/library")      
}

const handleRemovePlaylist = async () => {
    if (!playlist.value) return

    await removePlaylist(playlist.value.id)    
}

const handleAction = async () => {
    if (!playlist.value) return

    if (playlists.value.find((p: Playlist) => p.id === playlist.value?.id)) {
        handleRemovePlaylist()
    } else {
        handleAddPlaylist()
    }
}
</script>

<template>
    <div class="w-full h-full bg-[linear-gradient(90deg,#B18CFF_0%,#515ada_100%)] grid grid-rows-[300px_1fr]">
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
                        <div @click="handleAction" :class="{'bg-secondary border-secondary text-primary': playlists.find((p: Playlist) => p.id === playlist?.id)}" class="rounded-full border-2 border-white w-[50px] h-[50px] flex justify-center items-center ml-5 mt-2">
                            <i class="uil uil-bookmark text-3xl"></i>
                        </div>
                    </div>

                </div>
                <div class="text-xl font-bold">
                    {{ playlist?.author.name }}
                </div>
            </div>

        </div>
        <app-song-list @remove="handleRemove" class="bg-[#111] bg-opacity-75" playlist v-if="playlist" :songs="playlist?.songs">
        </app-song-list>
    </div>
</template>