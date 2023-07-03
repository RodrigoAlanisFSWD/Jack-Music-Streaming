<script lang="ts" setup>
import { User } from '~/models/auth';

defineProps<{
    logo: string,
    name: string,
    author: User
}>()

const emit = defineEmits(['remove', 'detail'])

const handleRemoveFromLibrary = () => {
    emit('remove')
}
</script>

<template>
    <div class="flex flex-col max-w-[150px]">
        <img @click="emit('detail')" :src="logo"
            class="w-[150px] h-[150px] bg-white mr-5 cursor-pointer rounded-sm mb-2" />
            <div class="flex justify-between">
                <h2 @click="emit('detail')" class="text-sm overflow-hidden max-w-[100px] text-ellipsis whitespace-nowrap cursor-pointer">
                    {{ name }}
                </h2>
                <app-options :right="true">
                    <app-option @click="handleRemoveFromLibrary" icon="uil uil-plus-circle">
                        Remove From Library
                    </app-option>
                </app-options>
            </div>
        
        <h3 class="text-[#dcdcdc] text-sm">
            {{ author.name }}
        </h3>
    </div>
</template>