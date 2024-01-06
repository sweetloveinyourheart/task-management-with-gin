import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

export interface User {
  id: string
  email: string
  username: string
  full_name: string
}

export interface AuthState {
  isAuthenticated: boolean
  user: User | null
}

const initialState: AuthState = {
  isAuthenticated: false,
  user: null
}

export const counterSlice = createSlice({
  name: 'counter',
  initialState,
  reducers: {
    setAuthenticatedStatus: (state, action: PayloadAction<boolean>) => {
      state.isAuthenticated = action.payload;
    },
    setUser: (state, action: PayloadAction<User>) => {
      state.user = action.payload;
    },
    logout: (state) => {
      state.isAuthenticated = false;
      state.user = null;
      localStorage.removeItem('refresh_token')
    }
  },
})

// Action creators are generated for each case reducer function
export const { setAuthenticatedStatus, setUser, logout } = counterSlice.actions

export default counterSlice.reducer