<script lang="ts" setup>
const { state: { song, duration, currentTime, barWidth, playing },  play, pause, seek, mainBar } = usePlayer()
</script> 

<template>
    <div class="w-full h-[50px] bg-secondary text-white flex justify-between items-center px-5">
        <div>
            <i v-if="!playing" class="uil uil-play text-2xl" @click="play()"></i>
            <i v-else class="uil uil-pause text-2xl" @click="pause()"></i>
        </div>
        {{ currentTime }}
        <div ref="mainBar" class="w-[1200px] h-[15px] rounded-lg bg-[#666]" @click="seek">
            <div :style="{'width': barWidth}" class="max-w-[100%] w-0 bg-primary rounded-lg h-full">

            </div>
        </div>
        {{ duration }}
    </div>
</template>

<!-- <style>

</style> -->