import React, { useState } from "react";
import styled from "styled-components";
import PrinterTable from "../components/PrinterTable";
import Tabbar from "../components/Tabbar";
import FileInput from "../components/FileInput";
import { Interface } from "ethers/lib/utils";
import Abi from "../assets/Abi";
// import { ethers } from "ethers";

interface INewTransactionPageProps {}

const printerInterface = new Interface(Abi.printer);

const NewTransactionPage: React.FunctionComponent<INewTransactionPageProps> = (
    props
) => {
    const submitNewTransaction = (url: string, lenPage: number) => {
        console.log(url, lenPage);
    };

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
            />
        );
    } else {
        fileInput = <div></div>;
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
