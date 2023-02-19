<template>
    <div class="flex items-center bg-transparent text-white border w-full rounded-[5px] p-1 px-2 transition-all duration-200" :class="{
        'border-primary': focus,
        'border-white': !focus
    }">
        <i class="uil uil-search text-xl mr-2 transition-all duration-200" :class="{
            'text-primary': focus
        }"></i>
        <input @focus="focus = true" @blur="focus = false" type="text" placeholder="What are you looking for?"
            class="bg-transparent text-sm placeholder:text-white placeholder:opacity-50 outline-none w-full" />
    </div>
</template>

<script lang="ts" setup>
const focus = ref(false)
</script>