import React from "react";

import Tabbar from "../components/Tabbar";
import { FaCheck } from "react-icons/fa";
import styled from "styled-components";

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

function ConfirmationPage() {
  return (
    <div className="flex flex-col h-screen">
      <Tabbar />
      <div className="flex flex-col h-screen justify-center items-center space-y-8 p-4">
        <FaCheck size={150} color="green" />
        <h1 className="text-[32px] font-mina font-medium">
          {" "}
          Successfully print your file. Please check that everything is OK and
          press the button.
        </h1>
        <div className="flex flex-row space-x-4">
          <RejectButton> Reject </RejectButton>
          <ConfirmButton>Confirm</ConfirmButton>
        </div>
      </div>
    </div>
  );
}

export default ConfirmationPage;
