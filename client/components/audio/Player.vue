<script lang="ts" setup>
import { Howl } from 'howler'
import { Ref } from 'vue';

let player = new Howl({
    src: ["http://localhost:8080/api/songs/media/4"],
    format: "mp4",
    preload: true,
    onplay: function () {
        requestAnimationFrame(step.bind(this))
    },
    onseek: function () {
        // Start updating the progress of the track.
        requestAnimationFrame(step.bind(self));
    }
})

const duration = ref("0:00")
const currentTime = ref("0:00")

const bar: Ref<HTMLDivElement | null> = ref(null)
const mainBar: Ref<HTMLDivElement | null> = ref(null) 

player.once("rate", () => {
})

player.once("load", () => {
    console.log("load")

    duration.value = formatTime(player.duration())
})

const formatTime = (secs: number) => {
    var minutes = Math.floor(secs / 60) || 0;
    var seconds = (secs - minutes * 60) || 0;

    return minutes + ':' + (seconds < 10 ? '0' : '') + seconds;
}

function step() {
    const seek = player.seek()
    currentTime.value = formatTime(Math.round(seek))

    if (bar.value) {
        bar.value.style.width = (((seek / player.duration()) * 100) || 0) + '%'
    }

    if (player.playing()) {
        requestAnimationFrame(step);
    }
}

const seek = (e: any) => {
    if (mainBar.value) {
        console.log(mainBar.value.scrollWidth)
        console.log(e.clientX)
        var per = e.offsetX / mainBar.value.scrollWidth

        if (player.playing()) {
            player.seek(player.duration() * per)
        }
    }
}
</script> 

<template>
    <div class="w-full h-[50px] bg-secondary text-white flex justify-between items-center px-5">
        <div>
            <i v-if="!player.playing()" class="uil uil-play text-2xl" @click="player.play()"></i>
            <i v-else class="uil uil-pause text-2xl" @click="player.pause()"></i>
        </div>
        {{ currentTime }}
        <div ref="mainBar" class="w-[1200px] h-[15px] rounded-lg bg-[#666]" @click="seek">
            <div  ref="bar" class="max-w-[100%] w-0 bg-primary rounded-lg h-full">

            </div>
        </div>
        {{ duration }}
    </div>
</template>

<!-- <style>

</style> -->