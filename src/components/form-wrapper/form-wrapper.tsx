import { Slot, component$ } from "@builder.io/qwik";

import styles from "./form-wrapper.module.css";

export default component$(() => {
  return (
    <div class={styles.root}>
      <h1 class={styles.heading}>
        <Slot name="heading" />
      </h1>
      <p class={styles.description}>
        <Slot name="description" />
      </p>

      <div class={styles.children}>
        <Slot />
      </div>

      <div class={styles.action}>
        <button class="save">Save</button>
      </div>
    </div>
  );
});
