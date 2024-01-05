import { component$ } from "@builder.io/qwik";
import { useLocation } from "@builder.io/qwik-city";
import MobileMock from "~/media/mobile-mock.svg?jsx";
import { useLinksFormLoader, useProfileFormLoader } from "~/routes/layout";
import styles from "./mock-view.module.css";

export default component$(() => {
  const profile = useProfileFormLoader();
  const links = useLinksFormLoader();
  const location = useLocation();
  const isLinksPage = location.url.pathname === "/links/";

  const { firstName, lastName, email } = profile.value;
  const hasValue = (str: string | null | undefined) => str !== "" || !!str;

  return (
    <div class={styles.root}>
      <div class={styles.imageWrapper}>
        <MobileMock />
      </div>
      <div class={styles.contents}>
        <div class={styles.profile}>
          <div class={[styles.picture, styles.skeleton]}></div>
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
            <div class={styles.link} key={item.link}>
              {item.platform}
            </div>
          ))}
          {isLinksPage && <div class={[styles.link, styles.skeleton]} />}
        </div>
      </div>
    </div>
  );
});
