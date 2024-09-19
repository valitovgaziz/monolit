document.addEventListener('DOMContentLoaded', () => {
    const authForm = document.getElementById('authForm');

    authForm.addEventListener('submit', async e => {
        e.preventDefault();

        const formData = new FormData(authForm);
        const phoneNumber = formData.get('phoneNumber');
        const password = formData.get('password');

        try {
            await fetch('/api/login', {
                method: 'POST',
                body: JSON.stringify({ phoneNumber, password }),
                headers: {
                    'Content-Type': 'application/json'
                },
            });

            if (response.ok) {
                // Успешная авторизация
                alert('Вы успешно вошли!');
                window.location.href = '/home';
            } else {
                throw new Error('Неправильный номер телефона или пароль');
            }
        } catch (error) {
            alert(`Ошибка при авторизации: ${error}`);
        }
    });
});