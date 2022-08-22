import {
  Add as AddIcon,
  ArrowDropDown as ArrowDropDownIcon,
  ArrowUpward as ArrowUpwardIcon,
  CheckCircle as CheckCircleIcon,
  Clear as ClearIcon,
  Delete as DeleteIcon,
  Error as ErrorIcon,
  ExitToApp as LogoutIcon,
  FileCopy as FileCopyIcon,
  Filter as FilterIcon,
  HourglassFull as HourglassFullIcon,
  Launch as LaunchIcon,
  NavigateBefore as NavigateBeforeIcon,
  NavigateNext as NavigateNextIcon,
  Person as PersonIcon,
  RemoveCircle as RemoveCircleIcon,
  SaveAlt as SaveAltIcon,
  Search as SearchIcon,
  SkipNext as SkipNextIcon,
  SkipPrevious as SkipPreviousIcon,
} from "@material-ui/icons";
import * as React from "react";
import styled from "styled-components";
import images from "../lib/images";
// eslint-disable-next-line
import { colors, spacing } from "../typedefs/styled";
import Flex from "./Flex";
import Text from "./Text";

export enum IconType {
  CheckMark,
  Account,
  ExternalTab,
  AddIcon,
  ArrowUpwardIcon,
  ArrowDropDownIcon,
  DeleteIcon,
  SaveAltIcon,
  ErrorIcon,
  CheckCircleIcon,
  HourglassFullIcon,
  NavigateNextIcon,
  NavigateBeforeIcon,
  SkipNextIcon,
  SkipPreviousIcon,
  RemoveCircleIcon,
  FilterIcon,
  ClearIcon,
  SearchIcon,
  LogoutIcon,
  SuccessIcon,
  FailedIcon,
  SuspendedIcon,
  FileCopyIcon,
  ReconcileIcon,
}

type Props = {
  className?: string;
  type: IconType;
  color?: keyof typeof colors;
  text?: string;
  size: keyof typeof spacing;
};

function getIcon(i: IconType) {
  switch (i) {
    case IconType.CheckMark:
      return CheckCircleIcon;

    case IconType.Account:
      return PersonIcon;

    case IconType.ExternalTab:
      return LaunchIcon;

    case IconType.AddIcon:
      return AddIcon;

    case IconType.ArrowUpwardIcon:
      return ArrowUpwardIcon;

    case IconType.DeleteIcon:
      return DeleteIcon;

    case IconType.SaveAltIcon:
      return SaveAltIcon;

    case IconType.CheckCircleIcon:
      return CheckCircleIcon;

    case IconType.HourglassFullIcon:
      return HourglassFullIcon;

    case IconType.ErrorIcon:
      return ErrorIcon;

    case IconType.NavigateNextIcon:
      return NavigateNextIcon;

    case IconType.NavigateBeforeIcon:
      return NavigateBeforeIcon;

    case IconType.SkipNextIcon:
      return SkipNextIcon;

    case IconType.SkipPreviousIcon:
      return SkipPreviousIcon;

    case IconType.RemoveCircleIcon:
      return RemoveCircleIcon;

    case IconType.FilterIcon:
      return FilterIcon;

    case IconType.ClearIcon:
      return ClearIcon;

    case IconType.SearchIcon:
      return SearchIcon;

    case IconType.LogoutIcon:
      return LogoutIcon;

    case IconType.SuccessIcon:
      return () => <img src={images.successSrc} />;

    case IconType.FailedIcon:
      return () => <img src={images.failedSrc} />;

    case IconType.SuspendedIcon:
      return () => <img src={images.suspendedSrc} />;

    case IconType.ReconcileIcon:
      return () => <img src={images.reconcileSrc} />;

    case IconType.ArrowDropDownIcon:
      return ArrowDropDownIcon;

    case IconType.FileCopyIcon:
      return FileCopyIcon;

    default:
      break;
  }
}

function Icon({ className, type, text, color }: Props) {
  return (
    <Flex align className={className}>
      {React.createElement(getIcon(type))}
      {text && (
        <Text color={color} bold>
          {text}
        </Text>
      )}
    </Flex>
  );
}

export default styled(Icon)`
  svg {
    fill: ${(props) => props.theme.colors[props.color as any]} !important;
    height: ${(props) => props.theme.spacing[props.size as any]};
    width: ${(props) => props.theme.spacing[props.size as any]};
  }
  &.downward {
    transform: rotate(180deg);
  }
  &.upward {
    transform: initial;
  }
  ${Text} {
    margin-left: 4px;
    color: ${(props) => props.theme.colors[props.color as any]};
  }
`;
