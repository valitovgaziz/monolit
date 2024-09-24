<template>
    <header class="header">
      <!-- Логотип -->
      <img src="logo.png" alt="Logo" />
      
      <!-- Выпадающее меню справа -->
      <nav class="menu-right">
        <ul>
          <li><a href="#">Login</a></li>
          <li><a href="#">Sign Up</a></li>
          <li><a href="#">About Us</a></li>
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
body {
  font-family: Arial, sans-serif;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background-color: #fff;
  border-bottom: 1px solid #ccc;
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