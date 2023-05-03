<script lang="ts" setup>
import useVuelidate from '@vuelidate/core';
import { helpers, required } from '@vuelidate/validators'
import { Ref } from 'vue';
import { SelectItem } from '~~/models/select';

definePageMeta({
    layout: "main",
    middleware: ["auth"]
})

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

const handleSubmit = () => {

}

const types: Array<SelectItem> = [
    {
        key: 1,
        name: "Public"
    },
    {
        key: 2,
        name: "Private"
    }
]

const user = useUser()
</script>

<template>
    <div class="w-full h-full flex justify-center items-center">
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
            <form-input :error="v$.name.$errors[0]" @change="v$.name.$touch" @blur="v$.name.$touch"
                v-model="playlistData.name" label="Name" placeholder="Mothership" type="text"
                class="mb-3 w-full"></form-input>
            <app-select :onSelect="(item) => selectedType = item" :selected-item="selectedType" :items="types" placeholder="Select Playlist Type" label="Type" class="mb-3"></app-select>
            <form-file :value="playlistLogo" :placeholder="playlistLogo ? playlistLogo.name : 'png | jpeg | jpg'"
                label="Song Logo" :error="playlistLogoError" @file="handleLogoChange" class="w-full"></form-file>
            <app-button @click="handleSubmit" class="mt-12">Submit</app-button>
        </div>
    </div>
</template>