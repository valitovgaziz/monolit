<template>
    <div class="min-h-screen flex flex-col justify-center items-center">
        <h1 class="text-4xl font-bold text-center">F&A Trip</h1>
        <form class="mt-5 w-full max-w-lg mx-auto flex flex-col">
            <input class="p-3.5 rounded-t border-b border-gray-300 text-gray-900 outline-none" type="text"
                placeholder="Phone" required />
            <input class="p-3.5 rounded-b text-gray-900 outline-none" type="password" placeholder="Password" required />
            <button class="bg-blue-500 hover:bg-blue-600 mt-2.5 py-2.5 rounded">
                Log in
            </button>
        </form>
        <div class="mt-5">
            Don't have an account?
            <RouterLink class="text-blue-400" :to="{ name: 'Signup' }">
                Sign up here
            </RouterLink>
        </div>
    </div>
</template>

<script setup></script>