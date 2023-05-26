const getUser = async (event) => {
    event.preventDefault();

    const urlParams = new URLSearchParams(event.target.search);
    const username = urlParams.get('username');
    console.log(username)
    await fetch('http://localhost:8080/profile?username=' + username, {
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
                event.target.href = "/profile?username=" + data.username
                await route(event)
                createProfileOnPage(data)
                getAllDialogs()

            } else {
                console.log(data.Text)
                route(event)
            }
        })
};
const getYouProfile = async (event) => {
    event.preventDefault();

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
                event.target.href = "/profile"
                await route(event)
                createProfileOnPage(data)
                getAllDialogs()
            } else {
                console.log(data.Text)
                route(event)
            }
        })
};
const createProfileOnPage = (profile) => {
    const prof = document.createElement('div');
    prof.className = 'profile'

    const username = document.createElement('a')
    username.className = 'username'
    username.href = 'profile?username=' + profile.username
    username.setAttribute('onclick', 'getUser(event)')
    username.appendChild(document.createTextNode(profile.username))

    const email = document.createElement('div')
    email.className = 'email'
    email.appendChild(document.createTextNode(profile.email))

    const chat = document.createElement('a')
    chat.className = 'chat'
    chat.href = 'chat?username=' + profile.username
    chat.setAttribute('onclick', 'startChatReq(event)')
    chat.appendChild(document.createTextNode('send message'))

    prof.appendChild(username)
    prof.appendChild(email)
    prof.appendChild(chat)

    const main = document.getElementById('main-page')
    main.appendChild(prof)
}
const createChatOnPage = () => {
    const chat = document.createElement('div');
    chat.className = 'chat'

    const log = document.createElement('div')
    log.id = 'log'

    const form = document.createElement('form');
    form.method = 'POST';
    // form.action = '/submit';


    const msg = document.createElement('input');
    msg.type = 'text';
    msg.id = 'msg';
    msg.placeholder = 'send msg';
    form.appendChild(msg);

    const submitButton = document.createElement('button');
    submitButton.type = 'submit';
    submitButton.textContent = 'Send';
    form.appendChild(submitButton);






    chat.appendChild(log)
    chat.appendChild(form)

    const main = document.getElementById('main-page')
    main.appendChild(chat)
}