<template>
    <li class="border-b-2 border-transparent hover:border-primary h-full flex items-center cursor-pointer transition-all duration-200 px-8">
        {{ text }}
    </li>
</template>

<script lang="ts" setup>
defineProps({
    text: String,
})
</script>