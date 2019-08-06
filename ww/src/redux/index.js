import {auth} from "./auth";
import {errors} from "./error"
import {configureStore, getDefaultMiddleware} from "redux-starter-kit";

const reducer = {
    auth,
    errors
};


export const store = configureStore({
    reducer,
    devTools: process.env.NODE_ENV !== 'production',
    middleware: [...getDefaultMiddleware()]
})