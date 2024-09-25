<template>
    <div id="app">
      <form @submit.prevent="registerUser">
        <label for="login">Логин:</label>
        <input type="text" id="login" v-model="login">
        <br>
        <label for="phone">Номер телефона:</label>
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
    name: 'RegisterForm',
    data() {
      return {
        login: '',
        phone: '',
        password: ''
      };
    },
    methods: {
      registerUser() {
        if (!this.login || !this.phone || !this.password) {
          alert('Все поля должны быть заполнены');
          return;
        }
  
        const url = '/api/register';
  
        try {
          await axios.post(url, {
            login: this.login,
            phone: this.phone,
            password: this.password
          });
          window.location.href = '/';
        } catch (err) {
          alert(`Ошибка при регистрации: ${err}`);
        }
      }
    }
  };
  </script>