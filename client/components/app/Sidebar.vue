<template>
    <div class="col-span-2 w-full h-full p-8 text-white text-opacity-60">
        <ul class="flex flex-col h-[150px] justify-around">
            <app-side-nav-item text="Home" to="/" />
            <app-side-nav-item text="Search" icon="uil uil-search" />
            <app-side-nav-item text="Your Library" to="/" icon="uil uil-books" />
        </ul>
        <div class="mt-12">
            <div class="border-b border-light pb-5">
                <div class="flex items-center hover:text-white transition-all duration-200 cursor-pointer">
                    <i
                        class="uil uil-plus text-xl bg-gray-300 text-black h-[25px] w-[25px] flex justify-center items-center rounded-sm"></i>
                    <h3 class="text-sm ml-5">
                        Create Playlist
                    </h3>
                </div>
            </div>
            <ul class="flex flex-col pt-2">
                <li class="py-2" v-for="playlist in playlists">
                    <NuxtLink :to="'/playlists/detail/' + playlist.id">
                        <h3 class="hover:text-white cursor-pointer">
                            {{ playlist.name }}
                        </h3>
                    </NuxtLink>
                </li>

            </ul>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { Ref } from 'vue'
import { Playlist } from '~~/models/playlist'

const { getPlaylistsByAuthor } = usePlaylistsService()
const user = useUser()

const playlists: Ref<Playlist[]> = ref([])

const fetchData = async (page: number) => {
    playlists.value = await getPlaylistsByAuthor(user.value?.id as number)
}

onMounted(async () => {
    await fetchData(1)
})
</script>