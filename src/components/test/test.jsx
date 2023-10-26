import { Button, Input } from "antd";
import React, { useRef, useState } from "react";
import './test.css'
function TestApp() {
    const [message, setMessage] = useState('');
    const [messages, setMessages] = useState([]);

    const handleSubmit = () => {
        setMessages(prev => [...prev, message.trim()]);
        setMessage('');
    }

    return (
        <>
            <Input value={message}
                onChange={e => setMessage(e.target.value)} />
            <Button onClick={handleSubmit}>Send</Button>
            <ul>
                {messages.map((message) => (
                    <li key={message}>{message}</li>
                ))}
           </ul>
        </>
    )
}



export default TestApp;

// ul {
//     list-style-type: none;
//     padding: 0;
// }

// li {
//     display: block;
//     padding: 10px;
//     margin-bottom: 5px;
//     background-color: #ffffff;
// }