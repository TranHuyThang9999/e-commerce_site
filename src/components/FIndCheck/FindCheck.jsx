import { Button, Checkbox, Form, Input, Radio, Select } from 'antd';
import React, { useState } from 'react';
const Option = Select.Option;
const { TextArea } = Input;


function FindCheck() {

    // checkbox 1 = tao , 2= cam , 3 = chanh
    // radionButton 1 = nu , 2 = nam
    // selection 1 = Anh , 2 = Mi , 3 = Duc

    // Khởi tạo state để lưu trạng thái form

    localStorage.setItem("data", fo)

    const hanlderFormSubmit = () => {
        var formData = new FormData();
        formData.append('typeFruit', typeFruit);
        formData.append('gender', gender);
        formData.append('typeCountry', typeCountry);
        formData.append('describe', describe);

    }

    return (
        <div>
            <Form>

                <Form.Item name='typeFruit'>
                    <Checkbox>Táo</Checkbox>
                    <Checkbox>Cam</Checkbox>
                    <Checkbox>Chanh</Checkbox>
                </Form.Item>

                <Form.Item name='gender'>
                    <Radio.Group>
                        <Radio value={1}>Nữ</Radio>
                        <Radio value={2}>Nam</Radio>
                    </Radio.Group>
                </Form.Item>

                <Form.Item name='typeCountry'>
                    <Select defaultValue="USA">
                        <Option value="Anh">Anh</Option>
                        <Option value="USA">USA</Option>
                        <Option value="France">France</Option>
                    </Select>
                </Form.Item>

                <Form.Item name='describe'>
                    <TextArea rows={4} />
                    <br />

                </Form.Item>

                <Form.Item>
                    <Button type="primary" htmlType='submit' >
                        Save
                    </Button>
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType='submit'>
                        Open
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
}

export default FindCheck;