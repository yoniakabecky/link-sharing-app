import type { QwikIntrinsicElements } from "@builder.io/qwik";
import { icons } from "./icons";

type IconProps = {
  name: keyof typeof icons;
  key?: string;
  size?: number;
} & QwikIntrinsicElements["svg"];

export const Icon = ({ name, size, key, ...props }: IconProps) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width={size ?? 24}
      height={size ?? 24}
      viewBox="0 0 24 24"
      fill="currentColor"
      {...props}
      key={key ?? name}
    >
      <title>{name}</title>
      <path d={icons[name]}></path>
    </svg>
  );
};
