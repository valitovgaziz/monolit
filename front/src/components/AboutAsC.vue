<template>
    <div id="app">
      <header class="header">
        <h1>О компании</h1>
      </header>
  
      <section class="container">
        <article class="card">
          <h2>История</h2>
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
        </article>
  
        <article class="card">
          <h2>Фотографии</h2>
          <div class="gallery">
            <img src="images/photo1.jpg" alt="Photo 1">
            <img src="images/photo2.jpg" alt="Photo 2">
            <img src="images/photo3.jpg" alt="Photo 3">
          </div>
        </article>
  
        <article class="card">
          <h2>Персонал</h2>
          <ul>
            <li>Иван Иванов</li>
            <li>Мария Петрова</li>
            <li>Николай Сидоров</li>
          </ul>
        </article>
  
        <article class="card">
          <h2>Проекты</h2>
          <ul>
            <li>Проект 1</li>
            <li>Проект 2</li>
            <li>Проект 3</li>
          </ul>
        </article>
  
        <article class="card">
          <h2>Заказы</h2>
          <table>
            <thead>
              <tr>
                <th>№ заказа</th>
                <th>Дата</th>
                <th>Статус</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>12345</td>
                <td>15.05.2023</td>
                <td>Выполнен</td>
              </tr>
              <tr>
                <td>67890</td>
                <td>20.05.2023</td>
                <td>В процессе</td>
              </tr>
            </tbody>
          </table>
        </article>
      </section>
    </div>
  </template>
  
  <script>
  export default {
    name: 'App'
  }
  </script>
  
  <style scoped>
  .header {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1rem;
    background-color: #fff;
    border-bottom: 1px solid #ccc;
  }
  
  .container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
  
  .card {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    padding: 1rem;
    background-color: #fafafa;
    border: 1px solid #ccc;
  }
  
  .gallery {
    display: flex;
    gap: 1rem;
  }
  
  .gallery img {
    width: 100px;
    height: 100px;
    object-fit: cover;
  }
  </style>