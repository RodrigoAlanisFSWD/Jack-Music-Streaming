<script lang="ts" setup>
import { SelectItem } from '../../models/select'

defineProps<{
    placeholder: string,
    label: string,
    error: object,
    items: Array<SelectItem>,
    selectedItem: SelectItem | null,
}>()

const show = ref(false)

const emit = defineEmits(['select'])

const handleSelect = (item: SelectItem) => {
    show.value = false;

    emit('select', item)
}
</script>

<template>
    <div class="w-full flex flex-col max-h-[86px]">
        <div class="w-full">
            <label class="block mb-2  transition-all duration-200">
                {{ label }}
            </label>
            <div @click="show = !show"
                class="w-full border border-white bg-transparent rounded-[8px] h-[54px] outline-none flex justify-between items-center px-3"
                :class="{
                    'rounded-b-none border-b-0': show
                }">

                <p v-if="!selectedItem" class="text-gray-400">{{ placeholder }}</p>
                <p v-else>
                    {{ selectedItem.name }}
                </p>

                <i class="uil uil-angle-down text-3xl transition-all duration-200" :class="{ 'rotate-90': show }"></i>
            </div>
        </div>
        <div v-if="show" class="relative w-full bg-secondary min-h-[100px] border border-white rounded-b-md flex flex-col">
            <div class="w-full h-[50px] hover:bg-[#222] bg-opacity-5 flex items-center px-5 transition-all duration-200 cursor-pointer"
                :class="{'bg-[#222]': selectedItem?.key === item.key }" @click="handleSelect(item)" v-for="item in items"
                :key="item.key">
                {{ item.name }}
            </div>
        </div>
    </div>
</template>