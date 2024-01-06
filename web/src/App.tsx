import { Outlet } from "react-router-dom";
import AuthenticationWrapper from "./components/AuthenticationWrapper";

function App() {
    return (  
        <AuthenticationWrapper>
            <Outlet />
        </AuthenticationWrapper>
    );
}

export default App;