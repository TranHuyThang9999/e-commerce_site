import React, { useState } from 'react';
import axios from 'axios';
import "./hadlerImages.css"
import { Card, Image, Col, Row, Upload } from 'antd';


function HandlerImages() {
    const [infoImages, setInfoImages] = useState([]);
    const [userId, setUserId] = useState('');

    const fetchData = async (id) => {
        try {
            const response = await axios.get(`http://localhost:8080/sell/imgbb?id_user=${id}`);
            setInfoImages(response.data);
        } catch (error) {
            console.error("Error fetching data:", error);
        }
    }

    const handleButtonClick = () => {
        fetchData(userId);
    };

    return (
        <div>
            <h2>Get all</h2>
            <div>
                <label htmlFor="userId">Enter User ID: </label>
                <input
                    type="text"
                    id="userId"
                    value={userId}
                    onChange={(e) => setUserId(e.target.value)}
                />
                <button onClick={handleButtonClick}>Fetch Data</button>
            </div>
            <div>
                <Row gutter={16}>
                    {infoImages.map((image) => (
                        <Col span={8}>
                            <Card

                                title={
                                    <Upload
                                        disabled={true}

                                        onPreview={() => false}  // Add this line

                                        beforeUpload={() => false}
                                        listType="picture-card"
                                        showUploadList={false}
                                        style={{ width: '100%' }}
                                    >
                                        <Image
                                            alt=""
                                            src={image.url}
                                            style={{ width: '100%', height: 'auto' }}
                                        />
                                    </Upload>
                                }
                            >

                                Code Product : {image.id_image}
                            </Card>
                        </Col>
                    ))}
                </Row>

            </div>
        </div>
    );
}

export default HandlerImages;
