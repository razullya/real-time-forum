let socket = new WebSocket("ws://localhost:8080/signup")

function getCookie(name) {
    const value = "; " + document.cookie;
    const parts = value.split("; " + name + "=");
    if (parts.length === 2) {
        return parts.pop().split(";").shift();
    }
}
function deleteCookie(name) {
    document.cookie = name + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}

const checkCookie = () => {
    return new Promise((resolve, reject) => {
        const socket = new WebSocket("ws://localhost:8080/token");
        socket.addEventListener('open', () => {
            const token = getCookie("token");
            if (!token) {
                reject('no token');
            } else {
                socket.send(JSON.stringify({ 'token': token }));
                socket.addEventListener('message', (event) => {
                    const data = JSON.parse(event.data);
                    if (data.error) {
                        reject(data.error);
                    } else {
                        resolve('ok');
                    }
                    socket.close();
                });
            }
        });
    });
};

const onSignUpSubmit = (event) => {
    event.preventDefault();
    socket.close();

    socket = new WebSocket("ws://localhost:8080/signup");
    socket.addEventListener('open', () => {
        const signUp = document.querySelector('#signup_form');
        const formData = new FormData(signUp);

        socket.send(JSON.stringify({
            'email': formData.get('email'),
            'username': formData.get('username'),
            'password': formData.get('password')
        }));

        socket.addEventListener('message', event => {
            const data = JSON.parse(event.data);

            if (data.error) {
                const newDiv = document.createElement('div').appendChild(document.createTextNode(data.error));
                document.body.appendChild(newDiv);
                return;
            }
            event.target.href = "/"
            route(event)
        });
    })

};

const onSignInSubmit = (event) => {
    event.preventDefault();
    socket.close();

    socket = new WebSocket("ws://localhost:8080/signin");
    socket.addEventListener('open', () => {
        const signIn = document.querySelector('#signin_form');
        const formData = new FormData(signIn);

        socket.send(JSON.stringify({
            'username': formData.get('username'),
            'password': formData.get('password')
        }));

        socket.addEventListener('message', event => {
            const data = JSON.parse(event.data);

            if (data.error) {
                // const newDiv = document.createElement('div').appendChild(document.createTextNode(data.error));
                // document.body.appendChild(newDiv);
                return;
            }
            var date = new Date();
            date.setTime(date.getTime() + ( 24 * 60 * 60 * 1000));
            document.cookie = "token=" + data.token + "; expires=" + date.toUTCString() + "; SameSite=None; Secure";

            event.target.href = "/"
            route(event)
        });
    });
    socket.close();

};


const onLogOut = (event) => {
    event.preventDefault();
    socket.close();

    socket = new WebSocket("ws://localhost:8080/logout");
    socket.addEventListener('open', () => {

        socket.send(JSON.stringify({
            'token': getCookie("token"),
        }));

        socket.addEventListener('message', event => {
            // const data = JSON.parse(event.data);
            deleteCookie("token");

            event.target.href = "/"
            route(event)
        });
    });
};

