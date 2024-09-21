// register.js
document.addEventListener('DOMContentLoaded', () => {
    const registrationForm = document.getElementById('registrationForm');

    // Обрабатываем событие отправки формы
    registrationForm.addEventListener('submit', async e => {
        e.preventDefault();

        const formData = new FormData(registrationForm);
        const phoneNumber = formData.get('phoneNumber');
        const role = formData.get('role');
        const password = formData.get('password');

        try {
            await fetch('http://localhost:8000/signupShort', {
                method: 'POST',
                body: JSON.stringify({ phoneNumber, role, password }),
                headers: {
                    'Content-Type': 'application/json',
                    'Access-Control-Allow-Origin': '*', // Разрешить любой домен
                    'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS', // Разрешенные методы
                    'Access-Control-Allow-Headers': 'X-Requested-With, Content-Type, Authorization' // Разрешенные заголовки
                },
            });
            
            window.location.href = '../auth/auth.html';
        } catch (error) {
            console.log('Ошибка при отправке данных: ', error);
        }
    });
});