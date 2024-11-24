import React, { useState } from "react";
import { Link } from "react-router-dom";
import 'react-phone-number-input/style.css';
import PhoneInput from 'react-phone-number-input';
import { Form, Button, Alert } from "react-bootstrap";
import { useUserAuth } from "../context/UserAuthContext";

const PhoneSignUp = () => {
    const [number, setNumber] = useState("");
    const [otp, setOtp] = useState("");
    const [error, setError] = useState("");
    const { setUpRecaptcha } = useUserAuth();
    const {flag, setFlag} = useState(false);
    const {confirmObj, setConfirmObj} = useState("");

    const getOtp = async (e) => {
        e.preventDefault();
        setError("");
        if (number === "" || number === undefined) {
            return setError("Please enter a valid phone number!");
        }
        try {
            const response = await setUpRecaptcha(number);
            console.log(response);
            setConfirmObj(response);
            setFlag(true);
        } catch (err) {
            setError(err.message);
        }
    };

    const verifyOtp = (e) => {
        e.preventDefault();
        console.log(otp); // Add OTP verification logic here
    };

    return (
        <>
            <div className="p-4 box">
                <h2 className="mb-3">Firebase Authentication</h2>
                {error && <Alert variant="danger">{error}</Alert>}

                {/* Get OTP */}
                <Form onSubmit={getOtp}>
                    <Form.Group className="mb-3" controlId="formBasicPhoneNumber">
                        <PhoneInput
                            defaultCountry="IN"
                            value={number}
                            onChange={setNumber}
                            placeholder="Enter phone number"
                        />
                        <div id="recaptcha-container"/>
                    </Form.Group>
                    <div className="btn-right">
                        <Link to="/">
                            <Button variant="secondary">Cancel</Button>
                        </Link>
                        &nbsp;
                        <Button variant="primary" type="submit">Send OTP</Button>
                    </div>
                </Form>

                {/* Verify OTP */}
                <Form onSubmit={verifyOtp}>
                    <Form.Group className="mb-3" controlId="formBasicOtp">
                        <Form.Control
                            type="text"
                            placeholder="Enter OTP"
                            value={otp}
                            onChange={(e) => setOtp(e.target.value)}
                        />
                    </Form.Group>
                    <div className="btn-right">
                        <Link to="/">
                            <Button variant="secondary">Cancel</Button>
                        </Link>
                        &nbsp;
                        <Button variant="primary" type="submit">Verify OTP</Button>
                    </div>
                </Form>
            </div>
        </>
    );
};

export default PhoneSignUp;
