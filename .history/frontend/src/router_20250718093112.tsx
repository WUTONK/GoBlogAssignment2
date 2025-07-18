import LoginPage from "./loginPage";

import { createBrowserRouter } from "react-router-dom";

const router = createBrowserRouter(
   [
        {
            path: "/user/login",
            element: <App/>
        },

        {
            path: "/user/info",
            element: <infoPage/>,
        }
   ]
)