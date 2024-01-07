import { component$, useStylesScoped$ } from "@builder.io/qwik";
import { type PlatformType, platforms } from "~/models/platform";
import { Icon } from "../icon/icon";
import styles from "./link-button.css?inline";

interface Props {
  type: PlatformType;
  link: string;
  view?: "mock" | "preview";
}

export const LinkButton = component$(
  ({ type, link, view = "preview" }: Props) => {
    useStylesScoped$(styles);

    const { icon, label, color } = platforms[type];
    const iconSize = view === "mock" ? 16 : 20;

    return (
      <div class={["root", view]} style={{ backgroundColor: color }}>
        <a href={link} target="_blank">
          <Icon name={icon} size={iconSize} />
          <span class="label">{label}</span>
          <Icon name="arrow_right" size={16} class="arrow" />
        </a>
      </div>
    );
  }
);
