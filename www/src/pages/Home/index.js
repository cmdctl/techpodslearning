import HomePage from "./Home";
import {connect} from "react-redux";
import {setCurrentUser} from "../../redux/auth";
import {addError} from "../../redux/error";

const mstp = (state) => {
    return {}
};

const mdtp = dispatch => {
    return {
        load: async () => {
            try {
                const res = await fetch("/api/me")
                const data = await res.json();
                dispatch(setCurrentUser(data))
            } catch (err) {
                dispatch(addError(err))
            }
        }
    }
};
export default connect(mstp, mdtp)(HomePage)