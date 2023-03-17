<script lang="ts" setup>
import { Ref } from 'vue';
import { Song } from '../../models/song'

definePageMeta({
    layout: "player"
})

const { getSongs } = useSongsService()

const { setSong } = usePlayer()

const songs: Ref<Song[]> = ref([])

const fetchData = async (page: number) => {
    songs.value = await getSongs(page)
}

onMounted(async () => {
    await fetchData(1)
})
</script>

<template>
    <div class="w-full h-[calc(100vh-125px)]">
        <app-pagination @change-page="fetchData">
            <div class="w-full h-full flex flex-col">
                <div class="w-full p-3 text-white flex justify-between" v-for="song in songs" :key="song.id">
                    <div class="flex items-center">
                        <h3 class="text-lg">
                            {{ song.name }}
                        </h3>
                        <h4 class="text-sm ml-5">
                            By: {{ song.author.name }}
                        </h4>
                    </div>
                    <div>
                        <span>
                            Duration: 00:00
                        </span>
                        <i class="uil uil-play text-xl hover:text-primary transition-all duration-200 cursor-pointer mx-5" @click="setSong(song)"></i>
                    </div>
                </div>
            </div>
        </app-pagination>
    </div>
</template>