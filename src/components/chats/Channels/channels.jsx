import { Avatar, Form, Input, Space, Button } from 'antd';
import React, { useState } from 'react';
import "./channels.css"
import { SendOutlined } from '@ant-design/icons';

const { TextArea } = Input;

function Channels() {

    const [message, setMessage] = useState('');
    const [messages, setMessages] = useState([]);

    const handleSubmit = () => {
        setMessages(prev => [...prev, message.trim()]);
        setMessage('');
    }

  

    return (
        <div>
            <Form className='form-container-textArea'
            >
                <Form.Item style={{ height: '32px' }} className='form-header-message'>
                    <Space size={10} wrap>
                        <div style={{ paddingLeft: '5px', paddingTop: '5px' }}>
                            <Avatar src={'https://gamek.mediacdn.vn/133514250583805952/2021/3/19/ce1-16161283334951633838261.jpg'} />

                        </div>
                        <Space
                            direction="vertical"
                            size="middle"
                            style={{
                                display: 'flex',
                            }}
                            className='form-name-and-status'>
                            <Form.Item className='form-name'>
                                <span style={{ fontWeight: 'bold' }}>
                                    Trần Huy Thắng
                                </span><br />
                                Đang hoạt động
                            </Form.Item>

                        </Space>
                    </Space>

                </Form.Item>

                <Form.Item>

                    <TextArea
                        style={{ color: 'blueviolet' }}
                        value={messages.join('\n')}
                         />
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
                            onClick={() => {
                                handleSubmit();
                            }}
                            disabled={message.trim() === ''}
                        >
                            <SendOutlined />
                        </Button>
                    </Space.Compact>
                </Form.Item>
            </Form>
        </div>
    );
}


export default Channels;