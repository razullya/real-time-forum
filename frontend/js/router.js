const route = (event) => {
    event = event || window.event;
    event.preventDefault();
    window.history.pushState({}, '', event.target.href);
    handleLocation();
};

const routes = {
    404: '/pages/404.html',
    '/': '/pages/main/mainno.html',
    '/post': '/pages/post/post.html',//

    '/signup': '/pages/auth/signup.html',
    '/signin': '/pages/auth/signin.html',

    '/profile': '/pages/user/user.html',//

}
const routesAuth = {
    404: '/pages/404.html',
    '/': '/pages/main/main.html',
    '/post': 'post.html',//

    '/create': '/pages/post/create.html',
    '/profile': 'pages/user/selfprofile.html',//

}


const handleLocation = async () => {
    const path = window.location.pathname

    checkCookie()
        .then(async result => {
            // if (path === '/') {
            const route = routesAuth[path] || routesAuth[404]
            const html = await fetch(route).then((data) => data.text());
            document.getElementById('main-nav').innerHTML = html
            //     return
            // }

            // const route = routesAuth[path] || routesAuth[404]
            // const html = await fetch(route).then((data) => data.text());
            // document.getElementById('main-page').innerHTML += html

        })
        .catch(async error => {
            // if (path === '/' || path === '/signup' || path === '/signin') {
            const route = routes[path] || routes[404]
            const html = await fetch(route).then((data) => data.text());
            document.getElementById('main-nav').innerHTML = html
            // return
            // }

            // const route = routes[path] || routes[404]
            // const html = await fetch(route).then((data) => data.text());
            // document.getElementById('main-page').innerHTML += html
        });
}

window.onpopstate = handleLocation;
window.route = route;

handleLocation();