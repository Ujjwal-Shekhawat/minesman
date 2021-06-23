import { Fragment, React, useContext, useEffect } from 'react'
import { Redirect } from 'react-router-dom'
import authContext from "./context/authContext"
import { socket } from './Socket';

function Console() {
    const wrapStyle = {
        overflow: 'scroll',
        overflowX: 'hidden'
    };
    const bgImages = [
        "https://www.wallpaperup.com/uploads/wallpapers/2012/02/19/1270/9fe3d70441dd3ea508a7f457f647ea9a.jpg",
        "https://www.pixelstalk.net/wp-content/uploads/2015/12/Minecraft-wallpaper-free-download-620x349.jpg",
        "https://cdn.wallpapersafari.com/45/66/IsiyQk.jpg",
        "https://cdn.statically.io/img/wallpaperaccess.com/full/171177.jpg"
    ];

    const authCtx = useContext(authContext)
    const { isAuth, authUser, username, logout } = authCtx

    useEffect(() => {
        // authUser()
        if (isAuth) {
            console.log("ggwp", isAuth)
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
                span.innerText = "server console > "
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
        }
    }, []);

    const sendCommand = (event) => {
        if (event.code === "Enter") {
            let message = document.getElementById("commandBox");
            console.log(message.value);
            socket.emit('notice', message.value);
            message.value = ""
        }
    }

    const Disconnect = () => {
        if (socket != null) { socket.emit('bye', "bye"); socket.close() }
        logout()
    }

    const changeBg = () => {
        let randNum = [Math.floor(Math.random() * bgImages.length)];
        return bgImages[randNum];
    }

    const bgStyle = {
        backgroundImage: `url(${changeBg()}`,
        backgroundPosition: 'center',
        backgroundRepeat: 'no-repeat',
        backgroundSize: 'cover',
        overflowY: 'hidden',
        overflowX: 'hidden',
    }

    return (
        <div>
            { (isAuth == true) ? <div style={bgStyle}>
                <main id="container">
                    <div id="terminal" style={wrapStyle}>
                        <section id="terminal__body">
                            <div id="terminal__prompt">
                            </div>
                        </section>
                        <div id='terminal__prompt'>
                            <label
                                style={{ color: "transparent", width: "auto", background: "rgba(0, 0, 0, 0.644)", color: "greenyellow", borderColor: "transparent", border: "thin", outline: "transparent", fontFamily: 'Ubuntu Mono', fontSize: "100%", fontWeight: "bold" }}>server console &gt; </label>
                            <input id="commandBox" type="text" placeholder="command" autoFocus onKeyDown={sendCommand} />
                        </div>
                    </div>
                    <button onClick={Disconnect}>logoff</button>
                </main>
            </div>
                :
                <Fragment>
                    <Redirect to='/' />
                </Fragment>
            }
        </div>
    )
}

export default Console;