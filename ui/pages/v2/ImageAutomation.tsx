import * as React from "react";
import { useQuery } from "react-query";
import styled from "styled-components";
import Page from "../../components/Page";
import { useImageAutomation } from "../../contexts/ImageAutomationContext";

type Props = {
  className?: string;
};

function ImageAutomation({ className }: Props) {
  const { api } = useImageAutomation();
  const { data, error } = useQuery(["imageupdateautomations"], () =>
    api.ListImageAutomations({})
  );

  console.log(data);
  return (
    <Page title="ImageAutomation" error={null} className={className}>
      ImageAutomation
    </Page>
  );
}

export default styled(ImageAutomation).attrs({
  className: ImageAutomation.name,
})``;
