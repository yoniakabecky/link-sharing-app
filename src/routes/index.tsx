import { component$ } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";

export default component$(() => {
  return (
    <div class="hero">
      <div class="card">
        <div class="contents">
          <h1>devlinks - Link Sharing App</h1>

          <p>
            This is a fully-functional link-sharing app for developers. It
            allows users to create a profile with their details and add links to
            their social media accounts. The app is built with Qwik, a new
            framework for building web apps with JSX, TypeScript, and Web
            Components.
          </p>

          <div class="brief">
            <p>In this app, users should be able to: </p>
            <ul>
              <li>
                Create, read, update, delete links and see previews in the
                mobile mockup
              </li>
              <li>
                Receive validations if the links form is submitted without a URL
                or with the wrong URL pattern for the platform
              </li>
              <li>
                <b>WIP</b>
                Drag and drop links to reorder them
              </li>
              <li>
                Add profile details like profile picture, first name, last name,
                and email
              </li>
              <li>
                Receive validations if the profile details form is saved with no
                first or last name
              </li>
              <li>
                <b>WIP</b>
                Preview their devlinks profile and copy the link to their
                clipboard
              </li>
              <li>
                View the optimal layout for the interface depending on their
                device's screen size
              </li>
              <li>
                See hover and focus states for all interactive elements on the
                page
              </li>
              <li>
                <b>WIP</b>
                Save details to a database (build the project as a full-stack
                app)
              </li>
              <li>
                <b>WIP</b>Create an account and log in (add user authentication
                to the full-stack app)
              </li>
            </ul>

            <a
              href="https://www.frontendmentor.io/challenges/linksharing-app-Fbt7yweGsT"
              target="_blank"
            >
              more details (frontendmentor.io)
            </a>
          </div>
        </div>
      </div>
    </div>
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
