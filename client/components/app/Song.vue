<script lang="ts" setup>
import { Song } from '~~/models/song';

defineProps<{
    song: Song
}>()

const { setSong } = usePlayer()
</script>

<template>
    <div class="w-full p-3 text-white grid grid-cols-4 justify-items-center items-center px-5">
        <div class="flex justify-self-start">
            <img :src="song ? 'http://localhost:8080/api/file/' + song?.logo_id : ''"
                class="w-[50px] h-[50px] bg-white mr-5 cursor-pointer" @click="setSong(song)" />
            <div>
                <h2 class="hover:text-primary cursor-pointer" @click="setSong(song)">
                    {{ song?.name }}
                </h2>
                <h3 class="text-sm text-[#dcdcdc] hover:text-primary cursor-pointer">
                    {{ song?.author.name }}
                </h3>
            </div>
        </div>
        <span>
            Album Name
        </span>
        <span>
            October 5, 2022
        </span>
        <span class="justify-self-end">
            {{ song?.duration }}
        </span>
    </div>
</template>