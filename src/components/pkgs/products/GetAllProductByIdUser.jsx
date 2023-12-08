import React, { useState, useEffect } from 'react';
import { Table } from 'antd';
import axios from 'axios';

const GetAllProductByIdUser = () => {
    const [products, setProducts] = useState([]);
 
    useEffect(() => {
        const getListProduct = async () => {
            try {
                const response = await axios.get('http://localhost:8080/sell/user/product/list', {
                    params:{
                        user_name: 'thangth1',
                    },
                    headers: {
                        Authorization: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDE1MjU3NTMsImlkIjo0MzU1Njg0NzQ0MDQ4NjQsInVzZXJfbmFtZSI6InRoYW5ndGgxIiwicm9sZSI6MTcwMTQzOTM1M30.k2bGaGumvzZN7-pEYxQLUiU6Ql3Q2PYug5fa8ZKih1Q',
                    },
                });
                console.log(response.data.product);
                setProducts(response.data.product);
            } catch (error) {
                console.error('Error while fetching flight information:', error);
            }
        };

        getListProduct();
    }, []);

    // "id": 435544398530048,
    // "id_user": 435523826097408,
    // "name_product": "Iphone 13",
    // "quantity": 11,
    // "sell_status": 11,
    // "price": 110000000,
    // "discount": 33,
    // "manufacturer": "Microsoft 11",
    // "created_at": 1701345306,
    // "updated_at": 1701345306,
    // "describe": "Good global 11",
    // "id_type_product": 33,
    // "number_of_photos": 2

    const columns = [
        {
            title: 'ID',
            dataIndex: 'id',
            key: 'id',
        },
        {
            title: 'Tên sản phẩm',
            dataIndex: 'name_product',
            key: 'name_product',
        },
        {
            title: 'Số lượng',
            dataIndex: 'quantity',
            key: 'quantity',
        },
        {
            title: 'Giá',
            dataIndex: 'price',
            key: 'price',
        },
        {
            title: 'So luong anh',
            dataIndex: 'number_of_photos',
            key: 'number_of_photos',
        },
        
    ];

    return (
        <div>
            <Table columns={columns} dataSource={products} />
        </div>
    );
};

export default GetAllProductByIdUser;
