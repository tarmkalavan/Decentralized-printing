import React, { useState } from "react";
import styled from "styled-components";
import PrinterTable from "../components/PrinterTable";
import Tabbar from "../components/Tabbar";
import FileInput from "../components/FileInput";
import { FaCheck, FaPrint } from "react-icons/fa";
import { Interface } from "ethers/lib/utils";
import Abi from "../assets/Abi";
// import { ethers } from "ethers";

interface INewTransactionPageProps {}

export enum webState {
    NEWTRANSACTION = "newTransaction",
    PRITING = "printing",
    CONFRIMATION = "confirmation",
}

const printerInterface = new Interface(Abi.printer);

const NewTransactionPage: React.FunctionComponent<INewTransactionPageProps> = (
    props
) => {
    const submitNewTransaction = (url: string, lenPage: number) => {
        console.log(url, lenPage);
    };

    const [state, setState] = useState(webState.NEWTRANSACTION);

    const [selectedNum, setSelectedNum] = useState(-1);
    let printerList = [
        {
            id: "f2cc5a27-6638-4287-8bfe-e872390777be",
            location: "OVG",
            name: "Romeo Zulu Quebec Papa Golf Whiskey Yankee X-ray Kilo Victor Sierra Juliett",
            price: 2.085,
        },
        {
            id: "3daa46ce-b5e5-4830-8d10-de42d899e38f",
            location: "EAR",
            name: "Juliett Tango Echo Oscar Kilo Quebec Mike Victor India Golf X-ray Papa",
            price: 4.5587,
        },
        {
            id: "823e56d8-4126-4955-adc4-fad1134ff8ea",
            location: "APU",
            name: "Romeo Whiskey Quebec Sierra Bravo Echo Victor Zulu Juliett Uniform Foxtrot Lima Kilo Charlie Oscar Golf Mike Alfa November",
            price: 1.7229,
        },
        {
            id: "1fb87ed7-6726-422b-b21d-e72bce3e99ce",
            location: "OBO",
            name: "Charlie Bravo Zulu Golf Sierra Romeo Oscar Foxtrot Whiskey Hotel Victor Uniform Yankee Alfa Papa X-ray",
            price: 1.1851,
        },
    ];
    let fileInput;

    if (selectedNum !== -1) {
        fileInput = (
            <FileInput
                price={printerList[selectedNum].price}
                submitNewTransaction={submitNewTransaction}
                setState={setState}
            />
        );
    } else {
        fileInput = <div></div>;
    }

    if (state === webState.PRITING) {
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
    } else if (state === webState.CONFRIMATION) {
        return (
            <div className="flex flex-col h-screen">
                <Tabbar />
                <div className="flex flex-col h-screen justify-center items-center space-y-8 p-4">
                    <FaCheck size={150} color="green" />
                    <h1 className="text-[32px] font-mina font-medium">
                        {" "}
                        Successfully print your file. Please check that
                        everything is OK and press the button.
                    </h1>
                    <div className="flex flex-row space-x-4">
                        <RejectButton> Reject </RejectButton>
                        <ConfirmButton>Confirm</ConfirmButton>
                    </div>
                </div>
            </div>
        );
    }
    return (
        <Container>
            <Tabbar />
            <Content>
                <Label>
                    <LabelText>Order new work</LabelText>
                </Label>

                <TableContainer>
                    <LabelSmallText>
                        Select printer you want to use
                    </LabelSmallText>
                    <PrinterTable
                        onSelectedPrinter={(index) => {
                            setSelectedNum(index);
                        }}
                        printers={printerList}
                    />
                </TableContainer>
                {fileInput}
            </Content>
        </Container>
    );
};

export default NewTransactionPage;

// Styled-component

const Container = styled.div`
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    padding: 0px;
`;

const Content = styled.div`
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    padding: 64px 96px;
    gap: 48px;

    flex: none;
    align-self: stretch;
    flex-grow: 0;
`;

const Label = styled.div`
    /* Auto layout */

    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 0px;
    gap: 16px;

    flex: none;
    align-self: stretch;
    flex-grow: 0;
`;

const LabelText = styled.p`
    /* Bold/36px */

    font-family: "Mina";
    font-style: normal;
    font-weight: 700;
    font-size: 36px;
    line-height: 57px;

    color: #000000;

    /* Inside auto layout */

    flex: none;
    flex-grow: 0;
`;

const TableContainer = styled.div`
    /* Auto layout */

    display: flex;
    flex-direction: column;
    align-items: flex-start;
    padding: 0px;
    gap: 0.5px;

    flex: none;
    align-self: stretch;
    flex-grow: 0;
`;

const LabelSmallText = styled.p`
    /* Bold/24px */

    font-family: "Mina";
    font-style: normal;
    font-weight: 700;
    font-size: 24px;
    line-height: 38px;

    color: #000000;
`;

const ConfirmButton = styled.button`
    /* Button */

    /* Auto layout */

    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    padding: 8px 16px;
    gap: 8px;

    /* blue/400 */

    background: #4280cb;
    border-radius: 8px;

    /* Regular/20px */

    font-family: "Mina";
    font-style: normal;
    font-weight: 400;
    font-size: 20px;
    line-height: 32px;
    /* identical to box height */

    /* gray/50 */

    color: #f1f1f1;

    /* Inside auto layout */

    flex: none;
    flex-grow: 0;
`;

const RejectButton = styled.button`
    /* Button */

    /* Auto layout */

    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    padding: 8px 16px;
    gap: 8px;

    /* blue/400 */

    background: #d63e33;
    border-radius: 8px;

    /* Regular/20px */

    font-family: "Mina";
    font-style: normal;
    font-weight: 400;
    font-size: 20px;
    line-height: 32px;
    /* identical to box height */

    /* gray/50 */

    color: #f1f1f1;

    /* Inside auto layout */

    flex: none;
    flex-grow: 0;
`;
