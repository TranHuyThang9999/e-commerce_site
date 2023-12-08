// dk
import React, { useState } from 'react';
import { Button, Form, Input, InputNumber, Radio, Upload, DatePicker, message } from 'antd';
import axios from 'axios';

import "./SingUp.css"

function SingUp() {
    const [form] = Form.useForm();
    const [imageFile, setImageFile] = useState(null);
    const formData = new FormData();

    const handleFormSubmit = async (values) => {
        try {
            // Upload file
            formData.append('image', imageFile);

            // Set form data
            formData.append('last_name', values.last_name);
            formData.append('first_name', values.first_name);
            formData.append('age', values.age);
            formData.append('address', values.address);
            formData.append('gender', values.gender); //oneof= 39 41 43"
            formData.append('email', values.email);
            formData.append('phone_number', values.phone_number);
            formData.append('date_of_birth', values.date_of_birth.unix());
            formData.append('user_name', values.user_name);
            formData.append('password', values.password);
            formData.append('notes', values.notes);

            // Gửi dữ liệu lên máy chủ sử dụng Axios
            const response = await axios.post('http://localhost:8080/sell/add', formData);

            // Xử lý kết quả từ máy chủ
            if (response.data.result.code === 10) {
                alert('Tài khoản đã tồn tại');
            } else if (response.data.result.code === 16) {
                alert('Tài khoản email đã tồn tại');
            } else if (response.data.result.code === 4
                || response.data.result.code === 14
                || response.data.result.code === 20
                || response.data.result.code === 12
                || response.data.result.code === 6) {
                alert('lỗi server ');
            } else if (response.data.result.code === 18) {
                alert('Số điện thoại đã tồn tại');
            } else {
                message.success("Đăng ký tài khoản thành công")
            }

        } catch (error) {
            console.error(error);
            alert('Có lỗi xảy ra. Vui lòng thử lại.');
            return;
        }
    };

    const layout = {
        labelCol: {
            span: 8,
        },
        wrapperCol: {
            span: 16,
        },
    };

    return (
        <div className='container'>
            <Form
                {...layout}
                form={form}
                onFinish={handleFormSubmit}
                label="create"
                className='form-sign-in'
                initialValues={formData}
            >
                {/* Ô input cho Họ và tên */}
                <Form.Item
                    label="Họ"
                    name="last_name"
                    rules={[{ required: true, message: 'Vui lòng nhập thông tin tên của bạn!' }]}
                >
                    <Input className='form-input' />
                </Form.Item>

                <Form.Item
                    label="Đệm và Tên"
                    name="first_name"
                    rules={[{ required: true, message: 'Vui lòng nhập thông tin tên của bạn!' }]}
                >
                    <Input className='form-input' />
                </Form.Item>

                {/* Ô input cho Tuổi */}
                <Form.Item
                    label="Tuổi"
                    name="age"
                    rules={[{ required: true, message: 'Vui lòng nhập tuổi của bạn!' }]}
                >
                    <InputNumber className='form-input' />
                </Form.Item>

                {/* Ô input cho Địa chỉ */}
                <Form.Item
                    label="Địa chỉ"
                    name="address"
                    rules={[{ required: true, message: 'Vui lòng nhập địa chỉ của bạn!' }]}
                >
                    <Input className='form-input' />
                </Form.Item>

                {/* Radio button cho Giới tính */}
                <Form.Item
                    label="Giới tính"
                    name="gender"
                    rules={[{ required: true, message: 'Vui lòng chọn giới tính của bạn!' }]}
                >
                    <Radio.Group>
                        <Radio value={39}>Nam</Radio>
                        <Radio value={41}>Nữ</Radio>
                        <Radio value={43}>Không tiết lộ</Radio>
                    </Radio.Group>
                </Form.Item>

                {/* Ô input cho Email */}
                <Form.Item
                    label="Email"
                    name="email"
                    rules={[{ required: true, message: 'Vui lòng nhập địa chỉ email của bạn!' }]}
                >
                    <Input className='form-input' type="email" />
                </Form.Item>



                {/* Ô input cho Số điện thoại */}
                <Form.Item
                    label="Số điện thoại"
                    name="phone_number"
                    rules={[{ required: true, message: 'Vui lòng nhập số điện thoại của bạn!' }]}
                >
                    <InputNumber className='form-input' />
                </Form.Item>

                <Form.Item
                    label="Ngày sinh"
                    name="date_of_birth"
                    rules={[{ required: true, message: 'Vui lòng nhập ngày sinh của bạn!' }]}
                >
                    <DatePicker className='form-input' />
                </Form.Item>

                {/* Ô input cho Tên đăng nhập */}
                <Form.Item
                    label="Tên đăng nhập"
                    name="user_name"
                    rules={[{ required: true, message: 'Vui lòng nhập tên đăng nhập của bạn!' }]}
                >
                    <Input className='form-input' />
                </Form.Item>

                {/* Ô input cho Mật khẩu */}
                <Form.Item
                    label="Mật khẩu"
                    name="password"
                    rules={[{ required: true, message: 'Vui lòng nhập mật khẩu của bạn!' }]}
                >
                    <Input.Password className='form-input' />
                </Form.Item>

                {/* Ô input cho Ghi chú */}
                <Form.Item
                    label="Ghi chú"
                    name="notes"
                >
                    <Input.TextArea className='form-input' />
                </Form.Item>



                <Form.Item
                    label="Ảnh đại diện"
                    name="image"
                    valuePropName="fileList"
                    getValueFromEvent={(e) => {
                        if (Array.isArray(e)) {
                            return e;
                        }
                        return e && e.fileList;
                    }}
                // rules={[{ required: true, message: 'Vui lòng tải lên ảnh đại diện của bạn!' }]}
                >

                    <Upload

                        maxCount={1}
                        type=''
                        listType='picture-card'
                        openFileDialogOnClick
                        accept="image/jpeg,image/png"
                        beforeUpload={(file) => {
                            setImageFile(file);
                            return false;
                        }}
                        onRemove={() => {
                            setImageFile(null);
                        }}
                    >
                        +Upload
                    </Upload>

                </Form.Item>

                {/* Nút Submit */}
                <Form.Item wrapperCol={{ ...layout.wrapperCol, offset: 8 }}>
                    <Button type="primary" htmlType="submit">
                        Đăng ký
                    </Button>
                </Form.Item>



            </Form>
        </div>
    );
}

export default SingUp;
