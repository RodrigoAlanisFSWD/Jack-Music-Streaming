<template>
    <div class="w-full h-full grid grid-cols-[400px_1fr]">
        <div class="flex flex-col items-center p-8 text-white">
            <img v-if="user?.profile" :src="user?.profile ? 'http://localhost:8080/api/file/' + user?.profile.logo_id : ''" class="w-[200px] h-[200px] rounded-full bg-white" />
            <input ref="fileInput" class="hidden" @change="$emit('change')"
            @input="handleChange" type="file">
            <app-text-button @click="handleClick" class="w-[200px] mt-5" left-icon="uil uil-camera">
                <h3 class="mr-5">
                    Select Avatar
                </h3>
            </app-text-button>
            <h2 class="text-4xl mt-12">
                {{ user?.name }}
            </h2>
            <h3 class="mt-2">
                {{ user?.email }} 
            </h3>
            <div class="w-full flex justify-between p-8 px-12">
                <span>
                    Followers
                </span>
                <span>
                    100
                </span>
            </div>
        </div>
        <NuxtPage />
    </div>
</template>

<script lang="ts" setup>
definePageMeta({
    middleware: 'auth',
    layout: "default"
})

const user = useUser()

const { uploadUserLogo } = useAuthService()

const fileInput: Ref<HTMLInputElement | null> = ref(null)

const handleClick = () => {
    fileInput.value?.click()
}

const handleChange = (e: any) => {
    uploadUserLogo(e.target.files[0])
}
</script>