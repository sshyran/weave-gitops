import * as React from "react";
import styled from "styled-components";
import Link from "../components/Link";
import SourceDetail from "../components/SourceDetail";
import Timestamp from "../components/Timestamp";
import { Kind, GitRepository } from "../hooks/objects";
import { convertGitURLToGitProvider } from "../lib/utils";

type Props = {
  className?: string;
  name: string;
  namespace: string;
  clusterName: string;
};

function GitRepositoryDetail({ name, namespace, clusterName, className }: Props) {
  return (
    <SourceDetail
      className={className}
      name={name}
      namespace={namespace}
      type={Kind.GitRepository}
      clusterName={clusterName}
      info={(s: GitRepository) => [
        ["Type", s.kind()],
        ["URL",
          <Link newTab href={convertGitURLToGitProvider(s.url())} >
            {s.url}
          </Link >,
        ],
        ["Ref", s.reference()],
        ["Last Updated", <Timestamp time={s.lastUpdated()} />],
        ["Cluster", s.clusterName()],

        ["Namespace", s.namespace()],
      ]}
    />
  );
}

export default styled(GitRepositoryDetail).attrs({
  className: GitRepositoryDetail.name,
})``;
