import {createAction, createReducer} from "redux-starter-kit";

export const addError = createAction('ADD_ERROR');
export const clearErrors = createAction('CLEAR_ERRORS');
export const removeError = createAction('REMOVE_ERRORS');

export const errors = createReducer([], {
    [addError.type]: (state, action) => {
        state.push(action.payload);
    },
    [clearErrors.type]: (state) => {
        return [];
    },
    [removeError.type]: (state, action) => {
        return state.filter(e => e.message !== action.payload.message);
    }
});