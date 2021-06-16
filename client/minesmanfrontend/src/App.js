import { useEffect } from 'react'
import "./App.css"
import X from './Socket';

function App() {
  const wrapStyle = {
    overflow: 'scroll',
    overflowX: 'hidden'
  };

  useEffect(() => {
    let socket = X();

    function sendMessage() {
      let message = document.getElementById("commandBox");
      console.log(message.value);
      socket.emit('notice', message.value);
      message.value = ""

    }

    function reconnect() {
      if (socket.disconnected) {
        socket.connect();
        console.log("Reconnected")
      }
    }

    socket.on('connect', () => {
      console.log('Connected');
    })

    socket.on('connect_error', (err) => {
      console.clear();
      console.log('Error connecting to server');
      console.log(err);
      socket.disconnect();
    })

    socket.on('reply', (msg) => {
      console.log(msg);
      let section = document.getElementById("terminal__body");

      let div = document.createElement('div');
      div.id = "terminal__prompt";
      let span = document.createElement('span');
      span.id = "terminal__prompt--user"
      span.innerText = "spex@opcarm64: "
      div.append(span)

      span = document.createElement('span');
      span.id = "terminal__prompt--location"
      span.innerText = "~ "
      div.append(span)

      span = document.createElement('span');
      span.id = "terminal__prompt--bling"
      span.innerText = "$ "
      div.append(span)

      span = document.createElement('span');
      span.id = "terminal__prompt--bling"
      span.innerText = msg
      span.style.color = "white"
      div.append(span)
      section.append(div)

      updateScroll()
    })

    function updateScroll() {
      let element = document.getElementById("terminal");
      element.scrollTop = element.scrollHeight;
    }

    document.addEventListener("DOMContentLoaded", () => {
      let commandInput = document.getElementById("commandBox")
      commandInput.addEventListener('keydown', (k) => {
        if (k.code === "Enter") {
          sendMessage();
        }
      })
    })
  });

  return (
    <div>
      <main id="container">
        <div id="terminal" style={wrapStyle}>
          <section id="terminal__body">
            <div id="terminal__prompt">
            </div>
          </section>
          <div id='terminal__prompt'>
            <label
              style={{ color: "transparent", width: "auto", background: "rgba(0, 0, 0, 0.644)", color: "greenyellow", borderColor: "transparent", border: "thin", outline: "transparent", fontFamily: 'Ubuntu Mono', fontSize: "100%", fontWeight: "bold" }}>spex@opcarm64: <span style={{ color: "lightblue", fontWeight: "normal", }}>~ </span><span style={{ color: "white", fontWeight: "normal" }}>$ </span></label>
            <input id="commandBox" type="text" placeholder="command" autoFocus />
          </div>
        </div>
      </main>
    </div >
  );
}

export default App;
