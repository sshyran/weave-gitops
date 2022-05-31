import * as React from "react";
import styled from "styled-components";
import SourceDetail from "../components/SourceDetail";
import Timestamp from "../components/Timestamp";
import { Kind, HelmChart } from "../hooks/objects";

type Props = {
  className?: string;
  name: string;
  namespace: string;
  clusterName: string;
};

function HelmChartDetail({ name, namespace, className, clusterName }: Props) {
  return (
    <SourceDetail
      name={name}
      namespace={namespace}
      type={Kind.HelmChart}
      clusterName={clusterName}
      className={className}
      info={(ch: HelmChart) => [
        ["Type", ch?.kind()],
        ["Chart", ch?.chartName()],
        ["Ref", ch?.sourceRef().name],
        ["Last Updated", <Timestamp time={ch?.lastUpdated()} />],
        ["Interval", ch?.interval()],
        ["Cluster", ch?.clusterName()],
        ["Namespace", ch?.namespace()],
      ]
      }
    />
  );
}

export default styled(HelmChartDetail).attrs({
  className: HelmChartDetail.name,
})``;
