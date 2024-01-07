import { FunctionComponent } from "react";
import GlassEffect from "./GlassEffect";
import { Col, Row } from "antd";

interface WorkspaceBarProps {

}

const WorkspaceBar: FunctionComponent<WorkspaceBarProps> = () => {
    return (
        <GlassEffect>
            <Row gutter={[24, 24]}>
                <Col span={6}></Col>
            </Row>
        </GlassEffect>
    );
}

export default WorkspaceBar;