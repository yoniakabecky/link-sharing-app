import type { Input } from "valibot";
import * as v from "valibot";

export const ProfileFormSchema = v.object({
  firstName: v.string([v.minLength(1, "Please enter your first name.")]),
  lastName: v.string([v.minLength(1, "Please enter your last name.")]),
  email: v.nullable(
    v.string([v.email("The email address is badly formatted.")])
  ),
  avatar: v.nullable(v.string([v.url("Please enter a valid URL.")])),
});

export type ProfileForm = Input<typeof ProfileFormSchema>;

export const LinksFormSchema = v.object({
  links: v.array(
    v.object({
      platform: v.string([v.minLength(1, "Please select a platform.")]),
      link: v.string([v.minLength(1, "Please enter a link.")]),
    })
  ),
});

export type LinksForm = Input<typeof LinksFormSchema>;
