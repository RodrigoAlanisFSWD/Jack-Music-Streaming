<script lang="ts" setup>
import useVuelidate from '@vuelidate/core';
import { helpers, required } from '@vuelidate/validators'
import { Ref } from 'vue';

const { state: { error } } = usePlaylistsService()

const emit = defineEmits(['submit'])

const albumData = reactive({
    name: ''
})

const rules = computed(() => ({
    name: { required: helpers.withMessage("Name is required", required), }
}))

const v$ = useVuelidate(rules, albumData)

const albumLogo: Ref<File | null> = ref(null)
const albumLogoError: Ref<string> = ref("")
const albumLogoURL: Ref<string | null> = ref(null)

const handleLogoChange = (file: File) => {
    albumLogo.value = file

    const fr = new FileReader();

    fr.onload = () => {
        albumLogoURL.value = fr.result as string
    }

    fr.readAsDataURL(file)
}

const user = useUser()
</script>

<template>
    <div class="text-white w-[400px] flex flex-col items-center">
        <div class="flex my-5 w-full justify-between">
            <div class="flex">
                <img v-if="albumData" :src="albumLogoURL ? albumLogoURL : 'http://localhost:8080/api/file/2'"
                    class="w-[75px] h-[75px] bg-white mr-5 rounded-full" />
                <div>
                    <h2>
                        {{ albumData.name ? albumData.name : 'Album Name' }}
                    </h2>
                    <h3 class="cursor-pointer hover:text-primary text-sm text-[#dcdcdc]">
                        {{ user?.name }}
                    </h3>
                </div>
            </div>
        </div>
        <form-input :error="v$.name.$errors[0]" @change="v$.name.$touch" @blur="v$.name.$touch" v-model="albumData.name"
            label="Name" placeholder="Mothership" type="text" class="mb-3 w-full"></form-input>
        <form-file :value="albumLogo" :placeholder="albumLogo ? albumLogo.name : 'png | jpeg | jpg'"
            label="Album Logo" :error="albumLogoError" @file="handleLogoChange" class="w-full"></form-file>
        <app-button @click="emit('submit', {
            name: albumData.name,
        }, albumLogo)" class="mt-12">Submit</app-button>
        <app-alert v-if="error" class="mt-5">{{ error }}</app-alert>
    </div>
</template>