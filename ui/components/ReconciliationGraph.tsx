import * as React from "react";
import styled from "styled-components";
import { useGetReconciledObjects } from "../hooks/flux";
import { Condition, ObjectRef } from "../lib/api/core/types.pb";
import { UnstructuredObjectWithChildren } from "../lib/graph";
import { removeKind } from "../lib/utils";
import DirectedGraph from "./DirectedGraph";
import { ReconciledVisualizationProps } from "./ReconciledObjectsTable";
import RequestStateHandler from "./RequestStateHandler";

export type Props = ReconciledVisualizationProps & {
  parentObject: {
    name?: string;
    namespace?: string;
    conditions?: Condition[];
    suspended?: boolean;
    children?: UnstructuredObjectWithChildren[];
  };
  source: ObjectRef;
};

function ReconciliationGraph({
  className,
  parentObject,
  automationKind,
  kinds,
  clusterName,
  source,
}: Props) {
  //grab data
  const {
    data: objects,
    error,
    isLoading,
  } = parentObject
    ? useGetReconciledObjects(
        parentObject?.name,
        parentObject?.namespace,
        automationKind,
        kinds,
        clusterName
      )
    : { data: [], error: null, isLoading: false };
  //add extra nodes
  const secondNode = {
    ...parentObject,
    kind: removeKind(automationKind),
    children: objects,
  };
  const rootNode = {
    ...source,
    kind: removeKind(source.kind),
    children: [secondNode],
  };

  return (
    <RequestStateHandler loading={isLoading} error={error}>
      <DirectedGraph className={className} rootNode={rootNode} />
    </RequestStateHandler>
  );
}

export default styled(ReconciliationGraph).attrs({
  className: ReconciliationGraph.name,
})``;
