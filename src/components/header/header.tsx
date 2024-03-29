import { Slot, component$ } from "@builder.io/qwik";
import type { LinkProps } from "@builder.io/qwik-city";
import { Link, useLocation } from "@builder.io/qwik-city";
import { Icon } from "~/components/icon/icon";
import Logo from "~/media/logo.svg?jsx";
import styles from "./header.module.css";

export const MenuItem = component$<LinkProps>(({ href }) => {
  const location = useLocation();
  const isActive = location.url.pathname === href;

  return (
    <li>
      <Link href={href} class={["button button-light", { active: isActive }]}>
        <Slot />
      </Link>
    </li>
  );
});

export default component$(() => {
  const location = useLocation();
  const isPreview = location.url.pathname === "/preview/";

  return (
    <header class={styles.header}>
      <div class={styles.outer}>
        {isPreview ? (
          <div class={["card", styles.inner]}>
            <div>
              <Link href="/links/" class="button button-outlined">
                Back to Editor
              </Link>
            </div>
            <div>
              <Link href="/" class="button">
                Share Link
              </Link>
            </div>
          </div>
        ) : (
          <div class={["card", styles.inner]}>
            <Link href="/" class={styles.logo}>
              <Logo />
              devlinks
            </Link>
            <ul>
              <MenuItem href="/links/">
                <Icon name="link" size={16} />
                Links
              </MenuItem>
              <MenuItem href="/profile/">
                <Icon name="user" size={16} />
                Profile Details
              </MenuItem>
            </ul>
            <div>
              <Link href="/preview/" class="button button-outlined">
                Preview
              </Link>
            </div>
          </div>
        )}
      </div>
    </header>
  );
});
