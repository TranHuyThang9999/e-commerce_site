thiết kế database bảng user với 3 vai trò --  người dùng ban đầu đk vào dùng app là người mua , sau đó có thể là đk bán , người quản trị có thể khóa trạng thái mua hoặc bán của người dùng

user có --
--------id, FullName ,Age , Address , Email (xác thực mã OTP gửi về), nếu là người dung thì thêm phần địa chỉ nhân hàng , UserName , password 
-------nếu là người bán có thêm gian hàng bản thân 
------- người mua có thể xem mặt hàng , add vào giỏ , mua (thanh toán online trước hoặc sau) , hoàn trả hàng , đánh giá sản phẩm 

// thieeus check logicn xem accoun mo chua
//api update otp
// lam producr all , car them xua xoa