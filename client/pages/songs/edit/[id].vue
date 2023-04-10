<script lang="ts" setup>
definePageMeta({
    middleware: ['auth']
})

const { createSong, uploadSongMedia, updateSong, getSong } = useSongsService()

const router = useRouter()

const route = useRoute()

const songData = ref(null)

const submitEdit = async ({ song, media, logo, touched }: any) => {
    const updatedSong = await updateSong(song)

    if (updatedSong && touched) {
        const formData = new FormData()

        formData.append("songMedia", media)
        formData.append("songLogo", logo)

        await uploadSongMedia(song.id, formData)

        router.push("/songs")
    }

    router.push("/dashboard")
}

onMounted(async () => {
    const song = await getSong(route.params.id)

    console.log(song)

    songData.value = song
})

</script>

<template>
    <songs-create-form v-if="songData" :song-data="songData" :edit="true" @submit-edit="submitEdit"></songs-create-form>
</template>

<style></style>