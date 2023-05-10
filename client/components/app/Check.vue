<script lang="ts" setup>
defineProps({
    active: Boolean,
})
</script>

<template>
    <div :class="{'bg-primary border-primary': active, 'border-[#333] bg-[#333]': !active}" class="w-[25px] h-[25px] rounded-full border  flex items-center justify-center cursor-pointer">
        <i v-if="active" class="uil uil-check"></i>
    </div>
</template>