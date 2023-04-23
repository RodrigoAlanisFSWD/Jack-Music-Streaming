<script lang="ts" setup>
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
<div class="w-full h-[calc(100vh-180px)]">
        <app-pagination @change-page="fetchData">
            <div class="w-full h-full flex flex-col">
                <div class="grid grid-cols-4 justify-items-center items-center px-5 text-[#dcdcdc] border-b border-[#343434] py-2">
                    <span class="justify-self-start">
                        Title
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