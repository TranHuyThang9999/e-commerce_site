import { Button, Select } from 'antd';
import React, { useState, useEffect } from 'react';


const options = [
    {
        value: '1',
        label: 'Jack',
    },
    {
        value: '2',
        label: 'Lucy',
    },
    {
        value: '3',
        label: 'yiminghe',
    },
];

function Selected() {

    const [selectedValue, setSelectedValue] = useState(null);

    useEffect(() => {
        const storedValue = localStorage.getItem('selectedValue');
        if (storedValue !== null) {
            // Chuyển đổi giá trị từ chuỗi sang số
            setSelectedValue(Number(storedValue));
        }
    }, []);

    const handleSelectChange = (value) => {
        setSelectedValue(value);
    };

    const handleSubmit = () => {
        localStorage.setItem('selectedValue', selectedValue);
        alert('State saved to localStorage!');
    };


    const labelEd = options.find(option => option.value === String(selectedValue))?.label;


    return (
        <div>
            <Select
                style={{
                    width: 120,
                }}
                options={options}
                value={labelEd}
                onChange={handleSelectChange}
            />
            <Button onClick={handleSubmit}>Submit</Button>
        </div>
    );
}
//
export default Selected;
