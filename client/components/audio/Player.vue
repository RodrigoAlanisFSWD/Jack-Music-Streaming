<script lang="ts" setup>
import { Howl } from 'howler';
import { Ref } from 'vue';
import { Song } from '~~/models/song';

const { state: { song, duration, currentTime, barWidth, playing, playlist, loading }, play, pause, seek, mainBar, setSong, playerStore } = usePlayer()
const { state: { songs } } = useSongsService()

const bar: Ref<HTMLDivElement | null> = ref(null)
const barEmpty: Ref<HTMLDivElement | null> = ref(null)

const volume = (val: number) => {
    if (bar.value) {
        Howler.volume(val)

        const barWidth = (val * 100) / 100;
        bar.value.style.width = (barWidth * 100) + '%'
    }
}

const captureVolume = (e: any) => {
    if (bar.value && barEmpty.value) {
        var val = e.offsetX / barEmpty.value.scrollWidth;

        volume(val)
    }
}

onMounted(() => {
    volume(Howler.volume())
})

const handleNextSong = () => {
    const index = playlist.value.findIndex((s: Song) => s.id === song.value?.id)

    if (index < 0) {
        setSong(playlist.value[0])
    } else if (playlist.value.length - 1 === index) {
        return
    } else {
        setSong(playlist.value[index + 1])
    }
}

const handlePrevSong = () => {
    const index = playlist.value.findIndex((s: Song) => s.id === song.value?.id)

    if (index < 0) {
        return
    } else if (index === 0) {
        return
    } else {
        setSong(playlist.value[index - 1])
    }
}
</script> 

<template>
    <div class="w-full bg-secondary text-white flex justify-center items-center border-t border-[#343434] h-[75px]">
        <div class="flex absolute left-5 bottom-0">
            <img v-if="song" :src="song ? 'http://localhost:8080/api/file/' + song?.logo_id : ''"
                class="w-[75px] h-[75px] bg-white mr-5" />
            <div class="pt-2">
                <h2>
                    {{ song?.name }}
                </h2>
                <h3 class="cursor-pointer hover:text-primary text-sm text-[#dcdcdc]">
                    {{ song?.author.name }}
                </h3>
            </div>
        </div>
        <div class="flex flex-col items-center">
            <div class="mb-1 flex items-center mt-4">
                <button :disabled="loading" @click="handlePrevSong">
                    <i class="uil uil-step-backward text-[#dcdc] transition-all duration-200"
                        :class="{ 'opacity-25': loading }"></i>
                </button>
                <div class="mx-5 cursor-pointer">
                    <i v-if="!playing" class="uil uil-play text-2xl" @click="play()"></i>
                    <i v-else class="uil uil-pause text-2xl" @click="pause()"></i>
                </div>
                <button :disabled="loading" @click="handleNextSong">
                    <i class="uil uil-skip-forward text-[#dcdc] transition-all duration-200"
                        :class="{ 'opacity-25': loading }"></i>
                </button>
            </div>

            <div class="flex items-center">
                {{ currentTime }}
                <div ref="mainBar" class="w-[700px] h-[8px] rounded-lg bg-[#343434] mx-3" @click="seek">
                    <div :style="{ 'width': barWidth }" class="max-w-[100%] w-0 bg-primary rounded-lg h-full">

                    </div>
                </div>
                {{ duration }}
            </div>

        </div>
        <div class="flex absolute right-5 items-center mt-5">
            <i class="uil uil-volume-up text-xl "></i>
            <div ref="barEmpty" class="w-[150px] h-[8px] rounded-lg bg-[#343434] mx-3" @click="captureVolume">
                <div ref="bar" class="max-w-[100%] w-0 bg-primary rounded-lg h-full">

                </div>
            </div>
        </div>

    </div>
</template>

<!-- <style>

</style> -->