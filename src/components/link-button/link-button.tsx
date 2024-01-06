import { component$, useStylesScoped$ } from "@builder.io/qwik";
import { type PlatformType, platforms } from "~/models/platform";
import { Icon } from "../icon/icon";
import styles from "./link-button.css?inline";

interface Props {
  type: PlatformType;
  link: string;
}
export default component$(({ type, link }: Props) => {
  useStylesScoped$(styles);
  const { icon, label, color } = platforms[type];

  return (
    <div class="root" style={{ backgroundColor: color }}>
      <a href={link} target="_blank">
        <Icon name={icon} size={16} />
        <span class="label">{label}</span>
        <Icon name="arrow_right" size={16} class="arrow" />
      </a>
    </div>
  );
});
