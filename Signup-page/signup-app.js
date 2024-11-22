
function toggleForm(form) {
    if (form === 'login') {
        document.getElementById('login-container').style.display = 'block';
        document.getElementById('signup-container').style.display = 'none';
    } else {
        document.getElementById('login-container').style.display = 'none';
        document.getElementById('signup-container').style.display = 'block';
    }
}

async function login() {
    const username = document.getElementById('login-username').value;
    const password = document.getElementById('login-password').value;

    const response = await fetch('/api/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
    });

    const result = await response.json();
    if (result.success) {
        alert('Login successful');
    } else {
        alert('Incorrect credentials');
    }
}

async function signup() {
    const username = document.getElementById('signup-username').value;
    const password = document.getElementById('signup-password').value;

    const response = await fetch('/api/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
    });

    const result = await response.json();
    if (result.success) {
        alert('Signup successful');
        toggleForm('login');
    } else {
        alert('Signup failed');
    }
}

async function forgotPassword() {
    const username = prompt('Enter your phone number:');
    if (username) {
        const response = await fetch('/api/forgot-password', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username })
        });

        const result = await response.json();
        if (result.success) {
            alert('Password reset link sent via SMS');
        } else {
            alert('Failed to send password reset link');
        }
    }
}