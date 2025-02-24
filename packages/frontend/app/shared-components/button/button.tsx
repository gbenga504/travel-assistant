import { NavLink } from "@remix-run/react";
import classNames from "classnames";

import { omit } from "~/utils/functional";

import { Spinner } from "../spinner";

import type { NavLinkProps } from "@remix-run/react";
import type React from "react";
import type { ReactElement } from "react";

import "./button.css";

interface CommonProps {
  variant?: "contained" | "text";
  size?: "small" | "medium" | "large";
  colorTheme?: "primary" | "white";
  fullWidth?: boolean;
  icon?: ReactElement;
  shape?: "default" | "rounded" | "circle";
  disabled?: boolean;
}

interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    CommonProps {
  type: "button";
  loading?: boolean;
  loadingText?: string;
}

interface AnchorProps
  extends React.AnchorHTMLAttributes<HTMLAnchorElement>,
    CommonProps {
  type: "anchor";
}

interface LinkProps extends NavLinkProps, CommonProps {
  type: "link";
}

export const Button: React.FC<ButtonProps | AnchorProps | LinkProps> = (
  props
) => {
  const {
    className,
    type,
    variant = "contained",
    size = "medium",
    colorTheme = "primary",
    fullWidth,
    icon,
    shape,
    disabled,
  } = props;

  const buttonClassName = classNames(
    "button",
    {
      "button-contained": variant === "contained",
      "button-text": variant === "text",
      "button-large": size === "large",
      "button-medium": size === "medium",
      "button-small": size === "small",
      "button-primary": colorTheme === "primary",
      "button-white": colorTheme === "white",
      "button-rounded": shape === "rounded",
      "button-circle": shape === "circle",
      "button-fullWidth": fullWidth,
      "button-disabled": disabled,
    },
    className
  );

  const commonKeys: (keyof CommonProps)[] = [
    "variant",
    "size",
    "colorTheme",
    "fullWidth",
    "icon",
    "shape",
  ];

  const renderIcon = () => {
    if (icon) {
      return <span className="mr-1">{icon}</span>;
    }

    return null;
  };

  const renderButton = (props: ButtonProps) => {
    const { loading, disabled, loadingText, children } = props;

    return (
      <button
        {...omit(props, [...commonKeys, "type", "loading", "loadingText"])}
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
      <a {...omit(props, commonKeys)} className={buttonClassName}>
        {renderIcon()}
        <span>{props.children}</span>
      </a>
    );
  };

  const renderLink = (props: LinkProps) => {
    const { className, children } = props;

    return (
      <NavLink
        {...omit(props, commonKeys)}
        className={(p) => {
          if (typeof className === "string") {
            return classNames(buttonClassName, className);
          }

          return classNames(buttonClassName, className?.(p));
        }}
      >
        {(p) => {
          if (typeof children === "function") {
            return (
              <div className="inline-flex items-center">
                {renderIcon()}
                {children(p)}
              </div>
            );
          }

          return (
            <div className="inline-flex items-center">
              {renderIcon()}
              {children}
            </div>
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
