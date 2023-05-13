<script lang="ts" setup>
import { required, email, minLength, helpers } from '@vuelidate/validators'
import { useVuelidate } from '@vuelidate/core'
import { Ref } from 'vue';
import { Howl } from 'howler';
import { Song } from '~~/models/song';

const props = defineProps<{
    edit: boolean,
    songData?: Song
}>()

const emit = defineEmits(['submitEdit', 'submit'])

const song = reactive<{
    name: string,
}>({
    name: '',
})

const songMedia: Ref<File | null> = ref(null)
const songLogo: Ref<File | null> = ref(null)
const songLogoError: Ref<string> = ref("")
const songLogoURL: Ref<string | null> = ref(null)
const songMediaDuration = ref('0:00')
const songMediaError: Ref<string> = ref("")
const filesTouched = ref(false)

const rules = computed(() => {
    return {
        name: { required: helpers.withMessage("Name is required", required) },
    }
})

const v$ = useVuelidate(rules, song)

const handleSubmit = async () => {
    v$.value.$validate();

    if (v$.value.$error) {
        return
    }

    if (!songMedia.value) {
            songMediaError.value = "Your song need... music!"
        }

        if (!songLogo.value) {
            songLogoError.value = "Your song need a logo"
        }

        if (!songLogo.value || !songMedia.value) return

        songMediaError.value = ""
        songLogoError.value = ""

    if (props.edit) {
        emit('submitEdit', {
            song: {
                ...props.songData,
                name: song.name,
                duration: songMediaDuration.value,
            }, 
            media: songMedia.value,
            logo: songLogo.value,
            touched: filesTouched.value
        })

    } else {
        emit('submit', {
            name: song.name,
            duration: songMediaDuration.value,
            media: songMedia.value,
            logo: songLogo.value
        })
    }
}

const authStore = useAuthStore()

const { status, user } = storeToRefs(authStore)

const formatTime = (secs: number) => {
    const m = Math.floor(secs % 3600 / 60).toString().padStart(2),
        s = Math.floor(secs % 60).toString().padStart(2, '0');

    return m + ':' + s;
}

const handleMediaChange = (file: File) => {
    songMedia.value = file

    const fr = new FileReader();

    fr.onload = () => {
        const media = new Howl({
            src: [fr.result as string],
            preload: true,
            format: ['mp3', 'mp4', 'm4a'],
            onload: () => {
                songMediaDuration.value = formatTime(media.duration())
            }
        })
    }

    fr.readAsDataURL(file)

    filesTouched.value = true
}

const handleLogoChange = (file: File) => {
    songLogo.value = file

    const fr = new FileReader();

    fr.onload = () => {
        songLogoURL.value = fr.result as string
    }

    fr.readAsDataURL(file)

    filesTouched.value = true
}

onMounted(async () => {
    songsStore.setError("")
    
    if (props.edit && props.songData) {
        let res = await fetch("http://localhost:8080/api/file/" + props.songData.media_id)
        let data = await res.blob()
        let file = new File([data], props.songData.media.name)
        handleMediaChange(file)

        res = await fetch("http://localhost:8080/api/file/" + props.songData.logo_id)
        data = await res.blob()
        console.log(props.songData)
        file = new File([data], props.songData.logo.name)
        handleLogoChange(file)

        song.name = props.songData.name
        filesTouched.value = false
    }
})

const { state: { error }, songsStore } = useSongsService()
</script>

<template>
    <div class="grid grid-cols-[1fr_700px] w-full h-full">
        <div class="background w-full bg-white">

        </div>
        <div class="flex justify-center items-center flex-col">
            <div class="text-white w-[450px]">
                <h2 class="text-4xl font-bold mb-2">Upload A Song</h2>
                <p class="mb-8 w-3/4">
                    Fill the fields to upload you new song!
                </p>
                <div class="flex my-5 w-full justify-between">
                    <div class="flex">
                        <img v-if="song" :src="songLogoURL ? songLogoURL : 'http://localhost:8080/api/file/2'" class="w-[75px] h-[75px] bg-white mr-5" />
                        <div>
                            <h2>
                                {{ song?.name ? song.name : 'Song Name' }}
                            </h2>
                            <h3 class="cursor-pointer hover:text-primary text-sm text-[#dcdcdc]">
                                {{ user?.name }}
                            </h3>
                        </div>
                    </div>

                    <span>
                        {{ songMediaDuration }}
                    </span>
                </div>
                <form-input :error="v$.name.$errors[0]" @change="v$.name.$touch" @blur="v$.name.$touch" v-model="song.name"
                    label="Name" placeholder="November Rain" type="text" class="mb-3"></form-input>
                <form-file :value="songMedia" :placeholder="songMedia ? songMedia.name : 'mp3 | mp4 | m3a | m4a'"
                    label="Song Media" :error="songMediaError" @file="handleMediaChange" class="mb-3"></form-file>
                <form-file :value="songLogo" :placeholder="songLogo ? songLogo.name : 'png | jpeg | jpg'" label="Song Logo"
                    :error="songLogoError" @file="handleLogoChange"></form-file>

                <app-button @click="handleSubmit" class="mt-12">Create</app-button>
                <app-alert v-if="error" class="mt-5">{{ error }}</app-alert>
            </div>
        </div>
    </div>
</template>

<style>
.background {
    background-image: url("../../assets/img/guitar.png");
    background-size: cover;
}
</style>