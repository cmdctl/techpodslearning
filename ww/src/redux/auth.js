import {createAction, createReducer} from "redux-starter-kit";

export const setCurrentUser = createAction('SET_CURRENT_USER');

export const auth = createReducer({} , {
    [setCurrentUser.type]: (state, action) => {
        return action.payload;
    }
});