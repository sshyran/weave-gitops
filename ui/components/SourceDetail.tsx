import _ from "lodash";
import * as React from "react";
import { useRouteMatch } from "react-router-dom";
import styled from "styled-components";
import { Source, HelmRelease, useListAutomations, useGetObject, Kind } from "../hooks/objects";
import Alert from "./Alert";
import AutomationsTable from "./AutomationsTable";
import DetailTitle from "./DetailTitle";
import EventsTable from "./EventsTable";
import Flex from "./Flex";
import InfoList, { InfoField } from "./InfoList";
import Link from "./Link";
import LoadingPage from "./LoadingPage";
import PageStatus from "./PageStatus";
import SubRouterTabs, { RouterTab } from "./SubRouterTabs";

type Props = {
  className?: string;
  type: Kind;
  name: string;
  namespace: string;
  clusterName: string;
  children?: JSX.Element;
  info: (s: Source) => InfoField[];
};

function SourceDetail({ className, name, namespace, info, type, clusterName }: Props) {
  const { data: source, isLoading, error } = useGetObject<Source>(name, namespace, type, clusterName);
  const { data: automations } = useListAutomations();
  const { path } = useRouteMatch();

  if (isLoading) {
    return <LoadingPage />;
  }


  if (!source) {
    return (
      <Alert
        severity="error"
        title="Not found"
        message={`Could not find source '${name}'`}
      />
    );
  }

  const items = info(source);

  const isNameRelevant = (expectedName: string) => {
    return expectedName == name;
  };

  const isRelevant = (expectedType: string, expectedName: string) => {
    return expectedType == type && isNameRelevant(expectedName);
  };

  const relevantAutomations = _.filter(automations, (a) => {
    if (!source) {
      return false;
    }

    if (type == Kind.HelmChart) {
      return a.kind() == Kind.HelmRelease && isNameRelevant((a as HelmRelease).helmChartRef().name)
    }

    return isRelevant(a?.sourceRef().kind, a?.sourceRef().name);
  });

  return (
    <Flex wide tall column className={className}>
      <DetailTitle name={name} type={type} />
      {error && (
        <Alert severity="error" title="Error" message={error.message} />
      )}
      <PageStatus conditions={source.conditions()} suspended={source.suspended()} />
      <SubRouterTabs rootPath={`${path}/details`}>
        <RouterTab name="Details" path={`${path}/details`}>
          <>
            <InfoList items={items} />
            <AutomationsTable automations={relevantAutomations} hideSource />
          </>
        </RouterTab>
        <RouterTab name="Events" path={`${path}/events`}>
          <EventsTable
            namespace={source.namespace()}
            involvedObject={{
              kind: type,
              name,
              namespace: source.namespace(),
            }}
          />
        </RouterTab>
      </SubRouterTabs>
    </Flex>
  );
}

export default styled(SourceDetail).attrs({ className: SourceDetail.name })`
  width: 100%;

  ${PageStatus} {
    padding: ${(props) => props.theme.spacing.small} 0px;
  }

  .MuiTabs-root ${Link} .active-tab {
    background: ${(props) => props.theme.colors.primary}19;
  }
`;
