<template>
    <header class="header">
      <!-- Логотип -->
      <img id="logo" src="/src/assets/goose_logo.png" alt="Logo" />
      
      <!-- Выпадающее меню справа -->
      <nav class="menu-right">
        <ul>
          <li><router-link to="http://localhost:8000/Signin">Login</router-link></li>
          <li><router-link href="http://localhost:8000/Signin">Sign Up</router-link></li>
          <li><router-link href="#">About Us</router-link></li>
        </ul>
      </nav>
    </header>
  </template>

<script>
export default {
  name: 'Header',
};
</script>

<style scoped>
.header {
  font-family: Arial, sans-serif;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background-color: #b1b1fd;
  border-bottom: 1px solid #ccc;
}

#logo {
    width: 25px;
}

.menu-right {
  list-style: none;
  margin: 0;
  padding: 0;
}

.menu-right li {
  display: inline-block;
  margin-right: 1rem;
}

.menu-right a {
  color: #333;
  text-decoration: none;
}

.menu-right a:hover {
  text-decoration: underline;
}
</style>