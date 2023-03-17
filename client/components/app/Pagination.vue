<script lang="ts" setup>
const emit = defineEmits(["changePage"])

const page = ref(1)

const pages = computed(() => {
    const current = page.value;

    const res = [];

    for (let i = 0; i < 5; i++) {
        res.push(current + i)
    }

    return res
})

const handleChangePage = (newPage: number) => {
    if (newPage < 1) return

    page.value = newPage;

    emit('changePage', newPage)
}
</script>

<template>
    <div class="w-full h-full grid grid-cols-1 grid-rows-[1fr_40px] justify-items-center">
        <div class="w-full h-full">
            <slot></slot>
        </div>
        <div class="text-white min-w-[200px] h-auto flex p-2 rounded-sm">
            <div class="cursor-pointer transition-all duration-200 hover:text-primary" @click="handleChangePage(page - 1)">
                Prev
            </div>
            <div class="flex mx-5 w-[200px] justify-around">
                <div @click="handleChangePage(item)" v-for="item in pages" :key="item" class="cursor-pointer transition-all duration-200 hover:text-primary" :class="{
                    'text-primary': page === item
                }">
                    {{ item }}
                </div>
            </div>
            <div class="cursor-pointer transition-all duration-200 hover:text-primary" @click="handleChangePage(page + 1)">
                Next
            </div>
        </div>
    </div>
</template>