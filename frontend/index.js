// const myDiv = document.getElementById('posts');
// const comment = document.createElement('Это мой комментарий');
// myDiv.appendChild(comment);
// const newDiv = document.createElement('div').appendChild(document.createTextNode(data.error));
// document.body.appendChild(newDiv);

const div = document.getElementById('main-page')




getAllPosts()
function createPostOnPage(post) {
    const postCont = document.createElement('div');
    postCont.className = 'post'

    const title = document.createElement('div')
    title.className = 'title'
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

    div.appendChild(postCont)
}
function getAllPosts() {

    // event.preventDefault();
    // socket.close();

    // socket = new WebSocket("ws://localhost:8080/post/all");
    // socket.addEventListener('open', () => {
    //     socket.send('');
    //     socket.addEventListener('message', event => {
    //         const data = JSON.parse(event.data);

    //         if (data.error) {
    //             // const newDiv = document.createElement('div').appendChild(document.createTextNode(data.error));
    //             // document.body.appendChild(newDiv);
    //             return;
    //         }

    //         data.forEach(element => {
    //             createPostOnPage(element)
    //         });
    //     });
    // })

}