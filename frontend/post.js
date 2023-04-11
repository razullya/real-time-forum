
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


            if (data.success) {
                console.log(data)
                event.target.href = "/"
                route(event)
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

    const categories = document.createElement('div')
    categories.className = 'categories'
    for (const iterator of post.category) {
        const category = document.createElement('div')
        category.className = 'category'
        category.appendChild(document.createTextNode(iterator))
        categories.appendChild(category)
    }

    postCont.appendChild(title)
    postCont.appendChild(description)
    postCont.appendChild(categories)
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
    await fetch('http://localhost:8080/post?id=' + id, {
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
                event.target.href = "/post?id=" + id
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
    console.log(post);
    const postCont = document.getElementById('post-info');


    const title = document.createElement('a')
    title.className = 'post__title'

    title.appendChild(document.createTextNode(post.title))

    const description = document.createElement('div')
    description.className = 'post__description'
    description.appendChild(document.createTextNode(post.description))

    const creator = document.createElement('a')
    creator.className = 'post__title'
    creator.href = 'profile?username=' + post.creator
    creator.setAttribute('onclick', 'getUser(event)')
    creator.appendChild(document.createTextNode(post.creator))

    const categories = document.createElement('div')
    categories.className = 'categories'
    post.category.forEach(element => {
        const category = document.createElement('div')
        category.className = "category"
        category.appendChild(document.createTextNode(element))
        categories.appendChild(category)
    });

    // const like = document.createElement('a')
    // like.className = 'post__like'
    // like.href = '/post/like'
    // like.setAttribute('onclick', 'likePost(event)')
    // like.appendChild(document.createTextNode('LIKE'))

    // const dislike = document.createElement('a')
    // dislike.className = 'post__dislike'
    // dislike.href = '/post/dislike'
    // dislike.setAttribute('onclick', 'dislikePost(event)')
    // dislike.appendChild(document.createTextNode('DISLIKE'))

    // const count = document.createElement('div')
    // count.className = 'post__count'

    // const likes = document.createElement('div')
    // likes.id = 'likes'
    // likes.appendChild(document.createTextNode('0'))

    // const dislikes = document.createElement('div')
    // dislikes.id = 'dislikes'
    // dislikes.appendChild(document.createTextNode('0'))


    // count.appendChild(likes)
    // count.appendChild(dislikes)


    postCont.appendChild(title)
    postCont.appendChild(description)
    postCont.appendChild(creator)
    postCont.appendChild(categories)
    // postCont.appendChild(like)
    // postCont.appendChild(dislike)
    // postCont.appendChild(count)

    postCont.setAttribute("post_id", post.id)
}
const likePost = async (event) => {
    event.preventDefault();

    const searchParams = new URLSearchParams(window.location.search);
    const id = searchParams.get('id');
    await fetch('http://localhost:8080/post/like', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'id': id,
            'token': getCookie('token'),
        })
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {

            console.log(data);
            if (data.success) {

                const likes = document.getElementById('likes')
                likes.innerHTML = ''

                console.log(likes);


            } else {
                console.log(data.Text)
                route(event)
            }
        })
}
const dislikePost = async (event) => {
    event.preventDefault();

    const searchParams = new URLSearchParams(window.location.search);
    const id = searchParams.get('id');
    await fetch('http://localhost:8080/post/dislike', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'id': id,
            'token': getCookie('token'),
        })
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {

            console.log(data);
            if (data.success) {

                const likes = document.getElementById('dislikes')
                likes.innerHTML = ''

                console.log(likes);


            } else {
                console.log(data.Text)
                route(event)
            }
        })
}

// event.preventDefault();

// socket.close();
// console.log('ya tut')
// socket = new WebSocket("ws://localhost:8080/chat/username")
// socket.addEventListener('open', () => {
//     const post_form = document.querySelector('#createpost_form')
//     const formData = new FormData(post_form)
//     socket.send(JSON.stringify({
//         'title': formData.get('title'),
//         'description': formData.get('description'),
//         'category': formData.get('category'),
//         'token': getCookie('token')
//     }))
// })
// socket.addEventListener('message', event => {
//     const data = JSON.parse(event.data)

//     if (data.error) {
//         const newDiv = document.createElement('div').appendChild(document.createTextNode(data.error));
//         document.body.appendChild(newDiv);
//         return
//     }
//     event.target.href = "/"
//     route(event)
// })