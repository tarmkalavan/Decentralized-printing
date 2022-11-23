import * as React from "react";
import styled from "styled-components";

interface ITabbarProps {}

const Tabbar: React.FunctionComponent<ITabbarProps> = (props) => {
    return (
        <MenuBar>
            <Text>Decentralized Aggregator Printing Services</Text>
            <Text></Text>
        </MenuBar>
    );
};

export default Tabbar;

const MenuBar = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: flex-end;
    padding: 8px 16px;

    background: #4280cb;

    flex: none;
    align-self: stretch;
    flex-grow: 0;
`;
const Text = styled.p`
    font-family: "Mina";
    font-style: normal;
    font-weight: 400;
    font-size: 16px;

    /* blue/100 */

    color: #c8dbf1;

    /* Inside auto layout */

    flex: none;
    order: 0;
    flex-grow: 0;
`;
