<script lang="ts" setup>
import useVuelidate from '@vuelidate/core';
import { helpers, required } from '@vuelidate/validators'
import { Ref } from 'vue';
import { SelectItem } from '~~/models/select';

const { state: { error } } = usePlaylistsService()

const emit = defineEmits(['submit'])

const playlistData = reactive({
    name: ''
})

const rules = computed(() => ({
    name: { required: helpers.withMessage("Name is required", required), }
}))

const v$ = useVuelidate(rules, playlistData)

const playlistLogo: Ref<File | null> = ref(null)
const playlistLogoError: Ref<string> = ref("")
const playlistLogoURL: Ref<string | null> = ref(null)

const selectedType: Ref<SelectItem | null> = ref(null)

const handleLogoChange = (file: File) => {
    playlistLogo.value = file

    const fr = new FileReader();

    fr.onload = () => {
        playlistLogoURL.value = fr.result as string
    }

    fr.readAsDataURL(file)

    // filesTouched.value = true
}

const types: Array<SelectItem> = [
    {
        key: "PUBLIC",
        name: "Public"
    },
    {
        key: "PRIVATE",
        name: "Private"
    }
]

const user = useUser()
</script>

<template>
    <div class="text-white w-[400px] flex flex-col items-center">
        <div class="flex my-5 w-full justify-between">
            <div class="flex">
                <img v-if="playlistData" :src="playlistLogoURL ? playlistLogoURL : ''"
                    class="w-[75px] h-[75px] bg-white mr-5 rounded-full" />
                <div>
                    <h2>
                        {{ playlistData.name ? playlistData.name : 'Playlist Name' }}
                    </h2>
                    <h3 class="cursor-pointer hover:text-primary text-sm text-[#dcdcdc]">
                        {{ user?.name }}
                    </h3>
                </div>
            </div>
        </div>
        <form-input :error="v$.name.$errors[0]" @change="v$.name.$touch" @blur="v$.name.$touch" v-model="playlistData.name"
            label="Name" placeholder="Mothership" type="text" class="mb-3 w-full"></form-input>
        <app-select @select="(item) => selectedType = item" :selected-item="selectedType" :items="types"
            placeholder="Select Playlist Type" label="Type" class="mb-3"></app-select>
        <form-file :value="playlistLogo" :placeholder="playlistLogo ? playlistLogo.name : 'png | jpeg | jpg'"
            label="Song Logo" :error="playlistLogoError" @file="handleLogoChange" class="w-full"></form-file>
        <app-button @click="emit('submit', {
            name: playlistData.name,
            type: selectedType?.key,
        }, playlistLogo)" class="mt-12">Submit</app-button>
        <app-alert v-if="error" class="mt-5">{{ error }}</app-alert>
    </div>
</template>