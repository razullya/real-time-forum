const onCreateComment = async (event) => {
    event.preventDefault();

    const comment_form = document.querySelector('#createcomment_form')
    const formData = new FormData(comment_form)
    const idFromPost = document.getElementById('post-info')
    const searchParams = new URLSearchParams(window.location.search);
    const id = searchParams.get('id');
    await fetch('http://localhost:8080/comment/create?id=' + id, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'id': idFromPost.getAttribute('post_id'),
            'token': getCookie('token'),
            'description': formData.get('form_auth_input'),
        })
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {


            if (data) {
                const comment_form = document.getElementById('form_auth_input')
                comment_form.value = ""
                getAllPostComments()

            } else {
                console.log(data.Text)
                route(event)
            }
        })
}
function prepairComments() {
    const div = document.getElementById('main-page')

    if (div.querySelectorAll("#comments").length == 0) {
        const comments = document.createElement('div')
        comments.id = 'comments'
        div.appendChild(comments)
    } else {
        const comments = document.getElementById("comments")
        comments.innerHTML = ""
    }
}
async function getAllPostComments() {
    prepairComments()
    const idFromPost = document.getElementById('post-info')
    const searchParams = new URLSearchParams(window.location.search);
    const id = searchParams.get('id');
    await fetch('http://localhost:8080/post/comments?id=' + id, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {


            if (data) {
                console.log(data)
                data.forEach(element => {
                    createCommentOnPostPage(element)
                });

            } else {
                console.log(data.Text)

            }
        })

}
function createCommentOnPostPage(comment) {

    
}
