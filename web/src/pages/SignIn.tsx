import { Alert, Button, Divider, Form, Input, Space, Typography } from "antd";
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { useNavigate } from "react-router-dom";
import _ from 'lodash'
import { useEffect, useState } from "react";

import * as userService from "../services/user.service";
import { storeAccessToken } from "../redux/slices/authSlice";
import { useAppDispatch, useAppSelector } from "../redux/hooks";

function SignInPage() {
    const [error, setError] = useState<string | null>(null)

    const { accessToken } = useAppSelector(state => state.auth)
    const dispatch = useAppDispatch()
    const navigate = useNavigate();

    useEffect(() => {
        if(accessToken) navigate('/dashboard')
    }, [accessToken])

    const onFinish = async (values: any) => {
        // Handle form submission
        const account: any = _.omit(values, ['confirm'])
        const res = await userService.signIn(account)

        if (res.error) {
            setError(res.error)
            return;
        }

        const accessToken = res.tokens.access_token
        dispatch(storeAccessToken(accessToken))

        const refreshToken = res.tokens.refresh_token
        localStorage.setItem('refresh_token', refreshToken)

        navigate("/")
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
                        <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Email" />
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

                    {error
                        ? (
                            <Space direction="vertical" style={{ width: '100%', marginBottom: 24 }}>
                                <Alert message={error} type="error" />
                            </Space>
                        )
                        : null
                    }

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