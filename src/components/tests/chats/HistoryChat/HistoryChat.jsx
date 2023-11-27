import React from 'react';
import "./HistoryChat.css"
import Channels from '../Channels/channels';



function HistoryChat() {


    return (
        <div>

            <div className='form-container-channel'>
                <div className='channels-container'>
                <Channels />
                </div>
            </div>

        </div>
    );
}

export default HistoryChat;