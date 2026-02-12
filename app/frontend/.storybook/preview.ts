import type { Preview } from '@storybook/sveltekit';
import '../src/routes/normalize.css';
import '../src/routes/global.css';

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
       color: /(background|color)$/i,
       date: /Date$/i,
      },
    },
  },
};

export default preview;