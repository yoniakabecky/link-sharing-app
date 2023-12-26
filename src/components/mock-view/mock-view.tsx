import { component$ } from "@builder.io/qwik";
import MobileMock from "~/media/mobile-mock.svg?jsx";
import styles from "./mock-view.module.css";

export default component$(() => {
  return (
    <div class={styles.root}>
      <div class={styles.imageWrapper}>
        <MobileMock />
      </div>
      <div class={styles.contents}>
        <div class={styles.profile}>
          <div class={[styles.picture, styles.skeleton]}></div>
          <div class={[styles.name, styles.skeleton]}></div>
          <div class={[styles.email, styles.skeleton]}></div>
        </div>
        <div class={styles.links}>
          <div class={[styles.link, styles.skeleton]}></div>
          <div class={[styles.link, styles.skeleton]}></div>
          <div class={[styles.link, styles.skeleton]}></div>
          <div class={[styles.link, styles.skeleton]}></div>
          <div class={[styles.link, styles.skeleton]}></div>
        </div>
      </div>
    </div>
  );
});
