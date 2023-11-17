-- Active: 1696642727227@@127.0.0.1@5432@sell_product@public

CREATE TABLE accounts (
    id BIGINT PRIMARY KEY,
    id_role BIGINT,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    age INT,
    address VARCHAR(255),
    gender INT,
    email VARCHAR(255) UNIQUE,
    phone_number VARCHAR(255),
    user_name VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    otp_code BIGINT, -- Đặt trường OTP làm kiểu VARCHAR với chiều dài 10 (độ dài mã OTP có thể thay đổi tùy thuộc vào yêu cầu của bạn)
    otp_expiry int, -- Nếu bạn muốn theo dõi thời gian hết hạn của OTP 
    is_verified INT,
    store_name VARCHAR(255) UNIQUE,
    Notes VARCHAR(255),
    created_at  INT,
	updated_at INT,
    avatar VARCHAR(255)
);
COMMENT ON COLUMN accounts.id_role IS 'ID của vai trò';


CREATE Table roles(
    id BIGINT PRIMARY KEY,            -- ID chính làm khóa chính
    admin INT,
    seller int,--
    buyer int--
);


CREATE TABLE products(
    id BIGINT PRIMARY KEY,            -- ID chính làm khóa chính
    id_user BIGINT,
    name_product VARCHAR(255),
    quantity INT, -- so luong
    sell_status INT,
    price NUMERIC(10, 2),  -- gia
    discount NUMERIC(5, 2),           -- Giảm giá (nếu có)
    manufacturer VARCHAR(255),        -- Nhà sản xuất
    created_at INT, -- Ngày tạo sản phẩm
    updated_at INT, -- Ngày cập nhật sản phẩm
    describe VARCHAR(510),
    id_type_product BIGINT,
    list_id_image  VARCHAR(255)
);

CREATE TABLE image_storages(
    id BIGINT PRIMARY KEY,
    url VARCHAR(255),
    id_user BIGINT,
    id_product BIGINT
)

CREATE TABLE orders(
    id BIGINT PRIMARY KEY,
    id_product BIGINT,
    id_buyer BIGINT,
    id_seller BIGINT,
    order_status INT,
    payment_status INT,
    created_at INT,                   -- Ngày tạo
    updated_at INT                    -- Ngày cập nhật
);


CREATE TABLE shipping_address (
    id BIGINT PRIMARY KEY,            -- ID chính làm khóa chính
    id_user BIGINT,                  -- ID của người dùng (user)
    full_name VARCHAR(255),  -- Họ và tên
    phone_number VARCHAR(20),-- Số điện thoại
    province VARCHAR(255),   -- Tỉnh/Thành phố
    district VARCHAR(255),   -- Quận/Huyện
    commune VARCHAR(255) ,    -- Xã/Phường
    village VARCHAR(255),    -- Làng/Đường
    street_name VARCHAR(255),-- Tên đường
    Notes VARCHAR(255),               -- Ghi chú
    created_at INT,                   -- Ngày tạo
    updated_at INT                    -- Ngày cập nhật
);


CREATE TABLE carts(
    id BIGINT PRIMARY KEY,
    id_product BIGINT,
    id_seller BIGINT,
    created_at INT,                   -- Ngày tạo
    updated_at INT                    -- Ngày cập nhật
);

CREATE TABLE product_type( ---máy tính , điện thoại , quần áo , mĩ phẩm , bàn phím , ram laptop , ....
    id BIGINT PRIMARY KEY,
);


SELECT *FROM accounts;
DELETE FROM accounts;

SELECT *FROM roles;
DELETE FROM roles;
SELECT * FROM "accounts" ORDER BY "accounts"."id" LIMIT 1

SELECT *FROM products;
SELECT *FROM image_storages;

DELETE FROM image_storages;
DELETE FROM products;
