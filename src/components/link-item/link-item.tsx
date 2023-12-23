import { component$ } from "@builder.io/qwik";

import styles from "./link-item.module.css";

export default component$(({ index }: { index: number }) => {
  return (
    <div class={styles.root}>
      <div class={styles.title}>
        <span></span>Link #{index}
      </div>
      <div class={styles.remove}>Remove</div>

      <label for={`platform${index}`}>Platform</label>
      <select name={`platform${index}`} id={`platform${index}`}>
        <option value="github">GitHub</option>
        <option value="youtube">YouTube</option>
        <option value="linkedin">Linkedin</option>
        <option value="x">X (Twitter)</option>
        <option value="portfolio">Portfolio</option>
        <option value="custom">Custom</option>
      </select>

      <label for={`link${index}`}>Link</label>
      <input
        id={`link${index}`}
        type="text"
        placeholder="https://www.yoniakabecky.com"
      />
    </div>
  );
});
