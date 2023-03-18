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

function getAllPostComments() {

    socket.close();

    socket = new WebSocket("ws://localhost:8080/post/comments");
    socket.addEventListener('open', () => {
        const idFromPost = document.getElementById('post-info')

        socket.send(JSON.stringify({ 'id': idFromPost.getAttribute('post_id') }));
        console.log('comment');
        socket.addEventListener('message', event => {
            const data = JSON.parse(event.data);

            if (data.error) {
                return;
            }

            console.log(data)
            data.forEach(element => {
                createCommentOnPostPage(element)
            });

        });
    })

}
function createCommentOnPostPage(comment) {
    const comCont = document.createElement('div');
    comCont.className = 'comment'

    const creator = document.createElement('a')
    creator.className = 'creator'
    creator.href = 'creator?name=' + comment.creator
    creator.setAttribute('onclick', 'getOtherUser(event)')////
    creator.appendChild(document.createTextNode(comment.creator))

    const description = document.createElement('div')
    description.className = 'description'
    description.appendChild(document.createTextNode(comment.text))

    comCont.appendChild(creator)
    comCont.appendChild(description)

    const comments = document.getElementById('comments')
    comments.appendChild(comCont)
}
