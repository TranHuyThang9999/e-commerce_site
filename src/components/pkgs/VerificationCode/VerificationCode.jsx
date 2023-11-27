import { useState } from "react";
import OTPInput, { ResendOTP } from "otp-input-react";
import { Button, Col, Form, Row, Space } from "antd";
import './VerificationCode.css';

export default function VerificationCode() {

    const [OTP, setOTP] = useState("");
    
    console.log(OTP)
    return (

        <div>
            <Form
                className="VerificationCode"
            >
                <Form.Item>
                    <OTPInput value={OTP} onChange={setOTP} autoFocus OTPLength={6} otpType="number" disabled={false} />
                </Form.Item>
                <Form.Item>
                    <Row  justify="space-between">
                        <Col span={8}offset={8}>
                            <Button>Submit</Button>
                        </Col>
                        <Col span={8}>
                            <Button>Back</Button>
                        </Col>
                    </Row>
                </Form.Item>
            </Form>

            {/* <ResendOTP onResendClick={() => console.log("Resend clicked")} /> */}
        </div>

    );
}
