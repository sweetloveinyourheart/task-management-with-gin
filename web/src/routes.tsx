import { Route, createBrowserRouter, createRoutesFromElements } from "react-router-dom";
import DashboardPage from "./pages/Dashboard";
import SignInPage from "./pages/SignIn";
import SignUpPage from "./pages/SignUp";
import App from "./App";
import ProtectedRoute from "./components/ProtectedRoute";
import HomePage from "./pages/Home";

const routers = createBrowserRouter(
    createRoutesFromElements(
        <Route path="/" element={<App />}>
            <Route index element={<HomePage />} />
            <Route path="dashboard" element={<ProtectedRoute children={<DashboardPage />} />} />
            <Route path="sign-in" element={<SignInPage />} />
            <Route path="sign-up" element={<SignUpPage />} />
        </Route>
    )
);

export default routers