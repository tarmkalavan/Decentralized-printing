// import React from 'react';
// import logo from './logo.svg';
// import { useMetaMask } from "metamask-react";
// import NewTransactionPage from './pages/NewTransaction';

import NewTransactionPage from "./pages/NewTransaction";

function App() {
    const basename =
        document.querySelector("base")?.getAttribute("href") ?? "/";
    return <NewTransactionPage />;
}

export default App;
