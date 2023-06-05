<script lang="ts" setup>
import { remove } from '@vue/shared'
import { Ref } from 'vue'
import { Album } from '~/models/album';
import { Song } from '~~/models/song'

const album: Ref<Album | null> = ref(null)

const route = useRoute()

const { getAlbum, removeSong } = useAlbumsService()

onMounted(async () => {
    const res = await getAlbum(route.params.album as any)

    album.value = res
})

const handleRemove = async (songID: number) => {
    if (album.value) {
        await removeSong(songID, album.value?.id)

        album.value.songs = album.value.songs.filter((song: Song) => song.id != songID)
    }
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

                        <div class="flex justify-center items-center w-[50px] h-[50px] bg-secondary rounded-full ml-5 mt-2">
                            <i class="uil uil-plus text-2xl text-primary"></i>
                        </div>
                    </div>

                </div>
                <div class="text-xl font-bold">
                    {{ album?.author.name }}
                </div>
            </div>

        </div>
        <app-album-song-list @remove="handleRemove" class="bg-[#111] bg-opacity-75" v-if="album" :songs="album?.songs">
        </app-album-song-list>
    </div>
</template>