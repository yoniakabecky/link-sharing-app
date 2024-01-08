import { component$ } from "@builder.io/qwik";
import { type DocumentHead } from "@builder.io/qwik-city";
import {
  getValues,
  insert,
  remove,
  useForm,
  valiForm$,
} from "@modular-forms/qwik";
import FormWrapper from "~/components/form-wrapper/form-wrapper";
import { Icon } from "~/components/icon/icon";
import MockView from "~/components/mock-view/mock-view";
import { Select } from "~/components/select/select";
import { TextField } from "~/components/textfield/textfield";
import { useLinksFormAction, useLinksLoader } from "../layout";
import type { LinksForm, LinkItem } from "../schema";
import { LinksFormSchema } from "../schema";
import styles from "./links.module.css";
import { platformOptions } from "~/models/platform";

export default component$(() => {
  const [linksForm, { Form, Field, FieldArray }] = useForm<LinksForm>({
    loader: useLinksLoader(),
    action: useLinksFormAction(),
    validate: valiForm$(LinksFormSchema),
    fieldArrays: ["links"],
  });

  const values = getValues(linksForm, "links") as LinkItem[];

  return (
    <div class="editor">
      <div class="card left">
        <MockView links={values} />
      </div>

      <div class="card right">
        <FormWrapper formName="LinksForm">
          <div q:slot="heading">Customize your links</div>
          <span q:slot="description">
            Add/edit/remove links below and then share all profiles with the
            world!
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
                        onClick$={() =>
                          remove(linksForm, "links", { at: index })
                        }
                      >
                        Remove
                      </div>

                      <Field name={`links.${index}.platform`}>
                        {(field, props) => (
                          <>
                            <Select
                              {...props}
                              label="Platform"
                              name={`platform${index}`}
                              value={field.value}
                              error={field.error}
                              options={platformOptions}
                            />
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
      </div>
    </div>
  );
});

export const head: DocumentHead = {
  title: "Links",
};
