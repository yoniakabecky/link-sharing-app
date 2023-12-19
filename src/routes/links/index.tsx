import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
import styles from "./links.module.css";

export default component$(() => {
  return (
    <div class={styles.root}>
      <div class="card left">left</div>
      <div class="card right">right</div>
    </div>
  );
});

export const head: DocumentHead = {
  title: "Links",
};
