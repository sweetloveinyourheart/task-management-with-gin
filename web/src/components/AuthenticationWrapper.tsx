import { useEffect, useState } from "react";
import { refreshToken } from "../services/user.service";
import { storeAccessToken } from "../redux/slices/authSlice";
import { useAppDispatch } from "../redux/hooks";

function AuthenticationWrapper({ children }: any) {
    const dispatch = useAppDispatch()

    useEffect(() => {
        initAuthentication()
    }, [])

    const initAuthentication = async () => {
        const tokens = await refreshToken()
        if(tokens) {
            dispatch(storeAccessToken(tokens.access_token))
        }
    }

    return (
        <>
            {children}
        </>
    );
}

export default AuthenticationWrapper;