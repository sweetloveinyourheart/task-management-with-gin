import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from "react-router-dom";
import { store } from './redux/store'
import { Provider } from 'react-redux'

import './index.css'
import routers from './routes';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Provider store={store}>
      <RouterProvider router={routers} />
    </Provider>
  </React.StrictMode>,
)
