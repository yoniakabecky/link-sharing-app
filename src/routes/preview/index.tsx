import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
import styles from "./preview.module.css";

export default component$(() => {
  return (
    <div class={styles.root}>
      <div class={styles.accent} />
      <div class={styles.preview}>Preview</div>
    </div>
  );
});

export const head: DocumentHead = {
  title: "Preview Your Links",
};
