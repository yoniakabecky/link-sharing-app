import { $, type QRL, component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
import type { SubmitHandler } from "@modular-forms/qwik";
import { useForm, valiForm$, getValues } from "@modular-forms/qwik";
import FormWrapper from "~/components/form-wrapper/form-wrapper";
import { Icon } from "~/components/icon/icon";
import MockView from "~/components/mock-view/mock-view";
import { TextField } from "~/components/textfield/textfield";
import { useProfileLoader, useProfileFormAction } from "../layout";
import { type ProfileForm, ProfileFormSchema } from "../schema";
import styles from "./profile.module.css";

export default component$(() => {
  const [profileForm, { Form, Field }] = useForm<ProfileForm>({
    loader: useProfileLoader(),
    action: useProfileFormAction(),
    validate: valiForm$(ProfileFormSchema),
  });
  const values = getValues(profileForm);

  const handleSubmit: QRL<SubmitHandler<ProfileForm>> = $((values, event) => {
    // Runs on client
    console.log(values, event);
  });

  return (
    <div class="editor">
      <div class="card left">
        <MockView profile={values as ProfileForm} />
      </div>

      <div class="card right">
        <FormWrapper formName="ProfileForm">
          <div q:slot="heading">Profile Details</div>
          <span q:slot="description">
            Add your details to create a personal touch to your profile.
          </span>

          <Form id="ProfileForm" onSubmit$={handleSubmit}>
            <div class={styles.pictureWrapper}>
              <p>Profile Picture</p>
              <div class={styles.imageWrapper}>
                <Field name="avatar">
                  {(field, props) => (
                    <input
                      {...props}
                      type="file"
                      accept="image/*"
                      value={field.value}
                    />
                  )}
                </Field>

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

            <div class={styles.inputWrapper}>
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
            </div>
          </Form>
        </FormWrapper>
      </div>
    </div>
  );
});

export const head: DocumentHead = {
  title: "Profile Details",
};
