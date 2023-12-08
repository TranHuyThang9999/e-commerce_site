import React from 'react';
import { Button, Form, Input, Checkbox, message } from 'antd';
import axios from 'axios';
import { LockOutlined, UserOutlined,LoginOutlined } from '@ant-design/icons';
import '../user/user.css';
function SignIn() {
    const [form] = Form.useForm();

    const errorMessage = () => {
        message.error('Lỗi hệ thống vui lòng thử lại');
    };

    const sucessMessage = ()=>{
        message.success("Đăng nhập thành công")
    }
    
    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();
            formData.append('user_name', values.user_name);
            formData.append('password', values.password);

            const response = await axios.post('http://localhost:9090/sell/user/login', formData);

            if (response.data.result.code === 0) {
                if (form.getFieldValue('remember')) {
                    localStorage.setItem('token', response.data.jwt_token.access_token);
                }
                sucessMessage();
                localStorage.setItem('user_name', values.user_name);
            } else {
                alert('Thông tin tài khoản hoặc mật khẩu không chính xác. Vui lòng thử lại.');
            }
        } catch (error) {
            console.error(error);
            // alert('Có lỗi xảy ra. Vui lòng thử lại.');
            errorMessage();
        }
    }

    const tailLayout = {
        wrapperCol: {
            offset: 0,
        },
    };

    return (
        <div className="container-login-user">
            <div className='form-login'>
                <Form
                    {...tailLayout}
                    className='login-form'
                    form={form}
                    onFinish={handleFormSubmit}
                    initialValues={{
                        remember: true,
                    }}
                >
                    <Form.Item
                        labelAlign='right'
                        className='form-login-user-label-header form-row'
                    >
                        <h2>Classy Login Form</h2>
                    </Form.Item>

                    <Form.Item
                        className='form-row'
                        name='user_name'
                        rules={[{ required: true, message: 'Vui lòng nhập tài khoản của bạn!' }]}
                    >
                        <Input
                            className='form-login-input'
                            prefix={<UserOutlined className="site-form-item-icon" />}
                            placeholder="Username"
                        />
                    </Form.Item>

                    <Form.Item
                        className='form-row'
                        name='password'
                        rules={[{ required: true, message: 'Vui lòng nhập mật khẩu của bạn!' }]}
                    >
                        <Input.Password
                            className='form-login-input'
                            prefix={<LockOutlined className="site-form-item-icon" />}
                            placeholder="Password"
                        />
                    </Form.Item>

                    <Form.Item>
                        <Form.Item name="remember" valuePropName="checked" noStyle>
                            <Checkbox>Remember me</Checkbox>
                        </Form.Item>

                        <a className="login-form-forgot" href="/">
                            Forgot password
                        </a>
                    </Form.Item>

                    <Form.Item>
                        <Button style={{ fontSize: '20px' }} type="primary" htmlType='submit'>
                            Sign in
                            <LoginOutlined />
                        </Button>
                    </Form.Item>

                    <Form.Item>
                        <h3>ok</h3>
                    </Form.Item>
                </Form>
            </div>
        </div>
    );
}

export default SignIn;