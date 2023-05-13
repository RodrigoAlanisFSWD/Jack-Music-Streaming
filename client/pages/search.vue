<script lang="ts" setup>
import { Ref } from 'vue';
import { Song } from '../models/song'

definePageMeta({
    layout: "main",
    middleware: "auth"
})

const { getSongs, searchSongs, state: { songs } } = useSongsService()

const { setSong, setPlaylist } = usePlayer()

const fetchData = async (page: number) => {
    const res = await getSongs(page)
    setPlaylist(res)
}

onMounted(async () => {
    await fetchData(1)
})

const user = useUser()

// Search

const query = ref('')

const handleSearch = async () => {
    if (query.value == '') {
        await fetchData(1)

        return
    }
    await searchSongs(query.value)
}
</script>

<template>
    <div class="w-full h-full">
        <div class="flex justify-between h-[75px] items-center px-5">
            <app-search @search="handleSearch" v-model="query" class="w-[400px] h-[45px]"></app-search>
            <div class="flex text-white items-center bg-black p-2 rounded-full w-[200px] justify-between">
                <div class="flex items-center">
                    <img v-if="user?.profile"
                        :src="user?.profile ? 'http://localhost:8080/api/file/' + user?.profile.logo_id : ''"
                        class="w-[30px] h-[30px] bg-white rounded-full" />
                    <NuxtLink to="/profile">
                        <h3 class="ml-5">
                            {{ user?.name }}
                        </h3>
                    </NuxtLink>
                </div>

                <i class="uil uil-angle-down text-2xl"></i>
            </div>
        </div>
        <app-song-list :songs="songs"></app-song-list>
    </div>
</template>