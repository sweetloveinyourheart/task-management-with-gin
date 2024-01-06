// Header.tsx
import { Layout, Avatar, Row, Col, Flex, Typography, Dropdown, Button } from 'antd';
import { DingdingOutlined, LogoutOutlined } from '@ant-design/icons';
import GlassEffect from './GlassEffect';
import { useEffect, useMemo } from 'react';
import { useAppDispatch, useAppSelector } from '../redux/hooks';
import { getUserProfile } from '../services/user.service';
import { logout, setUser } from '../redux/slices/authSlice';
import _ from 'lodash';

const { Header } = Layout;

const UserMenu = () => {

    const { user } = useAppSelector(state => state.auth)
    const dispatch = useAppDispatch()

    useEffect(() => {
        fetchUserProfile()
    }, [])

    const fetchUserProfile = async () => {
        const user = await getUserProfile()
        if (user) {
            dispatch(setUser(user))
        }
    }

    const items = useMemo(() => [
        {
            label: (
                <Button
                    type='link'
                    icon={<LogoutOutlined />}
                    onClick={() => dispatch(logout())}
                >
                    Logout
                </Button>
            ),
            key: '0',
        }
    ], []);

    return (
        <Flex justify='flex-end' align='center'>
            <div style={{ paddingInline: 24 }}>
                <Dropdown menu={{ items }}>
                    <Avatar style={{ background: "#f93707" }} size="default">
                        {_.get(user, 'full_name', 'user')[0]}
                    </Avatar>
                </Dropdown>
            </div>
        </Flex>
    )
}

const AppHeader = () => {
    return (
        <Header style={{ background: "transparent", marginTop: 12 }}>
            <GlassEffect rounded>
                <Row gutter={[24, 24]}>
                    <Col span={6}>
                        <div style={{ paddingInline: 24, display: 'flex', alignItems: 'center', height: '100%' }}>
                            <DingdingOutlined style={{ color: "#fff", fontSize: 32 }} />
                            &nbsp;
                            <Typography.Text strong style={{ color: "#fff", fontSize: 20 }} >
                                <span>TinyTask</span>
                            </Typography.Text>
                        </div>
                    </Col>
                    <Col span={12}></Col>
                    <Col span={6}>
                        <UserMenu />
                    </Col>
                </Row>
            </GlassEffect>
        </Header>
    );
};

export default AppHeader;
