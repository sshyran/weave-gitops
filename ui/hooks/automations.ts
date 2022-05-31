import { useContext } from "react";
import { useMutation } from "react-query";
import { CoreClientContext } from "../contexts/CoreClientContext";
import {
  SyncAutomationRequest,
  SyncAutomationResponse,
} from "../lib/api/core/core.pb";
import { RequestError, Syncable } from "../lib/types";

export function useSyncAutomation(obj: Syncable) {
  const { api } = useContext(CoreClientContext);
  const mutation = useMutation<
    SyncAutomationResponse,
    RequestError,
    SyncAutomationRequest
  >(({ withSource }) => api.SyncAutomation({ ...obj, withSource }));

  return mutation;
}
