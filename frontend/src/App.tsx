// import React from 'react';
// import logo from './logo.svg';
// import { useMetaMask } from "metamask-react";
// import NewTransactionPage from './pages/NewTransaction';

import { createRoot } from "react-dom/client";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import NewTransactionPage from "./pages/NewTransaction";
import PrintingPage from "./pages/PrintingPage";
import ConfirmationPage from "./pages/ConfirmationPage";

function App() {
  const basename = document.querySelector("base")?.getAttribute("href") ?? "/";
  return (
    <div className="flex flex-col relative bg-base-200 min-h-screen">
      <BrowserRouter basename={basename}>
        <Routes>
          <Route path="/" element={<NewTransactionPage />} />
          <Route path="printing" element={<PrintingPage />} />
          <Route path="confirmation" element={<ConfirmationPage />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
