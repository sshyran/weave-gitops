import * as React from "react";
import styled from "styled-components";
import { Kind, Source, Bucket, GitRepository } from "../hooks/objects";
import { formatURL, sourceTypeToRoute } from "../lib/nav";
import { convertGitURLToGitProvider, statusSortHelper } from "../lib/utils";
import { SortType } from "./DataTable";
import {
  filterConfigForStatus,
  filterConfigForString,
} from "./FilterableTable";
import KubeStatusIndicator, { computeMessage } from "./KubeStatusIndicator";
import Link from "./Link";
import Timestamp from "./Timestamp";
import URLAddressableTable from "./URLAddressableTable";

type Props = {
  className?: string;
  sources?: Source[];
  appName?: string;
};

function SourcesTable({ className, sources }: Props) {
  const [filterDialogOpen, setFilterDialog] = React.useState(false);

  const initialFilterState = {
    ...filterConfigForString(sources, "kind"),
    ...filterConfigForString(sources, "namespace"),
    ...filterConfigForStatus(sources),
    ...filterConfigForString(sources, "clusterName"),
  };

  return (
    <URLAddressableTable
      className={className}
      filters={initialFilterState}
      rows={sources}
      dialogOpen={filterDialogOpen}
      onDialogClose={() => setFilterDialog(false)}
      fields={[
        {
          label: "Name",
          value: (s: Source) => (
            <Link
              to={formatURL(sourceTypeToRoute(s.kind()), {
                name: s?.name(),
                namespace: s?.namespace(),
                clusterName: s.clusterName(),
              })}
            >
              {s?.name()}
            </Link>
          ),
          sortType: SortType.string,
          sortValue: (s: Source) => s.name() || "",
          textSearchable: true,
          maxWidth: 600,
        },
        { label: "Kind", value: (s) => s.kind() },
        { label: "Namespace", value: (s) => s.namespace() },
        {
          label: "Status",
          value: (s: Source) => (
            <KubeStatusIndicator
              short
              conditions={s.conditions()}
              suspended={s.suspended()}
            />
          ),
          sortType: SortType.number,
          sortValue: statusSortHelper,
        },
        {
          label: "Message",
          value: (s) => computeMessage(s.conditions()),
          maxWidth: 600,
        },
        {
          label: "Cluster",
          value: (s: Source) => s.clusterName(),
        },
        {
          label: "URL",
          value: (s: Source) => {
            let text;
            let url;
            let link = false;
            switch (s.kind()) {
              case Kind.GitRepository:
                text = s.url();
                url = convertGitURLToGitProvider(s.url());
                link = true;
                break;
              case Kind.Bucket:
                text = (s as Bucket).endpoint();
                break;
              case Kind.HelmChart:
                text = '-';
                break;
              case Kind.HelmRepository:
                text = s.url();
                url = text;
                link = true;
                break;
            }
            return link ? (
              <Link newTab href={url}>
                {text}
              </Link>
            ) : (
              text
            );
          },
          maxWidth: 600,
        },
        {
          label: "Reference",
          value: (s: Source) => {
            if (s.kind() === Kind.GitRepository) {
              return (s as GitRepository).reference()
            }
            return '-'
          },
        },
        {
          label: "Interval",
          value: (s: Source) => s.interval(),
        },
        {
          label: "Last Updated",
          value: (s: Source) => (
            <Timestamp time={s.lastUpdated()} />
          ),
        },
      ]}
    />
  );
}

export default styled(SourcesTable).attrs({ className: SourcesTable.name })``;
