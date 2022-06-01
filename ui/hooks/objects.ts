import _ from "lodash";
import { useContext } from "react";
import { useQuery } from "react-query";
import { CoreClientContext } from "../contexts/CoreClientContext";
import { NoNamespace, RequestError } from "../lib/types";
import { Object as ResponseObject } from "../lib/api/core/types.pb";
import { GetObjectResponse, ListObjectsResponse } from "../lib/api/core/core.pb";

export enum Kind {
  GitRepository = "GitRepository",
  Bucket = "Bucket",
  HelmRepository = "HelmRepository",
  HelmChart = "HelmChart",
  Kustomization = "Kustomization",
  HelmRelease = "HelmRelease",
}

const Automations = [Kind.Kustomization, Kind.HelmRelease];
const Sources = [Kind.GitRepository, Kind.HelmRepository, Kind.Bucket, Kind.HelmChart];

export class FluxObject {
  cluster: string
  obj: any

  clusterName() {
    return this.cluster
  }

  metadata() {
    const annotations = this.obj.metadata.annotations || {}
    return Object.keys(annotations).filter(key => {
      return key.startsWith('metadata.weave.works/')
    }).map(key => [key.slice(21), annotations[key]])
  }

  kind() {
    return this.obj.kind
  }

  name() {
    return this.obj.metadata.name
  }

  namespace() {
    return this.obj.metadata.namespace
  }

  conditions() {
    return this.obj.status.conditions
  }

  lastUpdated() {
    return _.get(_.find(this.conditions(), { type: "Ready" }), "lastTransitionTime");
  }

  interval() {
    return this.obj.spec.interval
  }

  suspended() {
    return this.obj.spec.suspend
  }
}

export class Source extends FluxObject {
  url() {
    return this.obj.spec.url
  }
}

export class Bucket extends Source {
  endpoint() {
    return this.obj.spec.endpoint
  }
}

export class GitRepository extends Source {
  reference() {
    const refObj = this.obj.spec.ref;
    return refObj?.branch ||
      refObj?.commit ||
      refObj?.tag ||
      refObj?.semver;
  }
}

export class HelmChart extends Source {
  chartName() {
    return this.obj.spec.chart
  }

  sourceRef() {
    const sourceRef = this.obj.spec.sourceRef;
    sourceRef.namespace = this.namespace();
    return sourceRef;
  }
}

export class HelmRepository extends Source {
}

export class Automation extends FluxObject {

  lastAttemptedRevision() {
    return this.obj.status.lastAttemptedRevision
  }

  lastAppliedRevision() {
    return this.obj.status.lastAppliedRevision
  }

  inventory() { return [] }

  sourceRef() {
    let sourceRef;
    if (this.kind() == Kind.HelmRelease) {
      sourceRef = this.obj.spec.chart.spec.sourceRef;
    } else {
      sourceRef = this.obj.spec.sourceRef;
    }
    if (!sourceRef.namespace) {
      sourceRef.namespace = this.namespace();
    }
    return sourceRef;
  }

}

export class Kustomization extends Automation {
  path() {
    return this.obj.spec.path
  }

  inventory() {
    const inventory = new Set();
    for (const inv of this.obj.status.inventory.entries) {
      const parts = inv.id.split('_');
      const group = parts[2];
      const kind = parts[3];
      inventory.add({ group, kind, version: inv.v });
    }
    return Array.from(inventory);
  }
}

export class HelmRelease extends Automation {
  rawInventory: any
  /* This is the name of the chart itself - not the name of the flux object */
  helmChartName() {
    return this.obj.status.helmChart
  }

  helmChartRef() {
    return {
      name: this.namespace() + '-' + this.name(),
      namespace: this.sourceRef().namespace,
      kind: Kind.HelmChart,
    }
  }

  inventory() {
    return this.rawInventory?.map(({ Group, Kind, Version }) => { return { group: Group, kind: Kind, version: Version } })
  }
}

function convertResponse(response: ResponseObject): FluxObject {
  const obj = JSON.parse(response.object);
  let result;
  switch (obj.kind) {
    case Kind.Kustomization:
      result = new Kustomization()
      break;
    case Kind.HelmRelease:
      result = new HelmRelease()
      if (response.inventory) {
        result.rawInventory = JSON.parse(response.inventory);
      }
      break;
    case Kind.GitRepository:
      result = new GitRepository()
      break
    case Kind.Bucket:
      result = new Bucket()
      break
    case Kind.HelmRepository:
      result = new HelmRepository()
      break
    case Kind.HelmChart:
      result = new HelmChart()
      break
  }
  result.obj = obj
  result.cluster = response.clusterName
  return result
}

function useListObjects<K extends FluxObject>(namespace = NoNamespace, kinds) {
  const { api } = useContext(CoreClientContext);

  return useQuery<K[], RequestError>(
    ["objects", ...kinds],
    () => {
      const p = kinds.map((kind) => api.ListObjects({ namespace, kind }))

      return Promise.all(p).then((results) => {
        return results.flatMap((response: ListObjectsResponse) => {
          return response.objects.map((resp) => {
            return convertResponse(resp) as K
          });
        });
      });
    },
    { retry: false, refetchInterval: 5000 }
  );
}

export function useListAutomations(namespace = NoNamespace) {
  return useListObjects<Automation>(namespace, Automations);
}

export function useListSources(namespace = NoNamespace) {
  return useListObjects<Source>(namespace, Sources);
}

export function useGetObject<K extends FluxObject>(
  name: string,
  namespace: string,
  kind: string,
  clusterName: string,
) {
  const { api } = useContext(CoreClientContext);

  return useQuery<K, RequestError>(
    ["object", name],
    () => api.GetObject({ name, namespace, kind, clusterName }).then(
      (result: GetObjectResponse) => (convertResponse(result.object) as K)
    ),
    { retry: false, refetchInterval: 5000 }
  )
}
