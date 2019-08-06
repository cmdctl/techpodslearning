import React from 'react';
import {Provider} from "react-redux";
import {MuiThemeProvider} from "@material-ui/core";
import Routes from "./routes";
import {store} from './redux'
import theme from './theme';



function App() {
  return (
    <Provider store={store}>
        <MuiThemeProvider theme={theme}>
            <Routes/>
        </MuiThemeProvider>
    </Provider>
  );
}

export default App;
