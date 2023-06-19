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
const router = useRouter()

const { getAlbum, removeSong } = useAlbumsService()
const { addAlbum, state: { albums } } = useLibraryService()

onMounted(async () => {
    const res = await getAlbum(route.params.album as any)

    album.value = res
})

const handleAddAlbum = async () => {
    if (!album.value) return

    await addAlbum(album.value.id)

    router.push("/library")
}
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
                        <div @click="handleAddAlbum"
                            :class="{ 'bg-secondary border-secondary text-primary': albums.find((p: Album) => p.id === album?.id) }"
                            class="rounded-full border-2 border-white w-[50px] h-[50px] flex justify-center items-center ml-5 mt-2">
                            <i class="uil uil-bookmark text-3xl"></i>
                        </div>
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