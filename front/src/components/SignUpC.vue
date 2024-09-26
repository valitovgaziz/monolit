<template>
    <div id="app">
      <form @submit.prevent="onSubmit">
        <label for="phone">Телефон:</label>
        <input type="tel" id="phone" v-model="phone">
        <br>
        <label for="password">Пароль:</label>
        <input type="password" id="password" v-model="password">
        <br>
        <button type="submit">Зарегистрироваться</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    name: 'FastSignUp',
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
  
        const url = 'http://localhost:8000/SingUp';
  
        try {
          axios.post(url, {
            phone: this.phone,
            password: this.password
          }).then(() => {
            window.location.href = 'http://localhost:8000/';
          });
        } catch (error) {
          alert(`Ошибка при регистрации: ${error}`);
        }
      }
    }
  };
  </script>