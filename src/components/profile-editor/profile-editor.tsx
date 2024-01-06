import { $, type QRL, component$ } from "@builder.io/qwik";
import { type SubmitHandler, useForm, valiForm$ } from "@modular-forms/qwik";
import FormWrapper from "~/components/form-wrapper/form-wrapper";
import { Icon } from "~/components/icon/icon";
import { useProfileLoader, useProfileFormAction } from "~/routes/layout";
import { type ProfileForm, ProfileFormSchema } from "~/routes/schema";
import { TextField } from "../textfield/textfield";
import styles from "./profile-editor.module.css";

export default component$(() => {
  const [, { Form, Field }] = useForm<ProfileForm>({
    loader: useProfileLoader(),
    action: useProfileFormAction(),
    validate: valiForm$(ProfileFormSchema),
  });

  const handleSubmit: QRL<SubmitHandler<ProfileForm>> = $((values, event) => {
    // Runs on client
    console.log(values, event);
  });

  return (
    <FormWrapper formName="ProfileForm">
      <div q:slot="heading">Profile Details</div>
      <span q:slot="description">
        Add your details to create a personal touch to your profile.
      </span>

      <div class={styles.pictureWrapper}>
        <p>Profile Picture</p>
        <div class={styles.imageWrapper}>
          <div class={styles.overlay}>
            <Icon name="pic_line" size={32} />
            <div>Change Image</div>
          </div>
        </div>
        <div>
          <small>Image must be below 1024x1024px.</small>
          <small>Use PNG, JPG, or BMP format.</small>
        </div>
      </div>

      <Form id="ProfileForm" onSubmit$={handleSubmit} class={styles.form}>
        <label for="firstName" class="text-md">
          First name*
        </label>
        <Field name="firstName">
          {(field, props) => (
            <TextField
              {...props}
              required
              name="firstName"
              value={field.value}
              error={field.error}
            />
          )}
        </Field>

        <label for="lastName" class="text-md">
          Last name*
        </label>
        <Field name="lastName">
          {(field, props) => (
            <TextField
              {...props}
              required
              name="lastName"
              value={field.value}
              error={field.error}
            />
          )}
        </Field>

        <label for="email" class="text-md">
          Email
        </label>
        <Field name="email">
          {(field, props) => (
            <TextField
              {...props}
              name="email"
              type="email"
              value={field.value ?? ""}
              error={field.error}
            />
          )}
        </Field>
      </Form>
    </FormWrapper>
  );
});
