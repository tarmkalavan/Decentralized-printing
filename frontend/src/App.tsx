// import React from 'react';
// import logo from './logo.svg';
// import { useMetaMask } from "metamask-react";
// import NewTransactionPage from './pages/NewTransaction';

import { ethers } from "ethers";
import { useEffect, useState } from "react";
import Abi from "./assets/Abi";
import NewTransactionPage from "./pages/NewTransaction";

function App() {
    const [isLoadContract, setLoadContract] = useState(false);
    const basename =
        document.querySelector("base")?.getAttribute("href") ?? "/";
    useEffect(() => {
        async function test() {
            const provider = new ethers.providers.Web3Provider(window.ethereum);
            await provider.send("eth_requestAccounts", []);

            const testContractAddress = `0x0DdC6059d7Ea36D9DD36E569b1e8E0a1E40D3fBd`;
            const printerContract = new ethers.Contract(
                testContractAddress,
                Abi.printer,
                provider
            );
            const data = await printerContract.printerData;
            console.log(data);
            setLoadContract(true);
        }

        if (!isLoadContract) {
            test();
        }
    });
    return <NewTransactionPage />;
}

export default App;
