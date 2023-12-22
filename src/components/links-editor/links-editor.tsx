import { component$ } from "@builder.io/qwik";

import FormWrapper from "~/components/form-wrapper/form-wrapper";
import styles from "./links-editor.module.css";

export default component$(() => {
  return (
    <FormWrapper>
      <div q:slot="heading">Customize your links</div>
      <span q:slot="description">
        Add/edit/remove links below and then share all profiles with the world!
      </span>

      <button class={["button button-outlined", styles.add]}>
        + Add new link
      </button>

      <div>
        <div>Link #1</div>
        <div>remove</div>
        <label for="platform">Platform</label>
        <input id="platform" type="text" />

        <label for="link">Link</label>
        <input id="link" type="text" />
      </div>
    </FormWrapper>
  );
});
