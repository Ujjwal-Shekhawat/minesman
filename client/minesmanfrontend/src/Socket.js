import io from "socket.io-client";

let socket;

function X() { 
    socket = io.connect('http://20.197.57.10:3000', {path: '/ws/', transports: ["websocket"]});
    return socket
}

export default X;

