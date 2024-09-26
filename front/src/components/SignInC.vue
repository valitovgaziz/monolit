<template>
    <div id="app">
      <form @submit.prevent="onSubmit">
        <label for="phone">Телефон:</label>
        <input type="tel" id="phone" v-model="phone">
        <br>
        <label for="password">Пароль:</label>
        <input type="password" id="password" v-model="password">
        <br>
        <button type="submit">Войти</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    name: 'Login',
    data() {
      return {
        phone: '',
        password: ''
      };
    },
    methods: {
      onSubmit() {
        if (!this.phone || !this.password) {
          alert('Все поля должны быть заполнены');
          return;
        }
  
        const url = 'http://localhost:8000/api/singin';
  
        try {
          axios.post(url, {
            phone: this.phone,
            password: this.password
          }).then(response => {
            if (response.data.success) {
              window.location.href = 'http://localhost:8000/account';
            } else {
              alert('Неправильный телефон или пароль');
            }
          });
        } catch (error) {
          alert(`Ошибка при входе: ${error}`);
        }
      }
    }
  };
  </script>