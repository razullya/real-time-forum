
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

function getAllPosts() {

    socket.close();

    socket = new WebSocket("ws://localhost:8080/post/all");
    socket.addEventListener('open', () => {
        socket.send('');
        socket.addEventListener('message', event => {
            const data = JSON.parse(event.data);

            if (data.error) {
                return;
            }
            prepairPosts()

            data.forEach(element => {
                createPostOnMainPage(element)
            });

        });
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

        socket.addEventListener('message', async (event) => {
            const data = JSON.parse(event.data);

            event.target.href = "/post"
            await route(event)

            createPostOnPage(data)

        });
    });
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
    creator.setAttribute('onclick', 'getPostById(event)')
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