import * as React from "react";
import styled from "styled-components";
import { HelmRelease } from "../hooks/objects";
import Alert from "./Alert";
import AutomationDetail from "./AutomationDetail";
import SourceLink from "./SourceLink";
import Timestamp from "./Timestamp";

type Props = {
  name: string;
  clusterName: string;
  helmRelease?: HelmRelease;
  className?: string;
};

function HelmReleaseDetail({ helmRelease, className }: Props) {
  return (
    <AutomationDetail
      className={className}
      automation={helmRelease}
      info={
        [
          ["Source", <SourceLink sourceRef={helmRelease.sourceRef()} clusterName={helmRelease.clusterName()} />],
          ["Chart", <SourceLink sourceRef={helmRelease.helmChartRef()} clusterName={helmRelease.clusterName()} />],
          ["Cluster", helmRelease?.clusterName()],
          ["Interval", helmRelease?.interval()],
          [
            "Last Updated",
            <Timestamp time={helmRelease.lastUpdated()} />,
          ],
        ]
      }
    />
  );
}

export default styled(HelmReleaseDetail).attrs({
  className: HelmReleaseDetail.name,
})`
  width: 100%;

  ${Alert} {
    margin-bottom: 16px;
  }
`;
