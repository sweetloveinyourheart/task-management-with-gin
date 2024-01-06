// GlassEffect.tsx
import { FunctionComponent } from 'react';

interface GlassEffectProps {
    children: JSX.Element
    rounded?: boolean
    styles?: React.CSSProperties
}

const GlassEffect: FunctionComponent<GlassEffectProps> = ({ children, rounded, styles }) => {
    return (
        <div className="glass-container" style={{
            ...styles,
            borderRadius: rounded ? 12 : 0
        }}>
            {children}
        </div>
    );
};

export default GlassEffect;
