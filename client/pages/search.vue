<script lang="ts" setup>
import { Ref } from 'vue';
import { Song } from '../../models/song'

definePageMeta({
    layout: "main"
})

const { getSongs } = useSongsService()

const { setSong, setPlaylist } = usePlayer()

const songs: Ref<Song[]> = ref([])

const fetchData = async (page: number) => {
    const res = await getSongs(page)
    setPlaylist(res)
    songs.value = res
}

onMounted(async () => {
    await fetchData(1)
})

const user = useUser()
</script>

<template>
    <div class="w-full h-full">
        <app-pagination @change-page="fetchData">
            <div class="flex justify-between h-[75px] items-center px-5">
                <app-search class="w-[400px] h-[45px]"></app-search>
                <div class="flex text-white items-center bg-black p-2 rounded-full w-[200px] justify-between">
                    <div class="flex items-center">
                        <img v-if="user?.profile" :src="user?.profile ? 'http://localhost:8080/api/file/' + user?.profile.logo_id : ''" class="w-[30px] h-[30px] bg-white rounded-full" />
                        <h3 class="ml-5">
                            {{ user?.name }}
                        </h3>
                    </div>

                    <i class="uil uil-angle-down text-2xl"></i>
                </div>
            </div>
            <div class="w-full h-full flex flex-col">
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
                <app-song v-for="song in songs" :key="song.id" :song="song"></app-song>
            </div>
        </app-pagination>
    </div>
</template>