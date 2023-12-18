import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <>
      <div>Hello World</div>
    </>
  );
});

export const head: DocumentHead = {
  title: "Link Sharing App",
  meta: [
    {
      name: "description",
      content: "A fully-functional link-sharing app for developers",
    },
  ],
};
