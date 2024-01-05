import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'
import axios from '../../helpers/axios'

interface User {
    id: string
    email: string
    username: string
    full_name: string
}

export interface AuthState {
  accessToken: string | null
  user: User | null
}

const initialState: AuthState = {
  accessToken: null,
  user: null
}

export const counterSlice = createSlice({
  name: 'counter',
  initialState,
  reducers: {
    storeAccessToken: (state, action: PayloadAction<string>) => { 
        state.accessToken = action.payload
        axios.defaults.headers.common['Authorization'] = state.accessToken;
    }
  },
})

// Action creators are generated for each case reducer function
export const { storeAccessToken } = counterSlice.actions

export default counterSlice.reducer