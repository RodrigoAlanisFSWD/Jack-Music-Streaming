<template>
    <NuxtLink :to="to">
        <li class="flex items-center text-2xl hover:text-white transition-all duration-200">
            <i :class="icon"></i>
            <h3 class="text-sm ml-5">
                {{ text }}
            </h3>
        </li>
    </NuxtLink>
</template>

<script lang="ts" setup>
defineProps({
    text: String,
    to: {
        type: String,
        default: "/search"
    },
    icon: {
        type: String,
        default: "uil uil-estate"
    },
})
</script>

<style scoped>
.router-link-active {
  @apply text-opacity-100 text-white;
}
</style>