import React, { useState } from 'react';
import axios from 'axios';
import "./hadlerImages.css"
import { Card, Image, Col, Row, Upload, Button, Input } from 'antd';


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

                <Row>
                    <Col span={8}>
                        <Input
                            type="text"
                            id="userId"
                            value={userId}
                            onChange={(e) => setUserId(e.target.value)}
                        />
                    </Col>
                    <Col span={8}>
                        <Button onClick={handleButtonClick}>Fetch Data</Button>
                    </Col>
                    <Col span={8}>
                        <Button>Add</Button>
                    </Col>
                </Row>
            </div>
            <div>
                <Row gutter={16}>
                    {infoImages.map((image) => (
                        <Col flex={3}>
                            <Card

                                title={

                                    <Upload
                                        disabled={true}

                                        onPreview={() => false}  // Add this line

                                        beforeUpload={() => false}
                                        listType="picture-card"
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

                                <Row>
                                    <Col flex={6}>
                                        <div style={{
                                            overflow: 'hidden',
                                            whiteSpace: 'nowrap',
                                            textOverflow: 'ellipsis',
                                            maxWidth: '200px',  // Adjust the maximum width as needed
                                            lineHeight: '1.8rem',
                                            display: 'block',
                                        }}>
                                            id: {image.id_image}
                                        </div>
                                    </Col>

                                    <Col flex={3}>
                                        <Button>
                                            Delete
                                        </Button>
                                    </Col>
                                </Row>
                            </Card>
                        </Col>
                    ))}
                </Row>

            </div>
        </div>
    );
}

export default HandlerImages;
