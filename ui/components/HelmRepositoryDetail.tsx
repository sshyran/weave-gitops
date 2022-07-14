import * as React from "react";
import styled from "styled-components";
import { removeKind } from "../lib/utils";
import { FluxObjectKind } from "../lib/api/core/types.pb";
import { HelmRepository } from "../lib/objects";
import Interval from "./Interval";
import Link from "./Link";
import SourceDetail from "./SourceDetail";
import Timestamp from "./Timestamp";

type Props = {
  className?: string;
  name: string;
  namespace: string;
  clusterName: string;
};

function HelmRepositoryDetail({
  name,
  namespace,
  className,
  clusterName,
}: Props) {
  return (
    <SourceDetail
      className={className}
      name={name}
      namespace={namespace}
      clusterName={clusterName}
      type={FluxObjectKind.KindHelmRepository}
      info={(hr: HelmRepository) => [
        ["Type", removeKind(FluxObjectKind.KindHelmRepository)],
        ["Repository Type", hr.repositoryType.toLowerCase()],
        ["URL", tryLink(hr.url)],
        [
          "Last Updated",
          hr.lastUpdatedAt ? <Timestamp time={hr.lastUpdatedAt} /> : "-",
        ],
        ["Interval", <Interval interval={hr.interval} />],
        ["Cluster", hr.clusterName],
        ["Namespace", hr.namespace],
      ]}
    />
  );
}

export function tryLink(url: string): React.ReactElement<any, any> | string {
  if (url.startsWith("http")) {
    return (
      <Link newTab href={url}>
        {url}
      </Link>
    );
  } else {
    return url;
  }
}

export default styled(HelmRepositoryDetail).attrs({
  className: HelmRepositoryDetail.name,
})``;
