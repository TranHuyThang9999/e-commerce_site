import React, { useState } from 'react';
import axios from 'axios';
import { Button, Col, Form, Input, Row } from 'antd';
import CheckBox from './CheckBox';

const OPTIONS = ["tre em", "nguoi lon", "nguoi gia"];

function FindCheck() {
    const [infoProduct, setProduct] = useState([]);
    const [userId, setUserId] = useState('');
    const [checkboxes, setCheckboxes] = useState(
        OPTIONS.reduce((options, option) => ({ ...options, [option]: false }), {})
    );

    const fetchData = async (id) => {
        try {
            const response = await axios.get(`http://localhost:8080/sell/proudcts?id=${id}`);
            setProduct(response.data.products);
        } catch (error) {
            console.error("Error fetching data:", error);
        }
    };

    const handleButtonClick = () => {
        fetchData(userId);
    };



    const renderProductDetails = () => {
        return infoProduct.map((product, index) => (
            <ul key={index}>
                <li>{product.name}</li>
                <li>{product.quantity}</li>
                <li>{product.price}</li>
                <li>{product.sell_status}</li>
                <li>{product.object_of_use}</li>
                <li>{product.product_type}</li>
            </ul>
        ));
    };

    return (
        <div>
            <Form>
                <Row>
                    <Col span={8}>
                        <Input
                            type="text"
                            onChange={(e) => setUserId(e.target.value)}
                        />
                    </Col>
                    <Col span={8}>
                        <Button onClick={handleButtonClick}>Fetch Data</Button>
                    </Col>
                </Row>
                <Row>
                    <Col>{renderProductDetails()}</Col>
                </Row>
            </Form>
        </div>
    );
}

export default FindCheck;
