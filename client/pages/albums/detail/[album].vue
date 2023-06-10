<script lang="ts" setup>
import { remove } from '@vue/shared'
import { Ref } from 'vue'
import { Album } from '~/models/album';
import { Song } from '~~/models/song'

definePageMeta({
    layout: "main",
    middleware: "auth"
})

const album: Ref<Album | null> = ref(null)

const route = useRoute()

const { getAlbum, removeSong } = useAlbumsService()

onMounted(async () => {
    const res = await getAlbum(route.params.album as any)

    album.value = res
})
</script>

<template>
    <div class="w-full h-full bg-[linear-gradient(90deg,#B18CFF_0%,#515ada_100%)] grid grid-rows-[300px_1fr]">
        <div class="p-12 flex text-white">
            <img :src="album ? 'http://localhost:8080/api/file/' + album?.logo_id : ''"
                class="w-[200px] h-[200px] bg-white mr-5 cursor-pointer shadow-2xl" />
            <div class="flex flex-col justify-between ml-8">
                <div>
                    <h3 class="mb-4">
                        Album
                    </h3>
                    <div class="flex items-center">
                        <h2 class="text-7xl font-bold">
                            {{ album?.name }}
                        </h2>
                    </div>

                </div>
                <div class="text-xl font-bold">
                    {{ album?.author.name }}
                </div>
            </div>

        </div>
        <albums-song-list class="bg-[#111] bg-opacity-75" v-if="album" :songs="album?.songs">
        </albums-song-list>
    </div>
</template>