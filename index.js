setInterval(function () {
    //user_token ,to_user_token message
    let obj = {
        user_token:'ca7c99cb-2120-4f3f-9a2b-a2d045848e89',
        to_user_token:'efb744d3-5ff5-462a-808b-ae3d1bdf4f96',
        message:"xxx",
    }
    websocket.send(JSON.stringify(obj))
}, 3000)
