<script lang="ts" setup>
const show = ref(false)

const handleClickOutside = () => {
    if (show.value) {
        console.log("click")
        show.value = false
    }
}
</script>

<template>
    <div v-click-outside="handleClickOutside">
        <i class="uil uil-ellipsis-v m-5 text-xl" @click="show = !show"></i>

        <div v-if="show" class="absolute w-[175px] min-h-[50px] bg-[#222] rounded-md z-30 right-5 mt-2 overflow-hidden">
            <slot></slot>
        </div>
    </div>
</template>