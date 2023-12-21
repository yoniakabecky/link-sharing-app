import { Slot, component$ } from "@builder.io/qwik";
import type { LinkProps } from "@builder.io/qwik-city";
import { Link, useLocation } from "@builder.io/qwik-city";
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
  return (
    <header class={styles.header}>
      <div class={styles.outer}>
        <div class={["card", styles.inner]}>
          <div class={styles.logo}>
            <a href="/">logo</a>
          </div>
          <ul>
            <MenuItem href="/links/">Links</MenuItem>
            <MenuItem href="/profile/">Profile Details</MenuItem>
          </ul>
          <div>
            <Link href="/preview" class="button button-outlined">
              Preview
            </Link>
          </div>
        </div>
      </div>
    </header>
  );
});
