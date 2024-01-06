import axios from "../helpers/axios"
import _ from 'lodash'

interface SignInData {
    email: string
    password: string
}

interface RegisterData extends SignInData {
    full_name: string
}

interface DataModifyingResponse {
    success: boolean
    error: string | null
}

interface Tokens {
    access_token: string
    refresh_token: string
}

interface SignInResponse {
    tokens: Tokens
    error: string | null
}

async function register(newUser: RegisterData): Promise<DataModifyingResponse> {
    try {
        const { data } = await axios.post('/user/register', newUser)

        const success: boolean = _.get(data, ['data', 'success'], false)
        const error: string | null = _.get(data, ['data', 'error'], null)

        return { success, error }
    } catch (err) {
        const data = _.get(err, ['response', 'data'])

        const success: boolean = _.get(data, ['success'], false)
        const error: string | null = _.get(data, ['error'], null)

        return { success, error }
    }
}

async function signIn(account: SignInData): Promise<SignInResponse> {
    try {
        const { data } = await axios.post('/user/sign-in', account)

        const tokens: Tokens = _.get(data, ['data'], null)
        const error: string | null = _.get(data, ['data', 'error'], null)

        return { tokens, error }
    } catch (err) {
        const data = _.get(err, ['response', 'data'])

        const tokens: Tokens = _.get(data, ['data'], null)
        const error: string | null = _.get(data, ['error'], null)

        return { tokens, error }
    }
}

async function refreshToken(): Promise<Tokens | null> {
    try {
        const refresh_token = localStorage.getItem('refresh_token')
        if(!refresh_token) return null

        const { data } = await axios.get(`/user/refresh-token?refresh-token=${refresh_token}`)

        const tokens: Tokens = _.get(data, ['data'], null)
        return tokens
    } catch (error) {
        return null
    }
}

export {
    register,
    signIn,
    refreshToken
}