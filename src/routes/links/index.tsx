import { component$ } from "@builder.io/qwik";
import { routeLoader$, type DocumentHead } from "@builder.io/qwik-city";
import {
  type InitialValues,
  formAction$,
  valiForm$,
} from "@modular-forms/qwik";
import type { Input } from "valibot";
import * as v from "valibot";

import LinksEditor from "~/components/links-editor/links-editor";
import MockView from "~/components/mock-view/mock-view";

export const LinksFormSchema = v.object({
  links: v.array(
    v.object({
      platform: v.string([v.minLength(1, "Please enter a platform.")]),
      link: v.string([v.minLength(1, "Please enter a link.")]),
    })
  ),
});

export type LinksForm = Input<typeof LinksFormSchema>;

export const useLinksFormLoader = routeLoader$<InitialValues<LinksForm>>(() => {
  return {
    links: [
      {
        platform: "github",
        link: "https://github.com/yoniakabecky",
      },
      {
        platform: "",
        link: "",
      },
    ],
  };
});

export const useLinksFormAction = formAction$<LinksForm>((values) => {
  // Runs on server
  console.log({ values });
}, valiForm$(LinksFormSchema));

export default component$(() => {
  return (
    <div class="editor">
      <div class="card left">
        <MockView />
      </div>
      <div class="card right">
        <LinksEditor />
      </div>
    </div>
  );
});

export const head: DocumentHead = {
  title: "Links",
};
