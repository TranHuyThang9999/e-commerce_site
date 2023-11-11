// Chat.js
import React, { useState } from 'react';
import  useWebSocket, { ReadyState }  from 'react-use-websocket';

const Chat = () => {
    const [message, setMessage] = useState('');
    const [sendMessage, lastMessage, readyState] = useWebSocket('ws://localhost:8080/ws');

    const handleSendMessage = () => {
        sendMessage(message);
        setMessage('');
    };

    return (
        <div>
            <div>
                <input
                    type="text"
                    value={message}
                    onChange={(e) => setMessage(e.target.value)}
                />
                <button onClick={handleSendMessage}>Send</button>
            </div>
            <div>
                <p>WebSocket Connection Status: {ReadyState[readyState]}</p>
                <p>Last Message: {lastMessage ? lastMessage.data : 'No messages yet.'}</p>
            </div>
        </div>
    );
};

export default Chat;
