import { Button, Divider, Form, Input, Typography } from "antd";
import { LockOutlined, UserOutlined, MailOutlined } from '@ant-design/icons';
import { useNavigate } from "react-router-dom";

function SignUpPage() {
    const navigate = useNavigate();

    const onFinish = (values: any) => {
        // Handle form submission
        console.log('Received values:', values);
    };

    const switchToLogin = () => navigate("/sign-in")

    return (
        <div className="center-container">
            <div style={{ border: "1px solid #dcdcdc", padding: 24, borderRadius: 12 }}>
                <Typography.Title style={{ marginBottom: 16, marginTop: 0 }}>Register new account</Typography.Title>
                <Divider />
                <Form
                    name="register"
                    onFinish={onFinish}
                    scrollToFirstError
                >
                    <Form.Item
                        name="email"
                        rules={[
                            {
                                type: 'email',
                                message: 'The input is not a valid email address!',
                            },
                            {
                                required: true,
                                message: 'Please input your email!',
                            },
                        ]}
                    >
                        <Input prefix={<MailOutlined />} placeholder="Email*" />
                    </Form.Item>

                    <Form.Item
                        name="password"
                        rules={[
                            {
                                required: true,
                                message: 'Please input your password!',
                            },
                        ]}
                        hasFeedback
                    >
                        <Input.Password prefix={<LockOutlined />} placeholder="Password*" />
                    </Form.Item>

                    <Form.Item
                        name="confirm"
                        dependencies={['password']}
                        hasFeedback
                        rules={[
                            {
                                required: true,
                                message: 'Please confirm your password!',
                            },
                            ({ getFieldValue }) => ({
                                validator(_, value) {
                                    if (!value || getFieldValue('password') === value) {
                                        return Promise.resolve();
                                    }
                                    return Promise.reject(new Error('The two passwords do not match!'));
                                },
                            }),
                        ]}
                    >
                        <Input.Password prefix={<LockOutlined />} placeholder="Confirm Password*" />
                    </Form.Item>

                    <Form.Item
                        name="full_name"
                    >
                        <Input prefix={<UserOutlined />} placeholder="Your full name" />
                    </Form.Item>

                    <Form.Item>
                        <Button block type="primary" htmlType="submit">
                            Register
                        </Button>
                    </Form.Item>

                    <Divider>Or</Divider>
                    <Form.Item>
                        <Button block type="link" htmlType="button" className="login-form-button" onClick={switchToLogin}>
                            Login to your account
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        </div>
    );
}

export default SignUpPage;