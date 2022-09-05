import * as React from "react";
import styled from "styled-components";
import { Automation } from "../hooks/automations";
import useNavigation from "../hooks/navigation";
import { getSourceRefForAutomation } from "../lib/utils";
import Button from "./Button";
import Flex from "./Flex";
import Icon, { IconType } from "./Icon";
import DependenciesGraph from "./DependenciesGraph";

function Message() {
  const { navigate } = useNavigation();

  return (
    <Flex column align center shadow>
      <h1>No Dependencies</h1>

      <p>
        There are no dependencies set up for your kustomizations or helm
        releases at this time. You can set them up using the dependsOn field on
        the kustomization or helm release object.
      </p>

      <h1>What are depedencies for?</h1>

      <p>
        Dependencies allow you to relate different kustomizations and helm
        releases, as well as specifying an order in which your resources should
        be started. For example, you can wait for a database to report as
        'Ready' before attempting to deploy other services.
      </p>

      <Button
        startIcon={<Icon type={IconType.AddIcon} size="base" />}
        onClick={() =>
          navigate.external(
            "https://fluxcd.io/flux/components/kustomize/kustomization/#kustomization-dependencies"
          )
        }
      >
        Learn More
      </Button>
    </Flex>
  );
}

type Props = {
  automation?: Automation;
};

function DependenciesView({ automation }: Props) {
  const hasDependencies = automation?.dependsOn?.length > 0;

  return hasDependencies ? (
    <DependenciesGraph
      automationKind={automation?.kind}
      automationName={automation?.name}
      kinds={automation?.inventory}
      parentObject={automation}
      clusterName={automation?.clusterName}
      source={getSourceRefForAutomation(automation)}
    />
  ) : (
    <Message />
  );
}

export default styled(DependenciesView).attrs({
  className: DependenciesView.name,
})`
  color: ${(props) => props.theme.colors.neutral30};
`;
