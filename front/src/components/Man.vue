<template>
    <div id="app">
      <header class="header">
        <!-- Логотип -->
        <img src="logo.png" alt="Logo" />
        
        <!-- Выпадающее меню слева -->
        <nav class="menu-left">
          <ul>
            <li><a href="#">Link 1</a></li>
            <li><a href="#">Link 2</a></li>
            <li><a href="#">Link 3</a></li>
          </ul>
        </nav>
        
        <!-- Выпадающее меню справа -->
        <nav class="menu-right">
          <ul>
            <li><a href="#">Login</a></li>
            <li><a href="#">Sign Up</a></li>
            <li><a href="#">About Us</a></li>
          </ul>
        </nav>
      </header>
      
      <!-- Поисковая строка -->
      <form class="search-bar">
        <input type="text" placeholder="Search..." />
        <button type="submit">Search</button>
      </form>
      
      <!-- Основной контент -->
      <main class="content">
        <!-- Тут будет основной контент страницы -->
        <p>Это пример контента для главной страницы.</p>
      </main>
      
      <!-- Подвал с информацией о компании -->
      <footer class="footer">
        <p>© 2023 Company Name Inc. All rights reserved.</p>
        <p>Company Address: 123 Main Street, City, State ZIP</p>
        <p>Phone Number: +1 (XXX) XXX-XXXX</p>
        <p>Email: <EMAIL></p>
      </footer>
    </div>
  </template>
  
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
  
  .menu-left,
  .menu-right {
    list-style: none;
    margin: 0;
    padding: 0;
  }
  
  .menu-left li,
  .menu-right li {
    display: inline-block;
    margin-right: 1rem;
  }
  
  .menu-left a,
  .menu-right a {
    color: #333;
    text-decoration: none;
  }
  
  .menu-left a:hover,
  .menu-right a:hover {
    text-decoration: underline;
  }
  
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
  
  .content {
    min-height: calc(100vh - 197px); /* Зазоры между элементами */
    padding: 1rem;
    background-color: #fafafa;
  }
  
  .footer {
    text-align: center;
    padding: 1rem;
    background-color: #eee;
    border-top: 1px solid #ccc;
  }
  </style>