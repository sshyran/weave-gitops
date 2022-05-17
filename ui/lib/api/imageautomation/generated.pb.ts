/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/
export type CommitSpec = {
  author?: CommitUser
  signingKey?: SigningKey
  messageTemplate?: string
}

export type CommitUser = {
  name?: string
  email?: string
}

export type CrossNamespaceSourceReference = {
  apiversion?: string
  kind?: string
  name?: string
  namespace?: string
}

export type GitCheckoutSpec = {
}

export type GitSpec = {
  checkout?: GitCheckoutSpec
  commit?: CommitSpec
  push?: PushSpec
}

export type ImageUpdateAutomation = {
  kind?: string
  apiversion?: string
  name?: string
  generateName?: string
  namespace?: string
  selfLink?: string
  resourceVersion?: string
  generation?: string
  deletionGracePeriodSeconds?: string
  labels?: {[key: string]: string}
  annotations?: {[key: string]: string}
  finalizers?: string[]
  clusterName?: string
  spec?: ImageUpdateAutomationSpec
  status?: ImageUpdateAutomationStatus
}

export type ImageUpdateAutomationSpec = {
  sourceRef?: CrossNamespaceSourceReference
  gitSpec?: GitSpec
  update?: UpdateStrategy
  suspend?: boolean
}

export type ImageUpdateAutomationStatus = {
  lastPushCommit?: string
  observedGeneration?: string
  lastHandledReconcileAt?: string
}

export type PushSpec = {
  branch?: string
}

export type SigningKey = {
}

export type UpdateStrategy = {
  strategy?: string
  path?: string
}