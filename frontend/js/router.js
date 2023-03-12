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

    '/profile': '/pages/user/user.html',//

}
const routesAuth = {
    404: '/pages/404.html',
    '/': '/pages/main/main.html',
    '/post': '/pages/post/post.html',//

    '/create': '/pages/post/create.html',
    '/profile': 'pages/user/selfprofile.html',//

}


const handleLocation = async (event) => {
    const path = window.location.pathname

    await checkCookie()
        .then(async result => {
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

            const route = routesAuth[path] || routesAuth[404]
            const html = await fetch(route).then((data) => data.text());
            document.getElementById('main-page').innerHTML += html

        })
        .catch(async error => {

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
        });
}

window.onpopstate = handleLocation;
window.route = route;

handleLocation();


