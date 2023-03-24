<script lang="ts" setup>
definePageMeta({
    middleware: ['auth']
})

const { createSong, uploadSongMedia } = useSongsService()

const router = useRouter()

const submit = async ({ name, duration, media, logo }: any) => {
    const formData = new FormData()

    formData.append("songMedia", media)
    formData.append("songLogo", logo)

    const song = await createSong({
        name,
        duration,
    })

    if (song) {
        await uploadSongMedia(song.id, formData)

        router.push("/songs")
    }
}
</script>

<template>
    <songs-create-form @submit="submit"></songs-create-form>
</template>

<style></style>