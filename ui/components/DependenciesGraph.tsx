import * as React from "react";
import styled from "styled-components";
import { useListObjects } from "../hooks/objects";
import { Condition, FluxObjectKind } from "../lib/api/core/types.pb";
import { fluxObjectKindToKind, FluxObjectWithChildren } from "../lib/objects";
import { removeKind } from "../lib/utils";
import DirectedGraph from "./DirectedGraph";
import RequestStateHandler from "./RequestStateHandler";

export type Props = {
  className?: string;
  currentObject: {
    name?: string;
    namespace?: string;
    conditions?: Condition[];
    suspended?: boolean;
    children?: FluxObjectWithChildren[];
  };
  automationKind: FluxObjectKind;
};

function DependenciesGraph({
  className,
  currentObject,
  automationKind,
}: Props) {
  const {
    data: objects,
    isLoading,
    error,
  } = currentObject
    ? useListObjects("", fluxObjectKindToKind(automationKind))
    : { data: [], error: null, isLoading: false };

  const rootNode = {
    ...currentObject,
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
