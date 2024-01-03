import { $, component$ } from "@builder.io/qwik";
import { insert, remove, useForm, valiForm$ } from "@modular-forms/qwik";

import FormWrapper from "~/components/form-wrapper/form-wrapper";
import { Icon } from "~/components/icon/icon";
import {
  type LinksForm,
  LinksFormSchema,
  useLinksFormAction,
  useLinksFormLoader,
} from "~/routes/links";
import { TextField } from "../textfield/textfield";
import styles from "./links-editor.module.css";

export default component$(() => {
  const [linksForm, { Form, Field, FieldArray }] = useForm<LinksForm>({
    loader: useLinksFormLoader(),
    action: useLinksFormAction(),
    validate: valiForm$(LinksFormSchema),
    fieldArrays: ["links"],
  });

  return (
    <FormWrapper formName="LinksForm">
      <div q:slot="heading">Customize your links</div>
      <span q:slot="description">
        Add/edit/remove links below and then share all profiles with the world!
      </span>

      <Form id="LinksForm" class={styles.scrollArea}>
        <button
          class={["button button-outlined", styles.addButton]}
          onClick$={() =>
            insert(linksForm, "links", {
              value: { platform: "", link: "" },
            })
          }
        >
          + Add new link
        </button>

        <FieldArray name="links">
          {(fieldArray) => (
            <>
              {fieldArray.items.map((item, index) => (
                <div class={styles.item} key={item}>
                  <div class={styles.itemName}>
                    <Icon name="drag_handle" size={18} />
                    Link #{index + 1}
                  </div>
                  <div
                    class={styles.removeButton}
                    onClick$={$(() =>
                      remove(linksForm, "links", { at: index })
                    )}
                  >
                    Remove
                  </div>

                  <Field name={`links.${index}.platform`}>
                    {(field, props) => (
                      <>
                        <label for={`platform${index}`}>Platform</label>
                        <select
                          {...props}
                          id={`platform${index}`}
                          value={field.value}
                        >
                          <option value="github">GitHub</option>
                          <option value="youtube">YouTube</option>
                          <option value="linkedin">Linkedin</option>
                          <option value="x">X (Twitter)</option>
                          <option value="portfolio">Portfolio</option>
                          <option value="custom">Custom</option>
                        </select>
                      </>
                    )}
                  </Field>

                  <Field name={`links.${index}.link`}>
                    {(field, props) => (
                      <TextField
                        {...props}
                        name={`link.${index}`}
                        label="Link"
                        value={field.value}
                        error={field.error}
                        placeholder="https://www.yoniakabecky.com"
                      />
                    )}
                  </Field>
                </div>
              ))}
            </>
          )}
        </FieldArray>
      </Form>
    </FormWrapper>
  );
});
