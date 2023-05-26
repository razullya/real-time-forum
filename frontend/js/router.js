const route = async (event) => {
    event = event || window.event;
    event.preventDefault();
    window.history.pushState({}, '', event.target.href);
    await handleLocation(event);
};

const routes = {
    404: '/pages/404.html',
    '/': '/pages/main/mainno.html',

    '/post': '/pages/post/post.html',//

    '/signup': '/pages/auth/signup.html',
    '/signin': '/pages/auth/signin.html',

    '/profile': '/pages/user/profile.html',
    // '/chat': '/pages/user/chat.html'

}
const routesAuth = {
    404: '/pages/404.html',
    '/': '/pages/main/main.html',
    '/post': '/pages/post/post.html',//

    '/create': '/pages/post/create.html',
    '/profile': 'pages/user/profile.html',//
    '/dialogs': 'pages/user/dialogs.html',

    '/chat': '/pages/user/chat.html'

}


const handleLocation = async (event) => {
    const path = window.location.pathname
    const token = checkCookie('token');

    await fetch('http://localhost:8080/token', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ token: token })
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok');
        })
        .then(async data => {

            if (data.success) {
                console.log('authorized');
                console.log(path)
                if (path === '/') {
                    document.getElementById('main-page').innerHTML = ""
                    getAllPosts()
                    const route = routesAuth[path] || routesAuth[404]
                    const html = await fetch(route).then((data) => data.text());
                    document.getElementById('main-nav').innerHTML = html

                    return
                }
                if (path === '/create' || path === '/profile') {
                    const route = routesAuth[path] || routesAuth[404]
                    const html = await fetch(route).then((data) => data.text());
                    document.getElementById('main-page').innerHTML = html
                    return
                }
                if (path === '/post') {

                    const route = routesAuth[path] || routesAuth[404]
                    const html = await fetch(route).then((data) => data.text());
                    document.getElementById('main-page').innerHTML = html
                    return
                }
                if (path === '/dialogs') {

                    const route = routesAuth[path] || routesAuth[404]
                    const html = await fetch(route).then((data) => data.text());
                    document.getElementById('main-page').innerHTML = html
                    return
                }
                if (path === '/chat') {
                    const route = routesAuth[path] || routesAuth[404]
                    const html = await fetch(route).then((data) => data.text());
                    document.getElementById('main-page').innerHTML = html
                    return
                }

                const route = routesAuth[path] || routesAuth[404]
                const html = await fetch(route).then((data) => data.text());
                document.getElementById('main-page').innerHTML += html

            } else {
                console.log('unauthorized');

                if (path === '/') {
                    document.getElementById('main-page').innerHTML = ""
                    getAllPosts()

                    const route = routes[path] || routes[404]
                    const html = await fetch(route).then((data) => data.text());
                    document.getElementById('main-nav').innerHTML = html
                    return
                }

                if (path === '/signup' || path === '/signin' || path === '/post') {
                    const route = routes[path] || routes[404]
                    const html = await fetch(route).then((data) => data.text());
                    document.getElementById('main-page').innerHTML = html
                    return
                }
                const route = routes[path] || routes[404]
                const html = await fetch(route).then((data) => data.text());
                document.getElementById('main-page').innerHTML += html
            }
        })
        .catch(error => {
            console.error('Произошла ошибка:', error);
        });




}

window.onpopstate = handleLocation;
window.route = route;

handleLocation();
