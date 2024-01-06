import { FunctionComponent } from 'react';
import { Navigate } from 'react-router-dom';
import { useAppSelector } from '../redux/hooks';

interface ProtectedRouteProps {
    children: JSX.Element
}

const ProtectedRoute: FunctionComponent<ProtectedRouteProps> = ({ children }) => {
    const { isAuthenticated } = useAppSelector(state => state.auth);

    if (!isAuthenticated) {
        // user is not authenticated
        return <Navigate to="/sign-in" />;
    }

    return children;
};

export default ProtectedRoute;