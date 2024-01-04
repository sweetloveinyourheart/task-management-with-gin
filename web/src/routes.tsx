import { createBrowserRouter } from "react-router-dom";
import HomePage from "./pages/Home";
import SignInPage from "./pages/SignIn";
import SignUpPage from "./pages/SignUp";

const routers = createBrowserRouter([
    {
        path: "/",
        element: <HomePage />,
    },
    {
        path: "/sign-in",
        element: <SignInPage />
    },
    {
        path: "/sign-up",
        element: <SignUpPage />
    }
]);

export default routers