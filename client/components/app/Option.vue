<template>
    <div class="w-full h-[50px] flex justify-around items-center px-2 hover:bg-[#333] transition-all duration-200 cursor-pointer">
        <i :class="icon"></i>
        <slot></slot>
    </div>
</template>

<script lang="ts" setup>
defineProps({
    icon: String,
})
</script>