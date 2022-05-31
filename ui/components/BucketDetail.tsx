import * as React from "react";
import styled from "styled-components";
import SourceDetail from "../components/SourceDetail";
import Timestamp from "../components/Timestamp";
import { Kind, Bucket } from "../hooks/objects";

type Props = {
  className?: string;
  name: string;
  namespace: string;
  clusterName: string;
};

function BucketDetail({ name, namespace, className, clusterName }: Props) {
  return (
    <SourceDetail
      className={className}
      name={name}
      clusterName={clusterName}
      namespace={namespace}
      type={Kind.Bucket}
      info={(b: Bucket) => [
        ["Type", b.kind()],
        ["Endpoint", b.endpoint()],
        ["Bucket Name", b.name()],
        ["Last Updated", <Timestamp time={b.lastUpdated()} />],
        ["Interval", b.interval()],
        ["Cluster", b.clusterName()],
        ["Namespace", b.namespace()],
      ]}
    />
  );
}

export default styled(BucketDetail).attrs({ className: BucketDetail.name })``;
