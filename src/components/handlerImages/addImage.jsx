import { Button, Form, Input, InputNumber, Upload } from 'antd';
import React, { useState } from 'react';
import "./hadlerImages.css"
import axios from 'axios';



function AddImage({onCreateSuccess}) {

    const layout = {
        labelCol: {
            span: 8,
        },
        wrapperCol: {
            span: 8,
        },

    };


    const [fileList, setFileList] = useState([]);

    const onChange = ({ fileList }) => {
        setFileList(fileList);
    };


    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();

            formData.append('name', values.name);
            formData.append('age', values.age);
            formData.append('address', values.address);


            fileList.forEach((file) => {
                formData.append('image', file.originFileObj);
            });


            const response = await axios.post('http://localhost:8080/sell/add', formData);

            if (response.data.result.code === 0) {
                alert('Tạo sp thành công');
                if (onCreateSuccess) {
                    onCreateSuccess();
                }
                
                return;
            }
            else {
                alert('Loix'); // User created successfully
                return

            }
        } catch (error) {
            console.error(error);
            alert('Có lỗi xảy ra. Vui lòng thử lại.'); // Handle error
            return;
        }
    };
    const tailLayout = {
        wrapperCol: {
            offset: 0,
        },
    };
    return (
        <div>

            <Form
                className='form-login'
                {...layout}
                onFinish={handleFormSubmit}
            >
                <Form.Item
                    label="Name"
                    name="name"
                    rules={[{ required: true, message: 'Vui lòng nhập thông tin tên của bạn!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label="Age"
                    name="age"
                    rules={[{ required: true, message: 'Vui lòng nhập thông tin age của bạn!' }]}
                >
                    <InputNumber />
                </Form.Item>
                <Form.Item
                    label="Address"
                    name="address"
                    rules={[{ required: true, message: 'Vui lòng nhập thông tin address của bạn!' }]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label='Nhập ảnh mô tả sản phẩm'
                    className='form-row'
                    name='file'
                    rules={[
                        { required: true, message: 'Vui lòng tên sản phẩm của bạn!' },
                    ]}
                >
                    <Upload
                        maxCount={5}
                        fileList={fileList}
                        listType='picture-card'
                        accept='image/jpeg,image/png'
                        onChange={onChange}
                        beforeUpload={() => false} // Prevent auto-upload
                    >
                        {fileList.length < 5 && '+ Upload'}
                    </Upload>
                </Form.Item>

                <Form.Item>
                    <Button type="primary" htmlType='submit'>
                        Tạo sản phẩm
                    </Button>
                </Form.Item>

            </Form>
        </div>
    );
}

export default AddImage;