import * as React from "react";
import styled from "styled-components";
import { useListObjects } from "../hooks/objects";
import { Condition, ObjectRef } from "../lib/api/core/types.pb";
import { UnstructuredObjectWithChildren } from "../lib/graph";
import { fluxObjectKindToKind } from "../lib/objects";
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

function DependenciesGraph({ className, parentObject, automationKind }: Props) {
  const {
    data: objects,
    isLoading,
    error,
  } = parentObject
    ? useListObjects(
        parentObject?.namespace,
        fluxObjectKindToKind(automationKind)
      )
    : { data: [], error: null, isLoading: false };

  const rootNode = {
    ...parentObject,
    kind: removeKind(automationKind),
    children: objects,
  };

  return (
    <RequestStateHandler loading={isLoading} error={error}>
      <DirectedGraph className={className} rootNode={rootNode} />
    </RequestStateHandler>
  );
}

export default styled(DependenciesGraph)`
  .MuiSlider-vertical {
    min-height: 400px;
  }
  .MuiSlider-vertical .MuiSlider-track {
    width: 6px;
  }
  .MuiSlider-vertical .MuiSlider-rail {
    width: 6px;
  }
  .MuiSlider-vertical .MuiSlider-thumb {
    margin-left: -9px;
  }
  .MuiSlider-thumb {
    width: 24px;
    height: 24px;
  }
`;
