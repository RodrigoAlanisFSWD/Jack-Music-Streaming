<script lang="ts" setup>
import { Ref } from 'vue'
import { Playlist } from '~~/models/playlist'
import { Song } from '~~/models/song'
import { watch } from 'vue'

const { show, selectedSong } = defineProps<{
    show: boolean,
    selectedSong: number
}>()

const { getPlaylistsByAuthor, addSong } = usePlaylistsService()
const user = useUser()
const router = useRouter()

const playlists: Ref<Playlist[]> = ref([])

const fetchData = async (page: number) => {
    if (user.value) {
        const res = await getPlaylistsByAuthor(user.value?.id as number)

        playlists.value = res
    }
}

onMounted(async () => {
    await fetchData(1)
})

const emit = defineEmits(['hide'])

const handleClickOutside = () => {
    if (show) {
        emit('hide', false)
    }
}

const handleSelectPlaylist = async (playlist: Playlist, song: any) => {
    await addSong(song, playlist.id)

    router.push("/playlists/detail/" + playlist.id)
}

const getPlaylists = computed(() => {
    return playlists.value.filter((playlist: Playlist) => {
        return !playlist.songs.find((song: Song) => song.id == selectedSong)
    })
})
</script>

<template>
    <Teleport to="body">
        <Transition enter-from-class="opacity-0" enter-active-class="transition duration-300" leave-to-class="opacity-0"
            leave-active-class="transition duration-300">
            <div v-if="show" v-click-outside="handleClickOutside"
                class="grid grid-cols-8 w-screen h-[calc(100vh-80px)] top-0 left-0 absolute">

                <div
                    class="col-start-3 col-end-9 w-full h-full bg-[#111] bg-opacity-75 z-50 flex justify-center items-center">

                    <div class="w-[450px] h-[350px] bg-secondary rounded-lg">
                        <div class="flex justify-between items-center pl-5 text-white">
                            <h2>
                                Add To Playlist
                            </h2>
                            <i @click="emit('hide')" class="uil uil-times text-2xl p-2"></i>
                        </div>
                        <div class="p-5 pt-0">
                            <app-search class="w-full"></app-search>
                        </div>
                        <div v-for="playlist in getPlaylists" :key="playlist.id">
                            <app-playlist :info="false" :playlist="playlist" 
                                 v-if="!playlist.songs.find((song: Song) => song.id == selectedSong)">
                                <template #action>
                                    <app-button @click="handleSelectPlaylist(playlist, selectedSong)"
                                        class="w-[44px] h-[45px]">
                                        <i class="uil uil-plus"></i>
                                    </app-button>
                                </template>
                            </app-playlist>
                        </div>
                    </div>
                </div>
            </div>
        </Transition>
    </Teleport>
</template>