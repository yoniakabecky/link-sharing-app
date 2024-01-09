# Link Sharing App 🚀

This is my [Qwik](https://qwik.builder.io/) practice project✨  
  
The original idea and designs are from [Frontend Mentor](https://www.frontendmentor.io/challenges/linksharing-app-Fbt7yweGsT).

## Overview 👀

### Built with

- [Qwik](https://qwik.builder.io/)
- [Modular Forms](https://modularforms.dev/)
- [Valibot](https://valibot.dev/)
- [json-server](https://github.com/typicode/json-server) - until I implement database

### The original challenge

Users should be able to:

- [x] Create, read, update, delete links and see previews in the mobile mockup
- [x] Receive validations if the links form is submitted without a URL or with the wrong URL pattern for the platform
- [ ] Drag and drop links to reorder them
- [ ] Add profile details like profile picture, first name, last name, and email  
    → profile picture is not done yet
- [x] Receive validations if the profile details form is saved with no first or last name
- [x] Preview their devlinks profile and copy the link to their clipboard
- [x] View the optimal layout for the interface depending on their device's screen size
- [x] See hover and focus states for all interactive elements on the page
- [ ] Bonus: Save details to a database (build the project as a full-stack app)
- [ ] Bonus: Create an account and log in (add user authentication to the full-stack app)


## Try out in a local environment 👩‍💻

### 1. Install dependencies 

```shell
pnpm i
```

I'm using [pnpm](https://pnpm.io/) 😉

### 2. Run json-server

```shell
pnpm run json # or npx json-server ./data/db.json
```

This will run on "https://localhost:3000" by default. If you want to change the port, use the `--port` flag.

### 3. Start development server

```shell
pnpm start
```

...and Voila! 🥳
