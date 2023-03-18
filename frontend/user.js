const getUser = (event) => {
    event.preventDefault();
    socket.close();

    socket = new WebSocket("ws://localhost:8080/user");
    socket.addEventListener('open', () => {

        // const urlParams = new URLSearchParams(event.target.search);
        // const creator = urlParams.get('creator');
        // console.log(creator)
        socket.send(JSON.stringify({
            'token': getCookie('token')
        }));

        socket.addEventListener('message', async (event) => {
            const data = JSON.parse(event.data);
            if (data.error) {
                const newDiv = document.createElement('div').appendChild(document.createTextNode(data.error));
                document.body.appendChild(newDiv);
                return
            }
            console.log(data);
            event.target.href = "/profile"
            await route(event)


        });
    });
};
const getOtherUser = (event) => {
    event.preventDefault();
    socket.close();

    socket = new WebSocket("ws://localhost:8080/user/other");
    socket.addEventListener('open', () => {

        const urlParams = new URLSearchParams(event.target.search);
        const creator = urlParams.get('name');
        console.log(creator)
        socket.send(JSON.stringify({
            'username': creator
        }));

        socket.addEventListener('message', async (event) => {
            const data = JSON.parse(event.data);
            if (data.error) {
                const newDiv = document.createElement('div').appendChild(document.createTextNode(data.error));
                document.body.appendChild(newDiv);
                return
            }
            console.log(data);
            event.target.href = "/profile"
            await route(event)


        });
    });
};