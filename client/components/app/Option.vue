<template>
    <div class="w-full h-[50px] flex justify-start items-center px-5 hover:bg-[#333] transition-all duration-200 cursor-pointer">
        <i :class="icon" class="mr-5"></i>
        <slot></slot>
    </div>
</template>

<script lang="ts" setup>
defineProps({
    icon: String,
})
</script>