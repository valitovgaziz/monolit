<template>
    <div class="container mx-auto px-4">
      <header class="text-center py-8">
        <h1 class="text-2xl font-semibold">Личный кабинет</h1>
      </header>
      
      <!-- Блок с фото -->
      <section class="py-6">
        <img src="../assets/logo.svg" alt="Фото пользователя" class="w-full h-auto object-cover rounded-lg shadow-md" />
      </section>
      
      <!-- Блок с информацией о аккаунте -->
      <section class="py-6 border-b border-gray-200">
        <div class="flex flex-col gap-4">
          <div class="text-left">
            <h2 class="text-lg font-medium">Информация о аккаунте</h2>
            <p class="text-sm text-gray-700">Ваше имя: {{ accountInfo.name }}</p>
            <p class="text-sm text-gray-700">Электронная почта: {{ accountInfo.email }}</p>
            <p class="text-sm text-gray-700">Дата регистрации: {{ accountInfo.registrationDate }}</p>
          </div>
        </div>
      </section>
      
      <!-- Блок с доступными объектами -->
      <section class="py-6">
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          <div v-for="object in objects" :key="object.id">
            <div class="card bg-white shadow-md rounded-lg p-4">
              <h3 class="font-semibold">{{ object.title }}</h3>
              <p class="text-sm text-gray-700">{{ object.description }}</p>
              <button @click="handleObjectAction(object)" class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded">Действие</button>
            </div>
          </div>
        </div>
      </section>
    </div>
  </template>
  
  <script >
  import { reactive } from 'vue';
  
  export default {
    data() {
      return {
        accountInfo: {
          name: '<NAME>',
          email: '<EMAIL>',
          registrationDate: new Date().toLocaleDateString('ru')
        },
        objects: [
          { id: 1, title: 'Объект 1', description: 'Описание объекта 1' },
          { id: 2, title: 'Объект 2', description: 'Описание объекта 2' },
          { id: 3, title: 'Объект 3', description: 'Описание объекта 3' },
          { id: 4, title: 'Объект 4', description: 'Описание объекта 4' }
        ],
      };
    },
    methods: {
      handleObjectAction(object) {
        // Тут можете реализовать логику действий с объектом
      },
    },
  };
  </script>
  

  <style lang="tailwind">
  @layer base {
    body {
      @apply bg-gray-100 text-gray-900;
    }
  }
  </style>