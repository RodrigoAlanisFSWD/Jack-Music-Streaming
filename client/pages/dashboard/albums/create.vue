<script lang="ts" setup>
const router = useRouter()

const { createAlbum, uploadAlbumLogo, } = useAlbumsService()

const handleSubmit = async (albumData: Object, logo: File) => {
    try {
        const album = await createAlbum(albumData)

        if (!album) return

        if (logo) {
            const formData = new FormData()

            formData.append("logo", logo)

            await uploadAlbumLogo(album.id, formData)
        }

        router.push("/dashboard/albums/detail/" + album.id)
    } catch (error) {
        console.log(error)
        return
    }
}
</script>

<template>
    <div class="w-full h-full flex justify-center items-center">
        <albums-form @submit="handleSubmit"></albums-form>
    </div>
</template>