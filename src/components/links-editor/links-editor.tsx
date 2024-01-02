import { component$ } from "@builder.io/qwik";

import FormWrapper from "~/components/form-wrapper/form-wrapper";
import LinkItem from "~/components/link-item/link-item";
import styles from "./links-editor.module.css";

export default component$(() => {
  return (
    <FormWrapper formName="linksForm">
      <div q:slot="heading">Customize your links</div>
      <span q:slot="description">
        Add/edit/remove links below and then share all profiles with the world!
      </span>

      <div class={styles.scrollArea}>
        <button class={["button button-outlined", styles.addButton]}>
          + Add new link
        </button>

        <LinkItem index={1} />
        <LinkItem index={2} />
        <LinkItem index={3} />
      </div>
    </FormWrapper>
  );
});
