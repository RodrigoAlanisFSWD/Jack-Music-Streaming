<template>
    <div :class="getStyles">
        <i v-if="leftIcon" :class="leftIcon + ' ' + 'text-xl'"></i>
        <slot></slot>
        <i v-if="rightIcon" :class="rightIcon + ' ' + 'text-xl'"></i>
    </div>
</template>

<script lang="ts" setup>
const { type, rightIcon, leftIcon } = defineProps({
    type: String,
    rightIcon: String,
    leftIcon: String,
})

const getStyles = computed(() => {
    const baseStyles =  `border cursor-pointer rounded-md p-3 w-full flex items-center transition-all duration-200 ${leftIcon || rightIcon ? 'justify-between px-5' : 'justify-center'}`

    switch(type) {
        case 'primary':
            return `${baseStyles} bg-primary border-primary hover:bg-transparent hover:text-primary text-secondary`
        case 'secondary':
            return `${baseStyles} bg-secondary border-white text-white hover:bg-primary hover:border-primary hover:text-secondary`
        default: 
            return `${baseStyles} bg-primary border-primary hover:bg-transparent hover:text-primary text-secondary`
    }
})
</script>