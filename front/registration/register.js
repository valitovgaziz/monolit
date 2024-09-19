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
            await fetch('/api/register', {
                method: 'POST',
                body: JSON.stringify({ phoneNumber, role, password }),
                headers: {
                    'Content-Type': 'application/json'
                },
            });
            
            window.location.href = '/login';
        } catch (error) {
            console.log('Ошибка при отправке данных: ', error);
        }
    });
});