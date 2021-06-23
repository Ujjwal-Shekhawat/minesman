import io from "socket.io-client";

let socket;

function X() {
    socket = io.connect('http://localhost:8080', { path: '/ws/', transports: ["websocket"] });
}

export { X, socket };

