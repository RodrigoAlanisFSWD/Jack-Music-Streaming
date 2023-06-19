<template>
    <div class="flex flex-col w-full h-full max-h-[calc(100vh-80px)] text-white overflow-y-scroll">
        <div v-if="songs.length > 0" class="w-full p-5">
            <h2 class="text-xl my-5">
                Songs
            </h2>
            <div class="flex">
                <div v-for="song in songs" :key="song.id" class="flex flex-col">
                    <img :src="song ? 'http://localhost:8080/api/file/' + song?.logo_id : ''"
                        class="w-[150px] h-[150px] bg-white mr-5 cursor-pointer rounded-sm mb-2" />
                    <h2 class="text-sm overflow-hidden max-w-[150px] text-ellipsis whitespace-nowrap">
                        {{ song.name }}
                    </h2>
                    <h3 class="text-[#dcdcdc] text-sm">
                        {{ song.author.name }}
                    </h3>
                </div>
            </div>

        </div>
        <div v-if="playlists.length > 0" class="w-full p-5">
            <h2 class="text-xl my-5">
                Playlists
            </h2>
            <div class="flex">

                <div v-for="playlist in playlists" :key="playlist.id" class="flex flex-col">
                    <img :src="playlist ? 'http://localhost:8080/api/file/' + playlist?.logo_id : ''"
                        class="w-[150px] h-[150px] bg-white mr-5 cursor-pointer rounded-sm mb-2" />
                    <h2 class="text-sm overflow-hidden max-w-[150px] text-ellipsis whitespace-nowrap">
                        {{ playlist.name }}
                    </h2>
                    <h3 class="text-[#dcdcdc] text-sm">
                        {{ playlist.author.name }}
                    </h3>
                </div>
            </div>
        </div>
        <div v-if="albums.length > 0" class="w-full p-5">
            <h2 class="text-xl my-5">
                Albums
            </h2>
            <div class="flex">
                <div v-for="album in albums" :key="album.id" class="flex flex-col">
                    <img :src="album ? 'http://localhost:8080/api/file/' + album?.logo_id : ''"
                        class="w-[150px] h-[150px] bg-white mr-5 cursor-pointer rounded-sm mb-2" />
                    <h2 class="text-sm overflow-hidden max-w-[150px] text-ellipsis whitespace-nowrap">
                        {{ album.name }}
                    </h2>
                    <h3 class="text-[#dcdcdc] text-sm">
                        {{ album.author.name }}
                    </h3>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
definePageMeta({
    layout: "main",
    middleware: ["auth"]
})

const { state: { songs, albums, playlists } } = useLibraryService()
</script>