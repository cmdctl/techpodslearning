import { createMuiTheme } from '@material-ui/core/styles';
import {blue, pink, red} from "@material-ui/core/colors";
import {responsiveFontSizes} from "@material-ui/core";

const theme = createMuiTheme({
    palette: {
        primary: blue,
        secondary: pink,
    },
    status: {
        danger: red,
    },
});

export default responsiveFontSizes(theme)