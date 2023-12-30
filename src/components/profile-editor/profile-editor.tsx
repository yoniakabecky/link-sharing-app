import { component$ } from "@builder.io/qwik";
import FormWrapper from "~/components/form-wrapper/form-wrapper";
import styles from "./profile-editor.module.css";

export default component$(() => {
  return (
    <FormWrapper>
      <div q:slot="heading">Profile Details</div>
      <span q:slot="description">
        Add your details to create a personal touch to your profile.
      </span>

      <div class={styles.pictureWrapper}>
        <p>Profile Picture</p>
        <div class={styles.imageWrapper}>
          <div class={styles.overlay}>
            <div>icon</div>
            <div>Change Image</div>
          </div>
        </div>
        <div>
          <small>Image must be below 1024x1024px.</small>
          <small>Use PNG, JPG, or BMP format.</small>
        </div>
      </div>

      <form class={styles.form}>
        <p aria-label="first-name">First name*</p>
        <input id="first-name" />
        <p aria-label="last-name">Last name*</p>
        <input id="last-name" />
        <p aria-label="email">Email</p>
        <input id="email" type="email" />
      </form>
    </FormWrapper>
  );
});
