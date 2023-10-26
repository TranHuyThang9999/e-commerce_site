import { Button, Form, Image, Input, Space } from 'antd';
import React, { useState } from 'react';
import "./HistoryChat.css"
import { SendOutlined } from '@ant-design/icons';



const tailLayout = {
    wrapperCol: {
        offset: 8,
        span: 16,
    },
};


function HistoryChat() {

    const [message, setMessage] = useState('');
    const [messages,setMessages]=useState([]);

 
   const handleSubmit=()=>{
    setMessages(prev =>[...prev,message]);
    setMessage('');
   }

    return (
        <div>
            <Form
                layout="horizontal"
                {...tailLayout}
                className='form-container'

            >

                <Form.Item
                >
                    <header>
                        <Image />
                    </header>
                    <Input.TextArea
                        className='form-TextArea'
                        value={messages}
                        onChange={e => setMessage(e.target.value)}

                    >
                    </Input.TextArea>
                    <Space.Compact
                        style={{
                            width: '100%',
                            height: '50'
                        }}
                    >
                        <Input
                            value={message}
                            onChange={e => setMessage(e.target.value)}
                        />
                        <Button style={{
                            height: '50px'
                        }} type="primary"
                        onClick={handleSubmit}
                        >
                            <SendOutlined />
                        </Button>
                    </Space.Compact>
                </Form.Item>

            </Form>
        </div>
    );
}

export default HistoryChat;