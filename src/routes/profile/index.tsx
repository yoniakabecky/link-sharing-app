import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";

import MockView from "~/components/mock-view/mock-view";
import ProfileEditor from "~/components/profile-editor/profile-editor";

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
