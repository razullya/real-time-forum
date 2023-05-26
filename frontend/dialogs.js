
let socket;

// const startChat = async (event) => {
//     event.preventDefault();
//     const urlParams = new URLSearchParams(event.target.search);
//     const username = urlParams.get('username');
//     var tokenChat = await checkChatToken(username)
//     if (!tokenChat) {
//         console.log('sorry unauth user')
//         return
//     }

//     console.log(tokenChat)
//     socket = new WebSocket(`ws://localhost:8080/chat/ws?token=${tokenChat}`);

//     socket.addEventListener('open', async (event) => {
//         event.target.href = `/chat?token=${tokenChat}`
//         await route(event)
//         console.log('WebSocket connection opened:', event);
//     });

//     socket.addEventListener('message', (event) => {
//         console.log(event.data)
//         // const datas = JSON.parse(event.data);
//         // console.log(`WebSocket message received from ${data.sender} to ${data.recipient}:`, data.text);
//         const messages = document.getElementById('messages');
//         messages.innerHTML += `<p>${event.data}</p>`;
//     });

//     socket.addEventListener('close', (event) => {
//         console.log('WebSocket connection closed:', event);
//     });

//     socket.addEventListener('error', (event) => {
//         console.error('WebSocket error:', event);
//     });
// };

const sendMessage = async (event) => {
    event.preventDefault();
    const message = document.getElementById('message').value;
    var username
    await fetch('http://localhost:8080/chat?token=' + getCookie('token'), {
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

            console.log(data);
            if (data) {
                username = data.username

            } else {
                route(event)
            }
        })

    if (socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify({ text: message, sender: username }));
    } else {
        console.error('WebSocket connection is not open');
    }
};


function prepairDialogs() {
    const div = document.getElementById('main-page')

    if (div.querySelectorAll("#dialogs").length == 0) {
        const dialogs = document.createElement('div')
        dialogs.id = 'dialogs'
        div.appendChild(dialogs)
    } else {
        const dialogs = document.getElementById("dialogs")
        dialogs.innerHTML = ""
    }
}

async function getAllDialogs() {
    console.log('ya tut?')
    prepairDialogs()
    await fetch('http://localhost:8080/chat/all?token=' + getCookie('token'), {
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
                console.log(data, 'dataaaaa')
                data.forEach(element => {
                    createDialogsOnPage(element)
                });

            }
        })

}
function createDialogsOnPage(dialog) {

    const dialogCont = document.createElement('div')
    dialogCont.className = 'dialog'


    const creator = document.createElement('a')
    creator.className = 'creator'
    creator.href = 'profile?username=' + dialog.username
    creator.setAttribute('onclick', 'getUser(event)')////
    creator.appendChild(document.createTextNode(dialog.sender))


    const chat = document.createElement('a')
    chat.className = 'creator'
    chat.href = 'chat/start?username=' + dialog.sender
    chat.setAttribute('onclick', 'startChat(event)')////
    chat.appendChild(document.createTextNode(dialog.sender))


    dialogCont.appendChild(creator)
    dialogCont.appendChild(chat)

    const dialogs = document.getElementById('dialogs')
    dialogs.appendChild(dialogCont)
}


const startChat = async (event) => {
    event.preventDefault();
    console.log('ya startanul')
    const urlParams = new URLSearchParams(event.target.search);
    const username = urlParams.get('username');
    await fetch('http://localhost:8080/chat/start', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'username': username,
            'token': getCookie('token')
        })
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {


            if (!data.status) {
                event.target.href = "/chat?token=" + data.token
                await route(event)

                // create chat view

                // createPostOnPage(data)
                // getAllPostComments()
            } else {
                console.log(data.text)
                route(event)
            }
        })
};

//waiting
const startChatReq = async (event) => {
    event.preventDefault();

    const urlParams = new URLSearchParams(event.target.search);
    const username = urlParams.get('username');
    await fetch('http://localhost:8080/chat/req', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'username': username,
            'token': getCookie('token')
        })
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {


            if (!data.status) {
                console.log(data, "uraaa token poluchil")

                event.target.href = "/chat?token=" + data.token
                await route(event)
                console.log(data.token, "uraaa token poluchil")
                // create chat view

                // createPostOnPage(data)
                // getAllPostComments()
            } else {
                console.log(data.text)
                route(event)
            }
        })
};

