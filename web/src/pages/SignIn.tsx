import { Button, Divider, Form, Input, Typography } from "antd";
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { useNavigate } from "react-router-dom";

function SignInPage() {
    const navigate = useNavigate();

    const onFinish = (values: any) => {
        console.log('Received values of form: ', values);
    };

    const switchToRegister = () => navigate("/sign-up")

    return (
        <div className="center-container">
            <div style={{ border: "1px solid #dcdcdc", padding: 24, borderRadius: 12 }}>
                <Typography.Title style={{ marginBottom: 16, marginTop: 0 }}>Login to Easy Task</Typography.Title>
                <Divider />
                <Form
                    name="normal_login"
                    className="login-form"
                    initialValues={{ remember: true }}
                    onFinish={onFinish}
                >
                    <Form.Item
                        name="username"
                        rules={[{ required: true, message: 'Please input your Username!' }]}
                    >
                        <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Username" />
                    </Form.Item>
                    <Form.Item
                        name="password"
                        rules={[{ required: true, message: 'Please input your Password!' }]}
                    >
                        <Input
                            prefix={<LockOutlined className="site-form-item-icon" />}
                            type="password"
                            placeholder="Password"
                        />
                    </Form.Item>
                    <Divider />
                    <Form.Item>
                        <Button block type="primary" htmlType="submit" className="login-form-button">
                            Log in
                        </Button>
                    </Form.Item>
                    <Divider>Or</Divider>
                    <Form.Item>
                        <Button block type="link" htmlType="button" className="login-form-button" onClick={switchToRegister}>
                            Register new account
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        </div>
    );
}

export default SignInPage;