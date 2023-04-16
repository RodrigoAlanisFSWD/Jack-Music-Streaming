<template>
    <div class="w-full h-[calc(100vh-180px)] grid grid-cols-[400px_1fr]">
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
        <div class="flex flex-col items-start">
            <div
                class="grid grid-cols-5 justify-items-center items-center px-5 text-[#dcdcdc] border-b border-[#343434] py-2 w-full">
                <span class="justify-self-start">
                    Title
                </span>
                <span class="mr-11">
                    Album
                </span>
                <span class="mr-9">
                    Date added
                </span>
                <span class="mr-1">
                    <i class="uil uil-clock"></i>
                </span>
                <span>
                    Actions
                </span>
            </div>
            <app-song-dashboard v-for="song in songs" :key="song.id" :song="song"></app-song-dashboard>
            <app-button @click="router.push('/songs/create')" class="mt-8 w-[250px]">Upload Song</app-button>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { Ref } from 'vue';

definePageMeta({
    middleware: 'auth',
    layout: "player"
})

const user = useUser()

const router = useRouter()

const { getSongsByAuthor, state: { songs } } = useSongsService()
const { uploadUserLogo } = useAuthService()

onMounted(async () => {
    if (user.value) {
        await getSongsByAuthor(user.value?.id as number)
    }
})

const fileInput: Ref<HTMLInputElement | null> = ref(null)

const handleClick = () => {
    fileInput.value?.click()
}

const handleChange = (e: any) => {
    uploadUserLogo(e.target.files[0])
}
</script>