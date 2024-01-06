/* eslint-disable qwik/loader-location */

import { routeLoader$ } from "@builder.io/qwik-city";
import {
  formAction$,
  valiForm$,
  type InitialValues,
  FormError,
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

export const useProfileFormAction = formAction$<ProfileForm, any>(
  async (values) => {
    const response = await fetch("http://localhost:3000/profile/", {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ ...values }),
    });
    if (response.ok) {
      const data = await response.json();
      return {
        status: "success",
        message: "Your changes have been successfully saved!",
        data,
      };
    } else {
      throw new FormError<ProfileForm>("Failed to save");
    }
  },
  valiForm$(ProfileFormSchema)
);

export const useLinksFormLoader = routeLoader$<InitialValues<LinksForm>>(
  async () => {
    const res = await fetch("http://localhost:3000/links");
    const links = (await res.json()) as LinksForm;
    return links;
  }
);

export const useLinksFormAction = formAction$<LinksForm>(async (values) => {
  const response = await fetch("http://localhost:3000/links/", {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ ...values }),
  });
  if (response.ok) {
    const data = await response.json();
    return {
      status: "success",
      message: "Your changes have been successfully saved!",
      data,
    };
  } else {
    throw new FormError<LinksForm>("Failed to save");
  }
}, valiForm$(LinksFormSchema));
