import React, { useEffect, useRef, useState } from "react";
import styled from "styled-components";
import PrinterTable from "../components/PrinterTable";
import Tabbar from "../components/Tabbar";
import FileInput from "../components/FileInput";
import { FaCheck, FaPrint } from "react-icons/fa";
import Abi from "../assets/Abi";
import { ethers } from "ethers";
import Bytecode from "../assets/Bytecode";
import { sign } from "crypto";
// import { ethers } from "ethers";

interface INewTransactionPageProps {}

export enum webState {
    NEWTRANSACTION = "newTransaction",
    PRITING = "printing",
    CONFRIMATION = "confirmation",
}

export interface Printer {
    id: string;
    location: string;
    name: string;
    price: number;
}

const NewTransactionPage: React.FunctionComponent<INewTransactionPageProps> = (
    props
) => {
    const [state, setState] = useState(webState.NEWTRANSACTION);
    const [isContractLoaded, setIsContracLoaded] = useState(false);
    const [isSubmitSubmission, setSubmitSubmission] = useState(false);
    const printersList = useRef<Printer[]>([]);
    const url = useRef("");
    const lenPage = useRef(0);

    function createNewTransaction(_url: string, _lenPage: number) {
        url.current = _url;
        lenPage.current = _lenPage;
        setSubmitSubmission(true);
    }

    useEffect(() => {
        async function submitNewTransaction(_url: string, _lenPage: number) {
            const provider = new ethers.providers.Web3Provider(window.ethereum);
            await provider.send("eth_requestAccounts", []);
            const signer = provider.getSigner();
            console.log(await signer.getBalance());

            const factory = new ethers.ContractFactory(
                Abi.transaction,
                Bytecode.transaction,
                signer
            );
            console.log(printersList.current[selectedNum].price * _lenPage);

            const contract = await factory.deploy(
                _url,
                printersList.current[selectedNum].id,
                _lenPage,
                { value: printersList.current[selectedNum].price * _lenPage }
            );

            const printerContract = new ethers.Contract(
                printersList.current[selectedNum].id,
                Abi.printer,
                signer
            );

            console.log(await contract.getOwner(), await signer.getAddress());
            await printerContract.addToQueue(contract.address);
            setState(webState.PRITING);
        }

        async function loadContract() {
            const provider = new ethers.providers.Web3Provider(window.ethereum);
            await provider.send("eth_requestAccounts", []);

            const contractAddress =
                "0xf27a721E7C970978AA5ea6655ef1e3FCC1e43fd6";
            const centralServerContract = new ethers.Contract(
                contractAddress,
                Abi.centralServer,
                provider
            );
            const addresses = await centralServerContract.getPrinters();
            console.log(addresses);
            let tempPrinterData: Printer[] = [];
            addresses.forEach(async (address: string) => {
                const printerContract = new ethers.Contract(
                    address,
                    Abi.printer,
                    provider
                );
                const data = await printerContract.printerData();
                tempPrinterData.push({
                    id: address,
                    location: data.location,
                    name: data.printerName,
                    price: parseInt(data.price._hex, 16),
                });
            });
            printersList.current = tempPrinterData;
            setIsContracLoaded(true);
        }

        if (!isContractLoaded) {
            loadContract();
        }

        if (isSubmitSubmission && state === webState.NEWTRANSACTION) {
            submitNewTransaction(url.current, lenPage.current);
        }
    });

    const [selectedNum, setSelectedNum] = useState(-1);
    let fileInput;

    if (selectedNum !== -1) {
        fileInput = (
            <FileInput
                price={printersList.current[selectedNum].price}
                submitNewTransaction={createNewTransaction}
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
                        <RejectButton>Reject</RejectButton>
                        <ConfirmButton>Confirm</ConfirmButton>
                    </div>
                </div>
            </div>
        );
    }
    let tableContent;
    if (!isContractLoaded) {
        tableContent = <div></div>;
    } else {
        tableContent = (
            <PrinterTable
                onSelectedPrinter={(index) => {
                    setSelectedNum(index);
                }}
                printers={printersList.current}
            />
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
                    {tableContent}
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
