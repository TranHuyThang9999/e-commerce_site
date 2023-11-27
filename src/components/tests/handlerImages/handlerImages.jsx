import React, { useState } from 'react';
import axios from 'axios';
import "./hadlerImages.css"
import { Card, Image, Col, Row, Upload, Button, Input, Modal, message, Popconfirm } from 'antd';
import { UploadOutlined } from '@ant-design/icons';
import AddImage from './addImage';


function HandlerImages() {
    const [infoImages, setInfoImages] = useState([]);
    const [userId, setUserId] = useState('');
    // const [idImage, setIdImage] = useState('');

    const fetchData = async (id) => {
        try {
            const response = await axios.get(`http://localhost:8080/sell/imgbb?id_user=${id}`);
            setInfoImages(response.data);
        } catch (error) {
            console.error("Error fetching data:", error);
        }
    }

    const deleteImage = async (id) => {
        try {
            const response = await axios.delete(`http://localhost:8080/sell/img/delete?id_image=${id}`);
            console.log(response);
            if (response.data.result.code === 0) {
                message.success("Delete success");
                fetchData(userId); // Refetch data after successful deletion
            } else {
                message.error("Error server");
            }
        } catch (error) {
            console.error("Error deleting image:", error);
        }
    };

    const handleButtonClick = () => {
        fetchData(userId);
    };

    const handleDeleteByid = (id) => {
        deleteImage(id);
    }

    const [isModalOpen, setIsModalOpen] = useState(false);

    const showModal = () => {
        setIsModalOpen(true);
    };
    const handleOk = () => {
        setIsModalOpen(false);
    };
    const handleCancel = () => {
        setIsModalOpen(false);
    };

    const handleCreateSuccess = () => {
        // Xử lý khi tạo thành công, ví dụ: refetch dữ liệu hoặc cập nhật state
        fetchData(userId);

        // Sau khi xử lý xong, tự đóng modal
        handleOk();
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
                        <Button icon={<UploadOutlined />} className='button-modal-add-product' type="primary" onClick={showModal}>
                            Thêm sản phẩm
                        </Button>
                        <Modal
                            width={650}
                            footer
                            className='button-modal-add-product'
                            open={isModalOpen}
                            onOk={handleOk}
                            onCancel={handleCancel}
                        >
                            <AddImage onCreateSuccess={handleCreateSuccess} />

                        </Modal>
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
                                        <Popconfirm
                                            title="Delete the task"
                                            okText="Yes"
                                            cancelText="No"
                                            onConfirm={() => handleDeleteByid(image.id_image)}

                                        >
                                            <Button  danger>Delete</Button>
                                        </Popconfirm>
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
