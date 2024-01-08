import { component$ } from "@builder.io/qwik";
import { Avatar } from "~/components/avatar/avatar";
import { LinkButton } from "~/components/link-button/link-button";
import MobileMock from "~/media/mobile-mock.svg?jsx";
import type { PlatformType } from "~/models/platform";
import { useLinksLoader, useProfileLoader } from "~/routes/layout";
import type { LinkItem, ProfileForm } from "~/routes/schema";
import styles from "./mock-view.module.css";

interface MockViewProps {
  profile?: ProfileForm;
  links?: LinkItem[];
}

export default component$((props: MockViewProps) => {
  const profile = useProfileLoader();
  const links = useLinksLoader();

  const { firstName, lastName, email, avatar } = props.profile ?? profile.value;
  const hasValue = (str: string | null | undefined) => str !== "" || !!str;

  const linksArray = props.links ?? links.value.links;

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
          {linksArray.map((item) => (
            <LinkButton
              type={item.platform as PlatformType}
              link={item.link as string}
              key={item.link}
              view="mock"
            />
          ))}
        </div>
      </div>
    </div>
  );
});
