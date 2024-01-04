import { component$ } from "@builder.io/qwik";
import { type DocumentHead } from "@builder.io/qwik-city";

import LinksEditor from "~/components/links-editor/links-editor";
import MockView from "~/components/mock-view/mock-view";

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
