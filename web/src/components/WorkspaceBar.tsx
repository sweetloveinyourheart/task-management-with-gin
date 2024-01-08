import { FunctionComponent } from "react";
import GlassEffect from "./GlassEffect";
import { Button, Col, Flex, Row, Select, Typography } from "antd";
import { PlusCircleOutlined } from "@ant-design/icons";

interface WorkspaceBarProps {

}

const WorkspaceBar: FunctionComponent<WorkspaceBarProps> = () => {
    return (
        <GlassEffect styles={{ marginBlock: 12 }}>
            <Row gutter={[24, 24]}>
                <Col span={6}></Col>
                <Col span={6}></Col>
                <Col span={6}></Col>
            </Row>
        </GlassEffect>
    );
}

export default WorkspaceBar;