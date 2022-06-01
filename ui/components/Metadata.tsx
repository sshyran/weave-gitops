import * as React from "react";
import Text from "./Text";
import Flex from "./Flex";
import styled from "styled-components";
import Spacer from "./Spacer";
import InfoList from "./InfoList";

type Props = {
  className?: string;
  metadata: any;
}

function Metadata({ metadata, className }) {
  if (!metadata.length) {
    return (<></>);
  }

  return (
    <Flex wide column className={className}>
      <Text size="large">Metadata</Text>
      <InfoList items={metadata} />
    </Flex>
  );
}

export default styled(Metadata).attrs({
  className: Metadata.name,
})`
  margin-top: ${(props) => props.theme.spacing.medium};
`;
