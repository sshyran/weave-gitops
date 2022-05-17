import * as React from "react";
import { ImageAutomation } from "../lib/api/imageautomation/imageautomation.pb";

type Props = {
  children: any;
  api: typeof ImageAutomation;
};

export type ImageAutomationContextType = { api: typeof ImageAutomation };

const ImageAutomationContext =
  React.createContext<ImageAutomationContextType>(null);

export function useImageAutomation() {
  return React.useContext(ImageAutomationContext);
}

export function ImageAutomationContextProvider({ api, children }: Props) {
  return (
    <ImageAutomationContext.Provider value={{ api }}>
      {children}
    </ImageAutomationContext.Provider>
  );
}
