import * as React from "react";
import styled from "styled-components";
import { Kustomization } from "../hooks/objects";
import Alert from "./Alert";
import AutomationDetail from "./AutomationDetail";
import SourceLink from "./SourceLink";
import Timestamp from "./Timestamp";

type Props = {
  kustomization?: Kustomization;
  className?: string;
};

function KustomizationDetail({ kustomization, className }: Props) {
  return (
    <AutomationDetail
      className={className}
      automation={kustomization}
      info={[
        ["Source", <SourceLink sourceRef={kustomization?.sourceRef()} clusterName={kustomization?.clusterName()} />],
        ["Applied Revision", kustomization?.lastAppliedRevision()],
        ["Cluster", kustomization?.clusterName()],
        ["Path", kustomization?.path()],

        ["Interval", kustomization?.interval()],
        [
          "Last Updated",
          <Timestamp time={kustomization.lastUpdated()} />,
        ],
      ]}
    />
  );
}

export default styled(KustomizationDetail).attrs({
  className: KustomizationDetail.name,
})`
  width: 100%;

  ${Alert} {
    margin-bottom: 16px;
  }
`;
