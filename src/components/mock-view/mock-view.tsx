import { component$ } from "@builder.io/qwik";
import { useLocation } from "@builder.io/qwik-city";
import MobileMock from "~/media/mobile-mock.svg?jsx";
import type { PlatformType } from "~/models/platform";
import { useLinksFormLoader, useProfileFormLoader } from "~/routes/layout";
import LinkButton from "../link-button/link-button";
import styles from "./mock-view.module.css";

const MIN_LINKS = 5;

export default component$(() => {
  const profile = useProfileFormLoader();
  const links = useLinksFormLoader();
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
          {avatar && hasValue(avatar) ? (
            <div class={styles.picture}>
              <img src={avatar as string} alt="avatar" width={80} height={80} />
            </div>
          ) : (
            <div class={[styles.picture, styles.skeleton]} />
          )}
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
