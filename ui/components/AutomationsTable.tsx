import _ from "lodash";
import * as React from "react";
import styled from "styled-components";
import { Automation, Kind } from "../hooks/objects";
import { formatURL } from "../lib/nav";
import { V2Routes } from "../lib/types";
import { statusSortHelper } from "../lib/utils";
import { Field, SortType } from "./DataTable";
import {
  filterConfigForStatus,
  filterConfigForString,
} from "./FilterableTable";
import KubeStatusIndicator, { computeMessage } from "./KubeStatusIndicator";
import Link from "./Link";
import SourceLink from "./SourceLink";
import Timestamp from "./Timestamp";
import URLAddressableTable from "./URLAddressableTable";

type Props = {
  className?: string;
  automations?: Automation[];
  appName?: string;
  hideSource?: boolean;
};

function AutomationsTable({ className, automations, hideSource }: Props) {
  const filterConfig = {
    ...filterConfigForString(automations, "kind"),
    ...filterConfigForString(automations, "namespace"),
    ...filterConfigForString(automations, "clusterName"),
    ...filterConfigForStatus(automations),
  };

  let fields: Field[] = [
    {
      label: "Name",
      value: (k) => {
        const route =
          k.kind() === Kind.Kustomization
            ? V2Routes.Kustomization
            : V2Routes.HelmRelease;
        return (
          <Link
            to={formatURL(route, {
              name: k.name(),
              namespace: k.namespace(),
              clusterName: k.clusterName(),
            })}
          >
            {k.name()}
          </Link>
        );
      },
      sortValue: (a) => a.name(),
      textSearchable: true,
      maxWidth: 600,
    },
    {
      label: "Kind",
      value: (a) => a.kind(),
    },
    {
      label: "Namespace",
      value: (a) => a.namespace(),
    },
    {
      label: "Cluster",
      value: (a) => a.clusterName(),
    },
    {
      label: "Source",
      value: (a: Automation) => {
        return (
          <SourceLink
            short
            sourceRef={a.sourceRef()}
            clusterName={a.clusterName()}
          />
        );
      },
      sortValue: (a: Automation) => a.sourceRef().name,
    },
    {
      label: "Status",
      value: (a: Automation) =>
        a.conditions().length > 0 ? (
          <KubeStatusIndicator
            short
            conditions={a.conditions()}
            suspended={a.suspended()}
          />
        ) : null,
      sortType: SortType.number,
      sortValue: statusSortHelper,
    },
    {
      label: "Message",
      value: (a: Automation) => computeMessage(a.conditions()),
      sortValue: (a: Automation) => computeMessage(a.conditions()),
      maxWidth: 600,
    },
    {
      label: "Revision",
      value: (a: Automation) => a.lastAttemptedRevision(),
    },
    {
      label: "Last Updated",
      value: (a: Automation) => (
        <Timestamp
          time={a.lastUpdated()}
        />
      ),
    },
  ];

  if (hideSource) fields = _.filter(fields, (f) => f.label !== "Source");

  return (
    <URLAddressableTable
      fields={fields}
      filters={filterConfig}
      rows={automations}
      className={className}
    />
  );
}

export default styled(AutomationsTable).attrs({
  className: AutomationsTable.name,
})``;
