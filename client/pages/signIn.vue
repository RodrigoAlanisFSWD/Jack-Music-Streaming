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
                <form-input label="Email" placeholder="axel@gmail.com" type="text" class="mb-3"></form-input>
                <form-input label="Password" placeholder="super_secure_password" type="text"></form-input>
                <Button text="Submit" class="mt-12" />
                <p class="text-sm mt-3">
                    By registering you agree with our <span class="text-primary">Privacy Policy</span>
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