<template>
    <div>
        <label class="block mb-2  transition-all duration-200" :class="{
            'text-primary': hover,
            'text-red-500': error
        }">
            {{ label }}
        </label>
        <input ref="fileInput" class="hidden" @change="$emit('change')"
            @input="handleChange" type="file" :placeholder="placeholder">
        <div @click="handleClick" @mouseover="hover = true" @mouseleave="hover = false" class="w-full border border-white bg-transparent rounded-[8px] p-3 outline-none transition-all duration-200 mb-1 text-gray-400 cursor-pointer"
            :class="{
                'bg-red-500 bg-opacity-25 border-red-500': error,
                'hover:border-primary': hover && !error
            }">
            {{ placeholder }}
        </div>
        <span v-if="error != ''" class=" text-red-500 p-1">
            <i class="uil uil-times-circle"></i>
            {{ error }}
        </span>
    </div>
</template>

<script lang="ts" setup>
import { Ref } from 'vue';

defineProps<{
    placeholder: string,
    label: string,
    error: string,
    value: File | null
}>()

const fileInput: Ref<HTMLInputElement | null> = ref(null)

const hover = ref(false)

const emit = defineEmits(['file'])

const handleClick = () => {
    fileInput.value?.click()
}

const handleChange = (e: any) => {
    emit('file', e.target.files[0])
}
</script>