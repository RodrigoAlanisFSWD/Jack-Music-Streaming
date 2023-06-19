<script lang="ts" setup>
import { Library } from '~/models/library';
import { Song } from '~~/models/song';

const { songs = [], playlist = false } = defineProps<{
    songs: Song[],
    playlist: boolean
}>()

const emit = defineEmits(['remove'])

const showPlaylistModal = ref(false)
const selectedSong = ref(0)

const { state, addSong } = useLibraryService()

const router = useRouter()

const handleAddSong = async (song: Song) => {
    await addSong(song.id)
    router.push("/library")      
}
</script>

<template>
    <div class="w-full h-full">
        <div class="w-full h-full flex flex-col">
            <div
                class="grid grid-cols-4 justify-items-center items-center px-5 text-[#dcdcdc] border-b border-[#343434] py-2 pr-[80px]">
                <span class="justify-self-start">
                    Title
                </span>
                <span class="mr-11">
                    Album
                </span>
                <span class="mr-9">
                    Date added
                </span>
                <span class="justify-self-end mr-4">
                    <i class="uil uil-clock"></i>
                </span>
            </div>
            <app-song v-for="song in songs" :key="song.id" :song="song">
                <template #options>
                    <app-options>
                        <app-option @click="() => {
                            selectedSong = song.id
                            showPlaylistModal = true
                        }" icon="uil uil-plus-circle">
                            Add To Playlist
                        </app-option>
                        <app-option v-if="playlist" @click="emit('remove', song.id)" icon="uil uil-minus-circle">
                            Remove From Playlist
                        </app-option>
                        <app-option v-if="!state.songs.value.find((s: Song) => s.id === song.id)" @click="handleAddSong(song)" icon="uil uil-plus-circle">
                            Add To Library
                        </app-option>
                        
                        <slot name="options">

                        </slot>
                    </app-options>
                </template>
            </app-song>
        </div>
        <app-select-playlist-modal :selected-song="selectedSong" @hide="showPlaylistModal = false"
            :show="showPlaylistModal"></app-select-playlist-modal>
    </div>
</template>