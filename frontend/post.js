
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






function createPostOnPage(post) {
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
                createPostOnPage(element)
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
