import { useEffect, useState } from "react";
import { refreshToken } from "../services/user.service";
import { setAuthenticatedStatus } from "../redux/slices/authSlice";
import { useAppDispatch } from "../redux/hooks";
import PageLoading from "./Loading";

function AuthenticationWrapper({ children }: any) {
    const [isLoading, setIsLoading] = useState<boolean>(true)

    const dispatch = useAppDispatch()

    useEffect(() => {
        initAuthentication()
    }, [])

    const initAuthentication = async () => {
        setIsLoading(true)
        const tokens = await refreshToken()
        if(tokens) {
            dispatch(setAuthenticatedStatus(true))
        } else {
            dispatch(setAuthenticatedStatus(false))
            localStorage.removeItem('refresh_token')
        }
        setIsLoading(false)
    }

    return (
        <>
            {isLoading ? <PageLoading /> : children}
        </>
    );
}

export default AuthenticationWrapper;