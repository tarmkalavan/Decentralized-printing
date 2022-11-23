import React, { useEffect, useState } from "react";
import styled from "styled-components";
// import styled from "styled-components";

interface IPrinterTableProps {
    printers: {
        name: string;
        id: string;
        location: string;
        price: number;
    }[];
    onSelectedPrinter: (i: number) => void;
}

const PrinterTable: React.FunctionComponent<IPrinterTableProps> = (props) => {
    const [selectedNum, setSelectedNum] = useState(-1);

    useEffect(() => {
        props.onSelectedPrinter(selectedNum);
    });

    return (
        <Container>
            {props.printers.map((printer, index) => {
                return (
                    <Row
                        key={index + 6}
                        onClick={() => {
                            setSelectedNum(index);
                        }}
                        rowSelected={index === selectedNum}
                    >
                        <div className="flex flex-col">
                            <h2 className="text-[18px]">{printer.name}</h2>
                            <h3 className="text-[14px]">ID: {printer.id}</h3>
                            <h3 className="text-[14px]">
                                Location: {printer.location}
                            </h3>
                        </div>
                        <div className="flex flex-row self-center space-x-2">
                            <h2 className="font-bold text-[18px]">
                                {printer.price} Wei
                            </h2>
                            <h2>per page</h2>
                        </div>
                    </Row>
                );
            })}
        </Container>
    );
};

export default PrinterTable;

const Row = styled.div<{
    rowSelected: boolean;
}>`
    /* List Item */

    /* Auto layout */

    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 8px 16px;
    gap: 16px;

    /* gray/50 */

    background: ${(props) => (props.rowSelected ? `#C4C4C4;` : `#F1F1F1;`)}

    /* Inside auto layout */

    flex: none;
    order: 0;
    align-self: stretch;
    flex-grow: 0;

    &:hover {
        background-color: #C4C4C4;
    }
`;

const Container = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 0px;

    /* Inside auto layout */

    height: 240px;
    overflow: scroll;

    flex: none;
    align-self: stretch;
    flex-grow: 0;
`;
