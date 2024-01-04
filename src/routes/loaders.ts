/* eslint-disable qwik/loader-location */

import { routeLoader$ } from "@builder.io/qwik-city";
import {
  formAction$,
  valiForm$,
  type InitialValues,
} from "@modular-forms/qwik";
import type { LinksForm, ProfileForm } from "./schema";
import { ProfileFormSchema, LinksFormSchema } from "./schema";

export const useProfileFormLoader = routeLoader$<InitialValues<ProfileForm>>(
  async () => {
    const res = await fetch("http://localhost:3000/profile");
    const profile = (await res.json()) as ProfileForm;
    return profile;
  }
);

export const useProfileFormAction = formAction$<ProfileForm>((values) => {
  // Runs on server
  console.log({ values });
}, valiForm$(ProfileFormSchema));

export const useLinksFormLoader = routeLoader$<InitialValues<LinksForm>>(
  async () => {
    const res = await fetch("http://localhost:3000/links");
    const links = (await res.json()) as LinksForm;
    return links;
  }
);

export const useLinksFormAction = formAction$<LinksForm>((values) => {
  // Runs on server
  console.log({ values });
}, valiForm$(LinksFormSchema));
