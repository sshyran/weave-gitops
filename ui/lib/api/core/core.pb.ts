/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as Gitops_coreV1Types from "./types.pb"
export type Pagination = {
  pageSize?: number
  pageToken?: string
}

export type ListError = {
  clusterName?: string
  namespace?: string
  message?: string
}

export type ListFluxRuntimeObjectsRequest = {
  namespace?: string
  clusterName?: string
}

export type ListFluxRuntimeObjectsResponse = {
  deployments?: Gitops_coreV1Types.Deployment[]
  errors?: ListError[]
}

export type ListObjectsRequest = {
  namespace?: string
  kind?: string
  pagination?: Pagination
}

export type ListObjectsResponse = {
  objects?: Gitops_coreV1Types.Object[]
  nextPageToken?: string
  errors?: ListError[]
}

export type GetObjectRequest = {
  name?: string
  namespace?: string
  kind?: string
  clusterName?: string
}

export type GetObjectResponse = {
  object?: Gitops_coreV1Types.Object
}

export type GetReconciledObjectsRequest = {
  automationName?: string
  namespace?: string
  automationKind?: string
  kinds?: Gitops_coreV1Types.GroupVersionKind[]
  clusterName?: string
}

export type GetReconciledObjectsResponse = {
  objects?: Gitops_coreV1Types.UnstructuredObject[]
}

export type GetChildObjectsRequest = {
  groupVersionKind?: Gitops_coreV1Types.GroupVersionKind
  namespace?: string
  parentUid?: string
  clusterName?: string
}

export type GetChildObjectsResponse = {
  objects?: Gitops_coreV1Types.UnstructuredObject[]
}

export type GetFluxNamespaceRequest = {
}

export type GetFluxNamespaceResponse = {
  name?: string
}

export type ListNamespacesRequest = {
}

export type ListNamespacesResponse = {
  namespaces?: Gitops_coreV1Types.Namespace[]
}

export type ListFluxEventsRequest = {
  involvedObject?: Gitops_coreV1Types.ObjectRef
}

export type ListFluxEventsResponse = {
  events?: Gitops_coreV1Types.Event[]
}

export type SyncAutomationRequest = {
  name?: string
  namespace?: string
  kind?: string
  clusterName?: string
  withSource?: boolean
}

export type SyncAutomationResponse = {
}

export type GetVersionRequest = {
}

export type GetVersionResponse = {
  semver?: string
  commit?: string
  branch?: string
  buildTime?: string
}

export type GetFeatureFlagsRequest = {
}

export type GetFeatureFlagsResponse = {
  flags?: {[key: string]: string}
}

export type ToggleSuspendResourceRequest = {
  kind?: string
  name?: string
  namespace?: string
  clusterName?: string
  suspend?: boolean
}

export type ToggleSuspendResourceResponse = {
}

export class Core {
  static ListObjects(req: ListObjectsRequest, initReq?: fm.InitReq): Promise<ListObjectsResponse> {
    return fm.fetchReq<ListObjectsRequest, ListObjectsResponse>(`/v1/object?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetObject(req: GetObjectRequest, initReq?: fm.InitReq): Promise<GetObjectResponse> {
    return fm.fetchReq<GetObjectRequest, GetObjectResponse>(`/v1/object/${req["name"]}?${fm.renderURLSearchParams(req, ["name"])}`, {...initReq, method: "GET"})
  }
  static ListFluxRuntimeObjects(req: ListFluxRuntimeObjectsRequest, initReq?: fm.InitReq): Promise<ListFluxRuntimeObjectsResponse> {
    return fm.fetchReq<ListFluxRuntimeObjectsRequest, ListFluxRuntimeObjectsResponse>(`/v1/flux_runtime_objects?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetReconciledObjects(req: GetReconciledObjectsRequest, initReq?: fm.InitReq): Promise<GetReconciledObjectsResponse> {
    return fm.fetchReq<GetReconciledObjectsRequest, GetReconciledObjectsResponse>(`/v1/reconciled_objects`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetChildObjects(req: GetChildObjectsRequest, initReq?: fm.InitReq): Promise<GetChildObjectsResponse> {
    return fm.fetchReq<GetChildObjectsRequest, GetChildObjectsResponse>(`/v1/child_objects`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetFluxNamespace(req: GetFluxNamespaceRequest, initReq?: fm.InitReq): Promise<GetFluxNamespaceResponse> {
    return fm.fetchReq<GetFluxNamespaceRequest, GetFluxNamespaceResponse>(`/v1/namespace/flux`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ListNamespaces(req: ListNamespacesRequest, initReq?: fm.InitReq): Promise<ListNamespacesResponse> {
    return fm.fetchReq<ListNamespacesRequest, ListNamespacesResponse>(`/v1/namespaces?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ListFluxEvents(req: ListFluxEventsRequest, initReq?: fm.InitReq): Promise<ListFluxEventsResponse> {
    return fm.fetchReq<ListFluxEventsRequest, ListFluxEventsResponse>(`/v1/events?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static SyncAutomation(req: SyncAutomationRequest, initReq?: fm.InitReq): Promise<SyncAutomationResponse> {
    return fm.fetchReq<SyncAutomationRequest, SyncAutomationResponse>(`/v1/sync`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetVersion(req: GetVersionRequest, initReq?: fm.InitReq): Promise<GetVersionResponse> {
    return fm.fetchReq<GetVersionRequest, GetVersionResponse>(`/v1/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetFeatureFlags(req: GetFeatureFlagsRequest, initReq?: fm.InitReq): Promise<GetFeatureFlagsResponse> {
    return fm.fetchReq<GetFeatureFlagsRequest, GetFeatureFlagsResponse>(`/v1/featureflags?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ToggleSuspendResource(req: ToggleSuspendResourceRequest, initReq?: fm.InitReq): Promise<ToggleSuspendResourceResponse> {
    return fm.fetchReq<ToggleSuspendResourceRequest, ToggleSuspendResourceResponse>(`/v1/suspend`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}