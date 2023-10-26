import { Button, Form, Image, Input, Space } from 'antd';
import React, { useState, useEffect, useReducer } from 'react';
import "./HistoryChat.css"
import { SendOutlined } from '@ant-design/icons';
const { TextArea } = Input;



const tailLayout = {
    wrapperCol: {
        offset: 8,
        span: 16,
    },
};


function HistoryChat() {
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
                        <Image  />
                    </header>
                    <TextArea
                        className='form-TextArea'
                        value={show ? message : ''}
                    >
                    </TextArea>
                    <Space.Compact
                        style={{
                            width: '100%',
                            height: '50'
                        }}
                    >
                        <Input
                            onChange={(e) => setMessage(e.target.value)}

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