import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
import { Avatar } from "~/components/avatar/avatar";
import { LinkButton } from "~/components/link-button/link-button";
import type { PlatformType } from "~/models/platform";
import { useLinksLoader, useProfileLoader } from "../layout";
import styles from "./preview.module.css";

export default component$(() => {
  const profile = useProfileLoader();
  const links = useLinksLoader();

  const { firstName, lastName, email, avatar } = profile.value;

  return (
    <div class={styles.root}>
      <div class={styles.accent} />
      <div class={styles.preview}>
        <Avatar src={avatar} view="preview" />

        <div class={styles.name}>
          {firstName} {lastName}
        </div>

        <div class={styles.email}>{email}</div>

        <div class={styles.links}>
          {links.value.links.map((item) => (
            <LinkButton
              type={item.platform as PlatformType}
              link={item.link as string}
              key={item.link}
              view="preview"
            />
          ))}
        </div>
      </div>
    </div>
  );
});

export const head: DocumentHead = {
  title: "Preview Your Links",
};
