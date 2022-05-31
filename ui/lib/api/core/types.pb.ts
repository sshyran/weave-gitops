/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/
export type ObjectRef = {
  kind?: string
  name?: string
  namespace?: string
}

export type Condition = {
  type?: string
  status?: string
  reason?: string
  message?: string
  timestamp?: string
}

export type GroupVersionKind = {
  group?: string
  kind?: string
  version?: string
}

export type Object = {
  object?: string
  inventory?: string
  clusterName?: string
}

export type Deployment = {
  name?: string
  namespace?: string
  conditions?: Condition[]
  images?: string[]
  suspended?: boolean
  clusterName?: string
}

export type UnstructuredObject = {
  groupVersionKind?: GroupVersionKind
  name?: string
  namespace?: string
  uid?: string
  status?: string
  conditions?: Condition[]
  suspended?: boolean
  clusterName?: string
  images?: string[]
}

export type Namespace = {
  name?: string
  status?: string
  annotations?: {[key: string]: string}
  labels?: {[key: string]: string}
  clusterName?: string
}

export type Event = {
  type?: string
  reason?: string
  message?: string
  timestamp?: string
  component?: string
  host?: string
  name?: string
}