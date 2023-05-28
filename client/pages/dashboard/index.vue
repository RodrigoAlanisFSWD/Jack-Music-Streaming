<template>
    <div class="flex flex-col items-start text-white">
            <h2 class="text-2xl mb-5">
                Albums
            </h2>
            <app-album v-for="album in albums" :key="album.id" :album="album" info></app-album>
            <div @click="router.push('dashboard/otherSongs')" class="mt-5 flex items-center hover:text-primary transition-all duration-200 cursor-pointer">
                <h3 class="text-xl mr-3">
                    Other Songs
                </h3>
                <i class="uil uil-angle-down text-2xl"></i>
            </div>
            <app-button @click="router.push('/songs/create')" class="mt-8 w-[250px]">Upload Song</app-button>
        </div>
</template>

<script lang="ts" setup>
import { Album } from '~/models/album';

const user = useUser()

const router = useRouter()

const { getAlbums } = useAlbumsService()

const albums: Ref<Album[]> = ref([])

onMounted(async () => {
    if (user.value) {
        albums.value = await getAlbums(user.value.id)
    }
})

</script>