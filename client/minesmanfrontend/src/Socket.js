import io from "socket.io-client";

let socket;

function X() {
    socket = io.connect('http://20.197.57.10:8080', { path: '/ws/', transports: ["websocket"] });
}

export { X, socket };

