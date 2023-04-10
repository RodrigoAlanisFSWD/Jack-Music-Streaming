<template>
    <NuxtLink :to="to" class="h-full">
        <li class="border-b-2 border-transparent hover:border-primary h-full flex items-center cursor-pointer transition-all duration-200 px-8">
        <slot></slot>
    </li>
    </NuxtLink>
    
</template>

<script lang="ts" setup>
defineProps({
    to: {
        type: String,
        default: "/"
    }
})
</script>