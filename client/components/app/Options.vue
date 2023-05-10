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

        <Transition enter-from-class="opacity-0" enter-active-class="transition duration-300" leave-to-class="opacity-0"
            leave-active-class="transition duration-300">
            <div v-if="show" class="absolute w-auto min-h-[50px] bg-[#222] rounded-md z-30 right-5 mt-2 overflow-hidden">
                <slot></slot>
            </div>
        </Transition>
    </div>
</template>