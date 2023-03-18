
const onCreatePost = async (event) => {
    event.preventDefault();
    const post_form = document.querySelector('#createpost_form')
    const formData = new FormData(post_form)
    await fetch('http://localhost:8080/post/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'title': formData.get('title'),
            'description': formData.get('description'),
            'category': formData.get('category'),
            'token': getCookie('token')
        })
    })
        .then(response => {
            console.log(response);
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {


            if (data) {
               console.log(data)
                // event.target.href = "/"
                // route(event)
            } else {
                console.log(data.Text)
                route(event)
            }
        })
}

function createPostOnMainPage(post) {
    const postCont = document.createElement('div');
    postCont.className = 'post'

    const title = document.createElement('a')
    title.className = 'title'
    title.href = 'post?id=' + post.id
    title.setAttribute('onclick', 'getPostById(event)')
    title.appendChild(document.createTextNode(post.title))

    const description = document.createElement('div')
    description.className = 'description'
    description.appendChild(document.createTextNode(post.description))

    const category = document.createElement('div')
    category.className = 'category'
    category.appendChild(document.createTextNode(post.category))

    postCont.appendChild(title)
    postCont.appendChild(description)
    postCont.appendChild(category)
    const posts = document.getElementById('posts')
    posts.appendChild(postCont)
}

async function getAllPosts() {
    await fetch('http://localhost:8080/post/all', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: ''
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {
            if (data) {
                prepairPosts()
                data.forEach(element => {
                    createPostOnMainPage(element)
                });
            }
        })

}
function prepairPosts() {
    const div = document.getElementById('main-page')

    if (div.querySelectorAll("#posts").length == 0) {

        const posts = document.createElement('div')
        posts.id = 'posts'
        div.appendChild(posts)
    } else {
        const posts = document.getElementById("posts")
        posts.innerHTML = ""
    }
}
const getPostById = async (event) => {
    event.preventDefault();

    const urlParams = new URLSearchParams(event.target.search);
    const id = urlParams.get('id');
    await fetch('http://localhost:8080/post', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'id': id
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
                event.target.href = "/post"
                await route(event)

                createPostOnPage(data)
                getAllPostComments()
            } else {
                console.log(data.Text)
                route(event)
            }
        })
};
function createPostOnPage(post) {

    const postCont = document.getElementById('post-info');


    const title = document.createElement('a')
    title.className = 'post__title'

    title.appendChild(document.createTextNode(post.title))

    const description = document.createElement('div')
    description.className = 'post__description'
    description.appendChild(document.createTextNode(post.description))

    const creator = document.createElement('a')
    creator.className = 'post__title'
    creator.href = 'profile?creator=' + post.creator
    creator.setAttribute('onclick', 'getOtherUser(event)')
    creator.appendChild(document.createTextNode(post.creator))

    const category = document.createElement('div')
    category.className = 'tags'
    post.category.forEach(element => {
        category.appendChild(document.createTextNode(element))
    });

    postCont.appendChild(title)
    postCont.appendChild(description)
    postCont.appendChild(creator)
    postCont.appendChild(category)
    postCont.setAttribute("post_id", post.id)
}