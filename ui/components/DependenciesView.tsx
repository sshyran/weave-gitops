import * as React from "react";
import styled from "styled-components";
import { Automation } from "../hooks/automations";

type Props = {
  className?: string;
  automation?: Automation;
};

function DependenciesView({ className, automation }: Props) {
  const hasDependencies = automation?.dependsOn?.length > 0;

  return hasDependencies ? <>DependenciesGraph</> : <>Message</>;
}

export default styled(DependenciesView).attrs({
  className: DependenciesView.name,
})`
  color: ${(props) => props.theme.colors.neutral30};
`;
