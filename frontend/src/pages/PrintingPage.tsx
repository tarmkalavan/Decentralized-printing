import React from "react";
import { FaPrint } from "react-icons/fa";
import Tabbar from "../components/Tabbar";

function PrintingPage() {
    return (
        <div className="h-screen">
            <Tabbar />
            <div className="flex flex-col justify-center h-screen items-center space-y-8 p-4">
                <FaPrint size={150} color="#4280cb" />
                <h1 className="text-[36px] font-medium font-mina">
                    {" "}
                    Your file is now printing. Please wait...{" "}
                </h1>
            </div>
        </div>
    );
}

export default PrintingPage;
