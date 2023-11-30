package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	qrcode "github.com/skip2/go-qrcode"
)

func generateQR(w http.ResponseWriter, r *http.Request) {
	// Điền thông tin thanh toán của bạn ở đây (ví dụ: số tài khoản ngân hàng)
	paymentInfo := "YourBankAccountNumber"
	qrCode := fmt.Sprintf("YourPaymentURL?info=%s", paymentInfo)

	// Tạo mã QR
	err := qrcode.WriteFile(qrCode, qrcode.Medium, 256, "qr.png")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Hiển thị mã QR trong trang HTML
	tmpl, err := template.New("qr").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>QR Code Payment</title>
	</head>
	<body>
		<h1>Scan the QR Code to make a payment</h1>
		<img src="qr.png" alt="QR Code">
	</body>
	</html>
	`)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", generateQR)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
