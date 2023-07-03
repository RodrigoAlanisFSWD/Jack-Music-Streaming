<template>
    <div class="flex flex-col w-full h-full max-h-[calc(100vh-80px)] text-white overflow-y-scroll">
        <div class="w-full p-5">
            <h2 v-if="songs.length > 0" class="text-xl my-5">
                Songs
            </h2>
            <div class="flex gap-x-8">
                <app-card @detail="setSong(song)" @remove="removeSong(song.id)" v-for="song in songs" :name="song.name" :author="song.author" :logo="song ? 'http://localhost:8080/api/file/' + song.logo_id : ''"></app-card>
            </div>

        </div>
        <div class="w-full p-5">
            <h2 v-if="playlists.length > 0" class="text-xl my-5">
                Playlists
            </h2>
            <div class="flex gap-x-8">
                <app-card @detail="router.push('/playlists/detail/' + playlist.id)" @remove="removePlaylist(playlist.id)" v-for="playlist in playlists" :name="playlist.name" :author="playlist.author" :logo="playlist ? 'http://localhost:8080/api/file/' + playlist.logo_id : ''"></app-card>
            </div>
        </div>
        <div class="w-full p-5">
            <h2 v-if="albums.length > 0" class="text-xl my-5">
                Albums
            </h2>
            <div class="flex gap-x-8">
                <app-card @detail="router.push('/albums/detail/' + detail.id)" @remove="removeAlbum(album.id)" v-for="album in albums" :name="album.name" :author="album.author" :logo="album ? 'http://localhost:8080/api/file/' + album.logo_id : ''"></app-card>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
definePageMeta({
    layout: "main",
    middleware: ["auth"]
})

const { state: { songs, albums, playlists }, removeAlbum, removePlaylist, removeSong } = useLibraryService()
const { setSong } = usePlayer()

const router = useRouter()
</script>