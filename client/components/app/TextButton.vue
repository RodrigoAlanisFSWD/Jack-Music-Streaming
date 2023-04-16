<template>
    <div :class="getStyles">
        <i v-if="leftIcon" :class="leftIcon + ' ' + 'text-xl mr-3'"></i>
        <slot></slot>
        <i v-if="rightIcon" :class="rightIcon + ' ' + 'text-xl ml-3'"></i>
    </div>
</template>

<script lang="ts" setup>
const { type, rightIcon, leftIcon } = defineProps({
    type: String,
    rightIcon: String,
    leftIcon: String,
})

const getStyles = computed(() => {
    const baseStyles =  `cursor-pointer rounded-md p-3 w-full flex items-center transition-all duration-200 hover:bg-opacity-25 text-white justify-center`

    switch(type) {
        case 'primary':
            return `${baseStyles} bg-transparent hover:bg-primary hover:text-primary`
        default: 
            return `${baseStyles} bg-transparent hover:bg-primary hover:text-primary`
    }
})
</script>