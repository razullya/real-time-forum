const onCreateComment = (event) => {
    event.preventDefault();
    socket.close();
    socket = new WebSocket("ws://localhost:8080/comment/create")
    socket.addEventListener('open', () => {
        const comment_form = document.querySelector('#createcomment_form')
        const formData = new FormData(comment_form)
        const idFromPost = document.getElementById('post-info')

        socket.send(JSON.stringify({
            'id': idFromPost.getAttribute('post_id'),
            'token': getCookie('token'),
            'description': formData.get('form_auth_input'),
        }))
    })
    socket.addEventListener('message', event => {
        const data = JSON.parse(event.data)

        if (data.error) {
            const newDiv = document.createElement('div').appendChild(document.createTextNode(data.error));
            document.body.appendChild(newDiv);
            return
        }
        const comment_form = document.getElementById('form_auth_input')

        comment_form.value = ""
    })
}