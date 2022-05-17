/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as Image_automationV1Generated from "./generated.pb"
export type ListImageAutomationRequest = {
}

export type ListImageAutomationResponse = {
  automations?: Image_automationV1Generated.ImageUpdateAutomation[]
}

export class ImageAutomation {
  static ListImageAutomations(req: ListImageAutomationRequest, initReq?: fm.InitReq): Promise<ListImageAutomationResponse> {
    return fm.fetchReq<ListImageAutomationRequest, ListImageAutomationResponse>(`/v1/image_automations?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}