import { component$ } from "@builder.io/qwik";
import { useLocation } from "@builder.io/qwik-city";
import { Avatar } from "~/components/avatar/avatar";
import { LinkButton } from "~/components/link-button/link-button";
import MobileMock from "~/media/mobile-mock.svg?jsx";
import type { PlatformType } from "~/models/platform";
import { useLinksLoader, useProfileLoader } from "~/routes/layout";
import styles from "./mock-view.module.css";

const MIN_LINKS = 5;

export default component$(() => {
  const profile = useProfileLoader();
  const links = useLinksLoader();
  const location = useLocation();
  const isLinksPage = location.url.pathname === "/links/";

  const { firstName, lastName, email, avatar } = profile.value;
  const hasValue = (str: string | null | undefined) => str !== "" || !!str;
  console.log(avatar);
  return (
    <div class={styles.root}>
      <div class={styles.imageWrapper}>
        <MobileMock />
      </div>
      <div class={styles.contents}>
        <div class={styles.profile}>
          <Avatar src={avatar as string} view="mock" />

          <div
            class={[
              styles.name,
              (!hasValue(firstName) || !hasValue(lastName)) && styles.skeleton,
            ]}
          >
            {firstName} {lastName}
          </div>
          <div class={[styles.email, !hasValue(email) && styles.skeleton]}>
            {email}
          </div>
        </div>

        <div class={styles.links}>
          {links.value.links.map((item) => (
            <LinkButton
              type={item.platform as PlatformType}
              link={item.link as string}
              key={item.link}
              view="mock"
            />
          ))}
          {isLinksPage &&
            links.value.links.length < MIN_LINKS &&
            new Array(MIN_LINKS - links.value.links.length)
              .fill(0)
              .map((_, i) => <div class={styles.linkSkeleton} key={i} />)}
        </div>
      </div>
    </div>
  );
});
