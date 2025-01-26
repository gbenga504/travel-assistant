import classNames from "classnames";
import React from "react";
import { NavLink, NavLinkProps } from "@remix-run/react";

import "./button.css";

import type { ReactElement } from "react";
import { Spinner } from "../spinner";

type variant = "contained" | "text";
type size = "small" | "medium" | "large";
type color = "primary";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  type: "button";

  variant?: variant;
  size?: size;
  color?: color;
  fullWidth?: boolean;
  icon?: ReactElement;

  loading?: boolean;
  loadingText?: string;
}

interface AnchorProps extends React.AnchorHTMLAttributes<HTMLAnchorElement> {
  type: "anchor";

  variant?: variant;
  size?: size;
  color?: color;
  fullWidth?: boolean;
  icon?: ReactElement;
}

interface LinkProps extends NavLinkProps {
  type: "link";

  variant?: variant;
  size?: size;
  color?: color;
  fullWidth?: boolean;
  icon?: ReactElement;
}

export const Button: React.FC<ButtonProps | AnchorProps | LinkProps> = (
  props
) => {
  const {
    className,
    type,
    variant = "contained",
    size = "medium",
    color = "primary",
    fullWidth,
    icon,
  } = props;

  const buttonClassName = classNames(className, "button", {
    "button-contained": variant === "contained",
    "button-text": variant === "text",
    "button-large": size === "large",
    "button-medium": size === "medium",
    "button-small": size === "small",
    "button-primary": color === "primary",
    "button-fullWidth": fullWidth === true,
  });

  const renderIcon = () => {
    if (icon) {
      return <span className="mr-1">{icon}</span>;
    }

    return null;
  };

  const renderButton = (props: ButtonProps) => {
    const { loading, disabled, loadingText, children, ...rest } = props;

    return (
      <button
        {...rest}
        className={buttonClassName}
        disabled={loading || disabled}
      >
        {loading ? (
          <>
            <Spinner />
            <span className="ml-2">{loadingText || "Loading"}</span>
          </>
        ) : (
          <>
            {renderIcon()}
            {children}
          </>
        )}
      </button>
    );
  };

  const renderAnchor = (props: AnchorProps) => {
    return (
      <a {...props} className={buttonClassName}>
        {renderIcon()}
        <span>{props.children}</span>
      </a>
    );
  };

  const renderLink = (props: LinkProps) => {
    return (
      <NavLink {...props} className={buttonClassName}>
        {(p) => {
          if (typeof props.children === "function") {
            return (
              <>
                {renderIcon()}
                {props.children(p)}
              </>
            );
          }

          return (
            <>
              {renderIcon()}
              {props.children}
            </>
          );
        }}
      </NavLink>
    );
  };

  switch (type) {
    case "button":
      return renderButton(props);
    case "anchor":
      return renderAnchor(props);
    case "link":
      return renderLink(props);
  }
};
