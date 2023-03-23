function getCookie(name) {
    const value = "; " + document.cookie;
    const parts = value.split("; " + name + "=");
    if (parts.length === 2) {
        return parts.pop().split(";").shift();
    }
}

function getNameCookie(name) {
    const cookies = document.cookie.split("; ");
    console.log(cookies)
    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].split("=");
        if (cookie[0] === name) {
            return cookie[1];
        }
    }
    return "";
}

function getCookieExpirationDate(cookieName) {

    return new Date(document.cookie
        .split('; ')
        .find(row => row.startsWith('expires='))
        ?.split('=')[1]);

}

function checkCookie(cookieName) {
    const expirationDate = getCookieExpirationDate(cookieName);
    // console.log(expirationDate);
    if (Date.now() > expirationDate) {
        console.log(cookieName, expirationDate);
        deleteCookie(cookieName)

        return ""
    } else {
        return getCookie(cookieName)
    }
}

function deleteCookie(name) {
    document.cookie = name + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}

const onSignUpSubmit = async (event) => {
    event.preventDefault();
    const signUp = document.querySelector('#signup_form');
    const formData = new FormData(signUp);
    await fetch('http://localhost:8080/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'email': formData.get('email'),
            'username': formData.get('username'),
            'password': formData.get('password')
        })
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {
            console.log(data);
            if (data.success) {
                event.target.href = "/"
                route(event)
            } else {
                console.log(data.Text)
                route(event)
            }
        })
};

const onSignInSubmit = async (event) => {
    event.preventDefault();
    event.preventDefault();
    const signUp = document.querySelector('#signin_form');
    const formData = new FormData(signUp);
    await fetch('http://localhost:8080/signin', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'username': formData.get('username'),
            'password': formData.get('password')
        })
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {
            if (data.token) {
                var date = new Date();
                date.setTime(date.getTime() + (24 * 60 * 60 * 1000));

                document.cookie = "token=" + data.token + "; expires=" + date.toUTCString() + "; SameSite=None; Secure";

                event.target.href = "/"
                route(event)
            } else {
                console.log(data.Text)
                route(event)
            }
        })

};


const onLogOut = async (event) => {
    event.preventDefault();

    await fetch('http://localhost:8080/logout', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'token': getCookie("token"),
        })
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {
            deleteCookie("token");
            event.target.href = "/"
            route(event)
        })

};

