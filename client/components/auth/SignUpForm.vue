<script lang="ts" setup>
import { required, email, minLength, helpers } from '@vuelidate/validators'
import { useVuelidate } from '@vuelidate/core'
import { AuthStatus } from '~~/store/authStore';

const emit = defineEmits(['submit'])

const userData = reactive({
    name: '',
    email: '',
    password: ''
})

const rules = computed(() => {
    return {
        name: { required: helpers.withMessage("Name is required", required) },
        email: { required: helpers.withMessage("Mail is required", required), email: helpers.withMessage("Mail is invalid", email) },
        password: { required: helpers.withMessage("Password is required", required), minLenght: helpers.withMessage("Password is too short", minLength(6)) }
    }
})

const v$ = useVuelidate(rules, userData)

const handleSubmit = async () => {
    v$.value.$validate();

    if (v$.value.$error) {
        return
    }

    emit('submit', userData)
}

const authStore = useAuthStore()

const { status } = storeToRefs(authStore)
</script>

<template>
    <div class="flex justify-center items-center flex-col">
        <div class="text-white w-[450px]">
            <h2 class="text-4xl font-bold mb-2">Sign up</h2>
            <p class="mb-8 w-3/4">
                Welcome to Jack, to start your need to create a new account.
            </p>
            <form-input :error="v$.name.$errors[0]" @change="v$.name.$touch" @blur="v$.name.$touch" v-model="userData.name"
                label="Name" placeholder="Axel Role" type="text" class="mb-3"></form-input>
            <form-input :error="v$.email.$errors[0]" @change="v$.email.$touch" @blur="v$.email.$touch"
                v-model="userData.email" label="Email" placeholder="axel@gmail.com" type="email" class="mb-3"></form-input>
            <form-input :error="v$.password.$errors[0]" @change="v$.password.$touch" @blur="v$.password.$touch"
                v-model="userData.password" label="Password" placeholder="super_secure_password"
                type="password"></form-input>
            <app-button @click="handleSubmit" class="mt-12">Submit</app-button>
            <app-alert v-if="status === AuthStatus.AUTH_ERROR" class="mt-5">Las credenciales no son validas</app-alert>
            <p class="text-sm mt-3">
                By registering you agree with our <span class="text-primary cursor-pointer">Privacy Policy</span>
            </p>
        </div>
        <NuxtLink to="/auth/signIn" class="absolute text-primary bottom-0 right-0 m-5 flex items-center">
            <h3>
                Sign In <i class="uil uil-arrow-right text-xl mt-1 ml-2"></i>
            </h3>
        </NuxtLink>

    </div>
</template>