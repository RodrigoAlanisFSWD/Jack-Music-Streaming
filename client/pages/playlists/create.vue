<script lang="ts" setup>
import { sortUserPlugins } from 'vite'


definePageMeta({
    layout: "main",
    middleware: ["auth"]
})

const router = useRouter()

const { createPlaylist, uploadPlaylistLogo, playlistsStore } = usePlaylistsService()

const handleSubmit = async (playlistData: Object, logo: File) => {
    try {
        const playlist = await createPlaylist(playlistData)

        if (!playlist) return
        
        playlistsStore.addPlaylist(playlist)

        if (logo) {
            const formData = new FormData()

            formData.append("logo", logo)

            await uploadPlaylistLogo(playlist.id, formData)
        }

        router.push("/playlists/detail/" + playlist.id)
    } catch (error) {
        console.log(error)
        return
    }
}
</script>

<template>
    <div class="w-full h-full flex justify-center items-center">
        <playlists-form @submit="handleSubmit"></playlists-form>
    </div>
</template>