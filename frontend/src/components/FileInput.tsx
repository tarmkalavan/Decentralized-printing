import React, { useEffect, useRef, useState } from "react";
import { PDFDocument } from "pdf-lib";
import styled from "styled-components";
import { useNavigate } from "react-router-dom";

interface IFileInputProps {
  price: number;
  submitNewTransaction: (url: string, lenPage: number) => void;
}

const FileInput: React.FunctionComponent<IFileInputProps> = (props) => {
  const [stateNum, setStateNum] = useState(0);
  const pdfUrl = useRef("");
  const lenPage = useRef(0);
  const navigate = useNavigate();

  useEffect(() => {
    async function checkFile(url: string) {
      try {
        const existingPdfBytes = await fetch(url).then((res) =>
          res.arrayBuffer()
        );
        const pdfDoc = await PDFDocument.load(existingPdfBytes);
        const pages = pdfDoc.getPages();
        lenPage.current = pages.length;
        setStateNum(2);
      } catch (error) {
        setStateNum(3);
      }
    }

    if (stateNum === 1) {
      checkFile(pdfUrl.current);
    }
  });

  if (stateNum === 0) {
    return (
      <Content>
        <Container>
          <Title>Enter your work's link</Title>
          <Input
            type="text"
            placeholder="Enter link here"
            onChange={(e) => {
              pdfUrl.current = e.currentTarget.value;
            }}
          />
        </Container>
        <Button
          onClick={() => {
            setStateNum(1);
          }}
        >
          Check file
        </Button>
      </Content>
    );
  } else if (stateNum === 1) {
    return (
      <Content>
        <Container>
          <Title>Enter your work's link</Title>
          <Input
            type="text"
            placeholder="Enter link here"
            onChange={(e) => {
              pdfUrl.current = e.currentTarget.value;
            }}
          />
        </Container>
      </Content>
    );
  } else if (stateNum === 2) {
    return (
      <Content>
        <Container>
          <Title>Enter your work's link</Title>
          <Input
            type="text"
            placeholder="Enter link here"
            onChange={(e) => {
              pdfUrl.current = e.currentTarget.value;
              setStateNum(0);
            }}
          />
        </Container>
        <RightContainer>
          <RightContext>
            <LabelVerySmallText>Total page</LabelVerySmallText>
            <LabelVerySmallText> {lenPage.current} pages </LabelVerySmallText>
          </RightContext>
          <RightContext>
            <LabelVerySmallText>Total price</LabelVerySmallText>
            <LabelVerySmallText>
              {" "}
              {lenPage.current}x{props.price} = {lenPage.current * props.price}{" "}
              ETH{" "}
            </LabelVerySmallText>
          </RightContext>
        </RightContainer>
        <Button
          onClick={() => {
            props.submitNewTransaction(pdfUrl.current, lenPage.current);
            navigate("/printing");
          }}
        >
          Finalized transaction
        </Button>
      </Content>
    );
  } else {
    return (
      <Content>
        <Container>
          <Title>Please enter REAL URL!!!!!!</Title>
          <Input
            type="text"
            placeholder="Enter link here"
            onChange={(e) => {
              pdfUrl.current = e.currentTarget.value;
            }}
          />
        </Container>
        <Button
          onClick={() => {
            setStateNum(1);
          }}
        >
          Check file
        </Button>
      </Content>
    );
  }
};

export default FileInput;

const RightContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  padding: 0px;
  gap: 16px;

  /* Inside auto layout */

  flex: none;
  flex-grow: 0;
`;

const RightContext = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  padding: 0px;
  gap: 141px;

  /* Inside auto layout */

  flex: none;
  align-self: stretch;
  flex-grow: 0;
`;

const LabelVerySmallText = styled.p`
  /* Bold/24px */

  font-family: "Mina";
  font-style: normal;
  font-weight: 700;
  font-size: 16px;
  line-height: 38px;

  color: #000000;
`;

const Container = styled.div`
  /* Frame 1 */

  /* Auto layout */

  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px;
  gap: 16px;

  /* Inside auto layout */

  flex: none;
  align-self: stretch;
  flex-grow: 0;
`;

const Button = styled.button`
  /* Button */

  /* Auto layout */

  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  padding: 8px 16px;
  gap: 8px;

  /* blue/400 */

  background: #649ad7;
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

const Input = styled.input`
    box-sizing: border-box;

    /* Auto layout */
    
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    padding: 8px 16px;
    gap: 10px;
    
    width: 100%;
    
    background: #FFFFFF;
    /* gray/800 */
    
    border: 1px solid #6C6C6C;
    border-radius: 4px;
    
    /* Inside auto layout */
    
    flex: none;
    align-self: stretch;
    flex-grow: 0;
    }
    `;

const Title = styled.p`
  /* Bold/24px */

  font-family: "Mina";
  font-style: normal;
  font-weight: 700;
  font-size: 24px;
  line-height: 38px;

  color: #000000;
`;

const Content = styled.div`
  /* Frame 5 */

  /* Auto layout */

  display: flex;
  flex-direction: column;
  align-items: flex-end;
  padding: 0px;
  gap: 16px;

  /* Inside auto layout */

  flex: none;
  align-self: stretch;
  flex-grow: 0;
`;
