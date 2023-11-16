import React, { useState, useEffect } from 'react';
import { Button, Radio } from 'antd';

function CheckRadion() {
    const [value, setValue] = useState();

    useEffect(() => {
        // Load trạng thái đã chọn từ localStorage khi component được render
        const storedValue = localStorage.getItem('selectedRadioValue');
        if (storedValue !== null) {
            // Chuyển đổi giá trị từ chuỗi sang số và cập nhật trạng thái
            setValue(Number(storedValue));
        }
    }, []);

    const onChange = (e) => {
        console.log('radio checked', e.target.value);
        setValue(e.target.value);
    };

    const handleSubmit = () => {
        // Lưu trạng thái đã chọn vào localStorage khi ấn nút Submit
        localStorage.setItem('selectedRadioValue', String(value));
        alert('State saved to localStorage!');
    };

    return (
        <div>
            <Radio.Group onChange={onChange} value={value}>
                <Radio value={1}>A</Radio>
                <Radio value={2}>B</Radio>
                <Radio value={3}>C</Radio>
                <Radio value={4}>D</Radio>
            </Radio.Group>
            <Button onClick={handleSubmit}>Submit</Button>
        </div>
    );
}

export default CheckRadion;
