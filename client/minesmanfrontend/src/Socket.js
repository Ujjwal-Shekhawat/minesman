import io from "socket.io-client";

let socket;

function X() {
    socket = io.connect('https://20.193.246.52', { path: '/ws/', transports: ["websocket"] });
}

export { X, socket };

