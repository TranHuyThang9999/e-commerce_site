import { Form } from 'antd';
import React from 'react';
import "./HistoryChat.css"
import Channels from '../Channels/channels';




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
                className='form-container'>

                <Form.Item>

                    <Channels className='form-TextArea' />

                </Form.Item>

            </Form>
        </div>
    );
}

export default HistoryChat;