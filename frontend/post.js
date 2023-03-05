
const onCreatePost = (event) => {
    event.preventDefault();
    socket.close();
    console.log('ya tut')
    socket = new WebSocket("ws://localhost:8080/post/create")
    socket.addEventListener('open', () => {
        const post_form = document.querySelector('#createpost_form')
        const formData = new FormData(post_form)
        socket.send(JSON.stringify({
            'title': formData.get('title'),
            'description': formData.get('description'),
            'category': formData.get('category'),
            'token': getCookie('token')
        }))
    })
    socket.addEventListener('message', event => {
        const data = JSON.parse(event.data)

        if (data.error) {
            const newDiv = document.createElement('div').appendChild(document.createTextNode(data.error));
            document.body.appendChild(newDiv);
            return
        }
        event.target.href = "/"
        route(event)
    })
}