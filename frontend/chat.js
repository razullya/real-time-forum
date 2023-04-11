
let socket;

const startChat = async (event) => {
    event.preventDefault();
    const urlParams = new URLSearchParams(event.target.search);
    const username = urlParams.get('username');
    var tokenChat = await checkChatToken(username)
    if (!tokenChat) {
        console.log('sorry unauth user')
        return
    }
    console.log('ya tutaaa')
    console.log(tokenChat)
    socket = new WebSocket(`ws://localhost:8080/chat/ws?token=${tokenChat}`);

    socket.addEventListener('open', async (event) => {
        event.target.href = "/chat"
        await route(event)
        console.log('WebSocket connection opened:', event);
    });

    socket.addEventListener('message', (event) => {
        console.log(event.data)
        // const datas = JSON.parse(event.data);
        // console.log(`WebSocket message received from ${data.sender} to ${data.recipient}:`, data.text);
        const messages = document.getElementById('messages');
        messages.innerHTML += `<p>${event.data}</p>`;
    });

    socket.addEventListener('close', (event) => {
        console.log('WebSocket connection closed:', event);
    });

    socket.addEventListener('error', (event) => {
        console.error('WebSocket error:', event);
    });
};

const sendMessage = async (event) => {
    event.preventDefault();
    const message = document.getElementById('message').value;
    var username
    await fetch('http://localhost:8080/profile?token=' + getCookie('token'), {
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


const checkChatToken = async (username) => {
    var token
    await fetch('http://localhost:8080/chat/check', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username: username, token: getCookie('token') })
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            }
            throw new Error('Network response was not ok')
        })
        .then(async data => {
            console.log(data)
            if (!data.status) {
                token = data.stoken
                return
            }
            return null
        })
    return token
}
