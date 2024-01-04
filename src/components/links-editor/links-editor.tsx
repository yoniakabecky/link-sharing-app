import { component$ } from "@builder.io/qwik";
import { insert, remove, useForm, valiForm$ } from "@modular-forms/qwik";

import FormWrapper from "~/components/form-wrapper/form-wrapper";
import { Icon } from "~/components/icon/icon";
import { Select } from "~/components/select/select";
import { TextField } from "~/components/textfield/textfield";
import { useLinksFormLoader } from "~/routes/layout";
import { type LinksForm, LinksFormSchema } from "~/routes/schema";
import styles from "./links-editor.module.css";

export default component$(() => {
  const [linksForm, { Form, Field, FieldArray }] = useForm<LinksForm>({
    loader: useLinksFormLoader(),
    // action: useLinksFormAction(),
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
                    onClick$={() => remove(linksForm, "links", { at: index })}
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
                          options={[
                            { label: "GitHub", value: "github" },
                            { label: "YouTube", value: "youtube" },
                            { label: "Linkedin", value: "linkedin" },
                            { label: "X (Twitter)", value: "x" },
                            { label: "Portfolio", value: "portfolio" },
                            { label: "Custom", value: "custom" },
                          ]}
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
  );
});
