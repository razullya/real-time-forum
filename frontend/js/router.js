const route = (event) => {
    event = event || window.event;
    event.preventDefault();
    window.history.pushState({}, '', event.target.href);
    handleLocation();
};
getAllPosts()
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


const handleLocation = async () => {
    const path = window.location.pathname

    checkCookie()
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
                getPostById()
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


// setInterval(getAllPosts, 2000);


const getPostById = (event) => {
    event.preventDefault();
    socket.close();


    socket = new WebSocket("ws://localhost:8080/post");
    socket.addEventListener('open', () => {

        const urlParams = new URLSearchParams(event.target.search);
        const id = urlParams.get('id');
        console.log(id)
        socket.send(JSON.stringify({
            'id': id
        }));

        socket.addEventListener('message', event => {
            const data = JSON.parse(event.data);

            console.log(data)

            event.target.href = "/post"
            route(event)
        });
    });
};

