<script lang="ts" setup>
import { Ref } from 'vue'
import { Album } from '~/models/album';
import { Song } from '~~/models/song'

const { show, album } = defineProps<{
    show: boolean,
    album: Album | null
}>()

const { getAlbum, addSong } = useAlbumsService()
const { searchSongs, getSongs } = useSongsService()
const router = useRouter()

const songs: Ref<Song[]> = ref([])

const fetchData = async (page: number) => {
    const res = await getSongs(1)

    songs.value = res
}

onMounted(async () => {
    await fetchData(1)
})

const emit = defineEmits(['hide', 'addSong'])

const handleClickOutside = () => {
    if (show) {
        emit('hide', false)
    }
}

const handleAddSong = async (song: Song, album: Album) => {
    await addSong(song.id, album?.id as any)
    emit('addSong', song)
}

const getSongsFiltered = computed(() => {
    return songs.value.filter((song: Song) => song.album_id != album?.id)
})
</script>

<template>
    <Teleport to="body">
        <Transition enter-from-class="opacity-0" enter-active-class="transition duration-300" leave-to-class="opacity-0"
            leave-active-class="transition duration-300">
            <div v-if="show" v-click-outside="handleClickOutside"
                class="grid grid-cols-[400px_1fr] w-screen h-screen top-0 left-0 absolute">

                <div
                    class="col-start-2 col-end-3 w-full h-full bg-[#111] bg-opacity-75 z-50 flex justify-center items-center">

                    <div class="w-[450px] h-[350px] bg-secondary rounded-lg">
                        <div class="flex justify-between items-center pl-5 text-white">
                            <h2>
                                Add To Album
                            </h2>
                            <i @click="emit('hide')" class="uil uil-times text-2xl p-2"></i>
                        </div>
                        <div class="p-5 pt-0">
                            <app-search class="w-full"></app-search>
                        </div>
                        <div>
                            <div class="p-3 justify-between px-5 w-full flex text-white items-center"
                                v-for="song in getSongsFiltered" :key="song.id">
                                <div class="flex justify-self-start">
                                    <img :src="song ? 'http://localhost:8080/api/file/' + song?.logo_id : ''"
                                        class="w-[50px] h-[50px] bg-white mr-5 cursor-pointer" @click="handleAddSong(song, album)" />
                                    <div>
                                        <h2 class="hover:text-primary cursor-pointer w-[175px] overflow-hidden text-ellipsis whitespace-nowrap"
                                            @click="handleAddSong(song, album)">
                                            {{ song?.name }}
                                        </h2>
                                        <h3 class="text-sm text-[#dcdcdc] hover:text-primary cursor-pointer">
                                            {{ song?.author.name }}
                                        </h3>
                                    </div>
                                </div>
                                <span class="justify-self-end pr-2">
                                    {{ song?.duration }}
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </Transition>
    </Teleport>
</template>