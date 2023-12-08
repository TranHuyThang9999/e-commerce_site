import { Button, Form, Input, InputNumber, Select, Upload, message } from 'antd';
import React, { useState } from 'react';
import axios from 'axios';


function AddProducts() {

    const [form] = Form.useForm();

    const [fileList, setFileList] = useState([]);

    const onChange = ({ fileList }) => {
        setFileList(fileList);
    };

    

    const options = [
        {
            value: '19',
            label: 'Mở bán',
        },
        {
            value: '21',
            label: 'Đóng bán',
        },
    ];

    const handleFormSubmit = async (values) => {
        try {

            var formData = new FormData();

            formData.append("user_name", "thangth1");
            formData.append("name_product", values.name_product);
            formData.append("quantity", values.quantity);
            formData.append('sell_status', values.sell_status ? values.sell_status : '19');
            formData.append('price', values.price);
            formData.append("discount", values.discount);
            formData.append("manufacturer",values.manufacturer);
            formData.append("describe",values.describe);

            fileList.forEach((file) => {
                formData.append('image', file.originFileObj);
            });

            const response = await axios.post(
                'http://localhost:8080/sell/user/product/add',
                formData,
                {
                    headers: {
                        Authorization: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDIxMzk5ODgsImlkIjo0MzU3MjU3MTg3MTI1NzYsInVzZXJfbmFtZSI6InRoYW5ndGgxIiwicm9sZSI6MTcwMjA1MzU4OH0.GFVbOFgVNc66znoPl5K4xjvSdNmFN6WJhs6vuLrrW6k",
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );
            if (response.data.result.code === 0) {
                message.success("tạo sản phẩm thành công")
            } else if (response.data.result === 1) {

            }
        } catch (error) {
            message.error("error server")
        }
    }
    const layout = {
        labelCol: {
            span: 8,
        },
        wrapperCol: {
            span: 16,
        },
    };


    return (
        <div>
            <Form
                {...layout}
                form={form}
                className='form-container'
                onFinish={handleFormSubmit}
            >
                <Form.Item
                    label='Tên sản phẩm'
                    name='name_product'
                >
                    <Input/>
                </Form.Item>

                <Form.Item
                    label='Số  lượng sản phẩm'
                    name='quantity' 
                >
                    <InputNumber/>
                </Form.Item>
                <Form.Item
                    label='Giá sản phẩm'
                    name='price' 
                >
                    <InputNumber/>
                </Form.Item>

                <Form.Item
                    label='Mô tả sản phẩm'
                    name='describe' 
                >
                    <InputNumber/>
                </Form.Item>

                <Form.Item
                    label='Hàng sản phẩm'
                    name='manufacturer' 
                >
                    <Input/>
                </Form.Item>
                <Form.Item
                    label='Mở bán sản phẩm'
                    className='form-row'
                    name='sell_status'
                >
                    <Select
                        defaultValue="19"
                        style={{
                            width: 120,
                            height: 42,
                        }}
                        options={options}
                    />
                </Form.Item>
                <Form.Item
                    label='Giảm giá sản phẩm'
                    className='form-row'
                    name='discount'
                    rules={[
                        { required: true, message: 'Vui lòng tên sản phẩm của bạn!' },
                        {
                            validator: (_, value) => {
                                if (value < 0) {
                                    return Promise.reject('Giảm giá không thể nhỏ hơn 0!');
                                }
                                return Promise.resolve();
                            },
                        },
                    ]}
                >
                    <InputNumber type='number' className='form-input' />
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

export default AddProducts;