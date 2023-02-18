<template>
    <div>
        <label class="block mb-2  transition-all duration-200" :class="{
            'text-primary': focus,
            'text-red-500': error
        }">
            {{ label }}
        </label>
        <input :value="modelValue" @change="$emit('change')" @input="$emit('update:modelValue', $event.target?.value)"
            @focus="focus = true" @blur="onBlur" :type="type" :placeholder="placeholder"
            class="w-full border border-white bg-transparent rounded-[8px] p-3 outline-none transition-all duration-200 mb-1"
            :class="{
                'bg-red-500 bg-opacity-25 border-red-500': error,
                'focus:border-primary': focus && !error
            }">
        <span v-if="error" class=" text-red-500 p-1">
            <i class="uil uil-times-circle"></i>
            {{ error.$message }}
        </span>
    </div>
</template>

<script lang="ts" setup>
defineProps({
    placeholder: String,
    label: String,
    type: {
        default: "text",
        type: String,
    },
    modelValue: String,
    error: Object,
})

const emit = defineEmits(['change', 'blur', 'update:modelValue'])

const focus = ref(false)

const onBlur = () => {
    focus.value = false;

    emit('blur')
}
</script>