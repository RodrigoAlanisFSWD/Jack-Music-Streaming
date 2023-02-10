<template>
    <div>
        <label class="block mb-2">
            {{ label }}
        </label>
        <input :type="type" :placeholder="placeholder" class="w-full border border-white bg-transparent rounded-[8px] p-3">
    </div>
</template>

<script lang="ts" setup>
defineProps({
    placeholder: String,
    label: String,
    type: {
        default: "text",
        type: String,
    }
})
</script>