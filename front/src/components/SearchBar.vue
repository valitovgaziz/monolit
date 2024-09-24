<template>
    <div class="search-bar">
      <input type="text" placeholder="Search..." v-model="query"/>
      <button @click="handleSubmit">Search</button>
    </div>
  </template>
  
  <script>
  export default {
    name: 'SearchBar',
    data() {
      return {
        query: ''
      };
    },
    methods: {
      handleSubmit() {
        // Код для обработки отправки поискового запроса
        console.log('Поисковый запрос:', this.query);
      }
    }
  };
  </script>
  
  <style scoped>
  .search-bar {
    display: flex;
    width: 100%;
    max-width: 400px;
    margin: 1rem auto;
  }
  
  .search-bar input {
    width: 100%;
    height: 36px;
    padding: 8px 12px;
    outline: none;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  .search-bar button {
    height: 36px;
    padding: 0 12px;
    background-color: #ccc;
    border: none;
    cursor: pointer;
    border-radius: 4px;
  }
  </style>