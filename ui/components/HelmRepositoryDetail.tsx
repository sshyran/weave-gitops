import * as React from "react";
import styled from "styled-components";
import Link from "../components/Link";
import SourceDetail from "../components/SourceDetail";
import Timestamp from "../components/Timestamp";
import { HelmRepository, Kind } from "../hooks/objects";

type Props = {
  className?: string;
  name: string;
  namespace: string;
  clusterName: string;
};

function HelmRepositoryDetail({ name, namespace, className, clusterName }: Props) {
  return (
    <SourceDetail
      className={className}
      name={name}
      namespace={namespace}
      clusterName={clusterName}
      type={Kind.HelmRepository}
      info={(hr: HelmRepository) => [
        ["Type", hr.kind()],
        [
          "URL",
          <Link newTab href={hr.url()}>
            {hr.url()}
          </Link>,
        ],
        ["Last Updated", <Timestamp time={hr.lastUpdated()} />],
        ["Interval", hr.interval()],
        ["Cluster", hr.clusterName()],
        ["Namespace", hr.namespace()],
      ]}
    />
  );
}

export default styled(HelmRepositoryDetail).attrs({
  className: HelmRepositoryDetail.name,
})``;
