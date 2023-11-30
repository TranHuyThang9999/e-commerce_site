

thiết kế database bảng user với 3 vai trò --  người dùng ban đầu đk vào dùng app là người mua , sau đó có thể là đk bán , người quản trị có thể khóa trạng thái mua hoặc bán của người dùng

user có --
--------id, FullName ,Age , Address , Email (xác thực mã OTP gửi về), nếu là người dung thì thêm phần địa chỉ nhân hàng , UserName , password 
-------nếu là người bán có thêm gian hàng bản thân 
------- người mua có thể xem mặt hàng , add vào giỏ , mua (thanh toán online trước hoặc sau) , hoàn trả hàng , đánh giá sản phẩm 

// thieeus check logicn xem accoun mo chua
//api update otp
// lam producr all , car them xua xoa
add date_of_birth -ok
--Home

//  có thể thêm 1 bảng nũa check thời gian sống của OTP

// phần update người thêm vào có thể xóa hoặc thêm ảnh từ bảng  có phần bảng ảnh mới xửa xóa


//list product







































///	now := time.Now()
	otherTime := now.Add(60 * time.Second)
	nowTimestamp := now.Unix()
	otherTimestamp := otherTime.Unix()
	if otherTimestamp-nowTimestamp <= 60 {
		fmt.Println("Hai thời điểm cách nhau chính xác 60 giây.", "*")
	} else {
		fmt.Println("Hai thời điểm không cách nhau đúng 60 giây.", "#")
	}
//