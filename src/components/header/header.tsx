import { component$ } from "@builder.io/qwik";
import styles from "./header.module.css";

export default component$(() => {
  return (
    <header class={styles.header}>
      <div class={styles.outer}>
        <div class={["card", styles.inner]}>
          <div class={styles.logo}>
            <a href="/">logo</a>
          </div>
          <ul>
            <li>
              <a href="/links">Links</a>
            </li>
            <li>
              <a href="/profile">Profile Details</a>
            </li>
          </ul>
          <div>
            <a href="/preview">Preview</a>
          </div>
        </div>
      </div>
    </header>
  );
});
