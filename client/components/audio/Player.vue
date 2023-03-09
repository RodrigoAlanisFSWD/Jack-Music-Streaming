<script lang="ts" setup>
import { Ref } from 'vue';

const playing = ref(false)
const duration = ref(0)
const playbackTime = ref(0)

const player: Ref<HTMLAudioElement | null> = ref(null)
const progress: Ref<HTMLProgressElement | null> = ref(null)

const togglePlaying = () => {
    playing.value = !playing.value
    console.log(player)

    if (player && player.value) {
        if (playing.value) {
            player.value.play()
        } else {
            player.value.pause()
        }
    }
}

const songEnded = () => {
    console.log("ended")
}

const setDuration = (e: any) => {
    duration.value = e.target.duration

    if (progress.value) {
        progress.value.max = duration.value
    }
}

const updateProgress = (e: any) => {
    if (player.value && progress.value) {
        playbackTime.value = player.value?.currentTime;
    }
}

const fortmatTime = (seconds: number) => {
    const format = (val: number) => `0${Math.floor(val)}`.slice(-2);
    var hours = seconds / 3600;
    var minutes = (seconds % 3600) / 60;

    console.log(seconds % 60)
    return [minutes, seconds % 60].map(format).join(":");
}

const totalTime = computed(() => {
    let audio = player.value;

    if (audio) {
        return fortmatTime(duration.value)
    } else {
        return '00:00'
    }
})

const elapsedTime = computed(() => {
    let audio = player.value;

    if (audio) {
        return fortmatTime(playbackTime.value)
    } else {
        return '00:00'
    }
})
</script> 

<template>
    <div class="w-[75%] h-[50px] bg-secondary text-white flex justify-between items-center px-5">
        <audio ref="player" src="http://localhost:8080/song" @ended="songEnded" @canplay="setDuration"
            @timeupdate="updateProgress">
        </audio>
        <div @click="togglePlaying">
            <i v-if="playing === false" class="uil uil-play hover:text-primary cursor-pointer text-xl"></i>
            <i v-if="playing === true" class="uil uil-pause hover:text-primary cursor-pointer text-xl"></i>
        </div>
        <progress :value="playbackTime" class="progress" ref="progress"></progress>
        <div class="text-white">
            <span class="text-white">
                {{ elapsedTime }}
            </span>

            /

            <span class="text-white">
                {{ totalTime }}
            </span>
        </div>
    </div>
</template>

<style>
.progress {
    border-radius: 10px;
    width: 85%;
}

.progress::-webkit-progress-bar {
    background-color: #666;
    border-radius: 8px;
}

.progress::-webkit-progress-value {
    background-color: #B18CFF;
    border-radius: 10px;
}
</style>