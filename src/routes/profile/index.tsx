import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
import { routeLoader$ } from "@builder.io/qwik-city";
import {
  formAction$,
  valiForm$,
  type InitialValues,
} from "@modular-forms/qwik";
import type { Input } from "valibot";
import * as v from "valibot";

import MockView from "~/components/mock-view/mock-view";
import ProfileEditor from "~/components/profile-editor/profile-editor";

export const ProfileFormSchema = v.object({
  firstName: v.string([v.minLength(1, "Please enter your first name.")]),
  lastName: v.string([v.minLength(1, "Please enter your last name.")]),
  email: v.nullable(
    v.string([v.email("The email address is badly formatted.")])
  ),
});

export type ProfileForm = Input<typeof ProfileFormSchema>;

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

export default component$(() => {
  return (
    <div class="editor">
      <div class="card left">
        <MockView />
      </div>
      <div class="card right">
        <ProfileEditor />
      </div>
    </div>
  );
});

export const head: DocumentHead = {
  title: "Profile Details",
};
