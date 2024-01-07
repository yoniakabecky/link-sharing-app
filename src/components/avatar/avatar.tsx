import { component$, useStylesScoped$ } from "@builder.io/qwik";
import styles from "./avatar.css?inline";

export interface AvatarProps {
  src?: string | null;
  view?: "mock" | "preview";
}

export const Avatar = component$<AvatarProps>(({ src, view = "preview" }) => {
  useStylesScoped$(styles);

  const size = view === "mock" ? 80 : 88;

  if (!src || src === "") return <div class={["avatar skeleton", view]} />;

  return (
    <div class={["avatar", view]}>
      <img src={src} alt="avatar" width={size} height={size} />
    </div>
  );
});
