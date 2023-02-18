<script lang="ts" setup>
import { useAuthService } from '~~/hooks/authService';
import { required, email, sameAs, minLength, helpers } from '@vuelidate/validators'
import { useVuelidate } from '@vuelidate/core'

const { authStore, signIn } = useAuthService()

const userData = reactive({
    email: '',
    password: ''
})

const rules = computed(() => {
    return {
        email: { required: helpers.withMessage("Mail is required", required), email: helpers.withMessage("Mail is invalid", email) },
        password: { required: helpers.withMessage("Password is required", required), minLenght: helpers.withMessage("Password is too short", minLength(6)) }
    }
})

const v$ = useVuelidate(rules, userData)

const router = useRouter()

const handleSubmit = async () => {
    v$.value.$validate();

    if (v$.value.$error) {
        return
    }

    try {
        await signIn(userData)
        
        router.push("/profile")
    } catch (error) {
        return
    }
}
</script>

<template>
    <div class="grid grid-cols-[1fr_650px] w-screen h-[calc(100vh-70px)]">
        <div class="background w-full bg-white">

        </div>
        <div class="flex justify-center items-center flex-col">
            <div class="text-white w-[450px]">
                <h2 class="text-4xl font-bold mb-2">Sign In</h2>
                <p class="mb-8 w-3/4">
                    Welcome to Jack, please enter your account credentials.
                </p>
                <form-input :error="v$.email.$errors[0]" @change="v$.email.$touch" @blur="v$.email.$touch"
                    v-model="userData.email" label="Email" placeholder="axel@gmail.com" type="email"
                    class="mb-3"></form-input>
                <form-input :error="v$.password.$errors[0]" @change="v$.password.$touch" @blur="v$.password.$touch"
                    v-model="userData.password" label="Password" placeholder="super_secure_password"
                    type="password"></form-input>
                <Button @click="handleSubmit" text="Submit" class="mt-12" />
                <p class="text-sm mt-3">
                    By registering you agree with our <span class="text-primary cursor-pointer">Privacy Policy</span>
                </p>
            </div>
            <NuxtLink to="/signUp" class="absolute text-primary bottom-0 m-5 flex items-center mr-[525px]">
                <h3 >
                <i class="uil uil-arrow-left text-xl mt-1 ml-2"></i> Sign Up 
                </h3>
            </NuxtLink>   
        </div>
    </div>
</template>

<style>
    .background {
        background-image: url("../assets/img/guitar.png");
        background-size: cover;
    }
</style>