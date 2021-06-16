import io from "socket.io-client";

let socket;

function X() { 
    socket = io.connect('ws://localhost:8080', {transports: ["websocket"]});
    return socket
}

export default X;

